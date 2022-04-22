package cacmd

import (
	"bytes"
	"crypto/ed25519"
	"errors"
	"fmt"
	"os/exec"

	"github.com/cosmos/cosmos-sdk/crypto/hd"

	cryptoed25519 "github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	cryptossecp256k1 "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptosm2 "github.com/cosmos/cosmos-sdk/crypto/keys/sm2"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/glodnet/chain/pkg/cacmd/ca"
	tmcrypto "github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/algo"
	tmed25519 "github.com/tendermint/tendermint/crypto/ed25519"
	tmsecp256k1 "github.com/tendermint/tendermint/crypto/secp256k1"
	tmsm2 "github.com/tendermint/tendermint/crypto/sm2"
)

func Genkey(privKey tmcrypto.PrivKey) ([]byte, error) {
	switch pk := privKey.(type) {
	case tmsm2.PrivKeySm2:
		privKey := privKey.(tmsm2.PrivKeySm2)
		return ca.Sm2Cert{PrivateKey: privKey.GetPrivateKey()}.WritePrivateKeytoMem()
	case tmed25519.PrivKey:
		priKey := make([]byte, ed25519.PrivateKeySize)
		copy(priKey, pk[:])
		return ca.X509Cert{PrivateKey: ed25519.PrivateKey(priKey)}.WritePrivateKeytoMem()
	case tmsecp256k1.PrivKey:
		return nil, errors.New("unsupported algorithm type")
	default:
		return nil, errors.New("unsupported algorithm type")
	}
}

func ToTmPrivKeyInterface(priv cryptotypes.PrivKey) (tmcrypto.PrivKey, error) {
	switch tp := priv.(type) {
	case *cryptoed25519.PrivKey:
		privKeyBytes := make([]byte, tmed25519.PrivateKeySize)
		copy(privKeyBytes[:], tp.Key)
		return tmed25519.PrivKey(privKeyBytes), nil
	case *cryptosm2.PrivKey:
		tmpriv := tmsm2.PrivKeySm2{}
		copy(tmpriv[:], tp.Key)
		return tmpriv, nil
	case *cryptossecp256k1.PrivKey:
		privKeyBytes := make([]byte, tmsecp256k1.PrivKeySize)
		copy(privKeyBytes[:], tp.Key)
		return tmsecp256k1.PrivKey(privKeyBytes[:]), nil
	default:
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "cannot convert %v to Tendermint private key", tp)
	}
}

func GenRootCert(algoStr string, keyPath, certPath, subj string, days string) error {
	switch algoStr {
	case algo.SM2:
		cmd := exec.Command("openssl", "ecparam", "-genkey", "-name", "SM2", "-out", keyPath)
		if err := executeCmd(cmd); err != nil {
			return err
		}
	case algo.ED25519:
		cmd := exec.Command("openssl", "genpkey", "-algorithm", "ED25519", "-out", keyPath)
		if err := executeCmd(cmd); err != nil {
			return err
		}
	case string(hd.Secp256k1Type):
		cmd := exec.Command("openssl", "ecparam", "-genkey", "-name", "secp256k1", "-out", keyPath)
		if err := executeCmd(cmd); err != nil {
			return err
		}
	default:
		return fmt.Errorf("provided algorithm %q is not supported", algoStr)
	}
	return GenSelfSignCert(algoStr, keyPath, certPath, subj, days)
}

func GenSelfSignCert(algoStr string, keyPath, certPath, subj string, days string) error {
	switch algoStr {
	case algo.SM2:
		cmd := exec.Command(
			"openssl", "req", "-new", "-x509", "-sm3", "-sigopt", "distid:1234567812345678",
			"-key", keyPath, "-subj", subj, "-out", certPath, "-days", days,
		)
		return executeCmd(cmd)
	case algo.ED25519:
		cmd := exec.Command(
			"openssl", "req", "-new", "-x509",
			"-key", keyPath, "-subj", subj, "-out", certPath, "-days", days,
		)
		return executeCmd(cmd)
	case string(hd.Secp256k1Type):
		cmd := exec.Command(
			"openssl", "req", "-new", "-x509",
			"-key", keyPath, "-subj", subj, "-out", certPath, "-days", days,
		)
		return executeCmd(cmd)
	default:
		return fmt.Errorf("provided algorithm %q is not supported", algoStr)
	}
}

func GenCertRequest(algoStr string, keyPath, cerPath, subj string) error {
	switch algoStr {
	case algo.SM2:
		cmd := exec.Command(
			"openssl", "req", "-new", "-sm3", "-sigopt", "distid:1234567812345678",
			"-key", keyPath, "-subj", subj, "-out", cerPath,
		)
		return executeCmd(cmd)
	case algo.ED25519:
		cmd := exec.Command(
			"openssl", "req", "-new",
			"-key", keyPath, "-subj", subj, "-out", cerPath,
		)
		return executeCmd(cmd)
	case string(hd.Secp256k1Type):
		cmd := exec.Command(
			"openssl", "req", "-new",
			"-key", keyPath, "-subj", subj, "-out", cerPath,
		)
		return executeCmd(cmd)
	default:
		return fmt.Errorf("provided algorithm %q is not supported", algoStr)
	}
}

func IssueCert(algoStr string, cerPath, caPath, caKeyPath, certPath string, days string) error {
	switch algoStr {
	case algo.SM2:
		cmd := exec.Command(
			"openssl", "x509", "-req", "-in", cerPath,
			"-CA", caPath, "-CAkey", caKeyPath, "-CAcreateserial", "-out", certPath, "-days", days,
			"-sm3", "-sigopt", "distid:1234567812345678", "-vfyopt", "distid:1234567812345678",
		)
		return executeCmd(cmd)
	case algo.ED25519:
		cmd := exec.Command(
			"openssl", "x509", "-req", "-in", cerPath,
			"-CA", caPath, "-CAkey", caKeyPath, "-CAcreateserial", "-out", certPath, "-days", days,
		)
		return executeCmd(cmd)
	case string(hd.Secp256k1Type):
		cmd := exec.Command(
			"openssl", "x509", "-req", "-in", cerPath,
			"-CA", caPath, "-CAkey", caKeyPath, "-CAcreateserial", "-out", certPath, "-days", days,
		)
		return executeCmd(cmd)
	default:
		return fmt.Errorf("provided algorithm %q is not supported", algoStr)
	}
}

func executeCmd(cmd *exec.Cmd) error {
	var stdOut bytes.Buffer
	var stdErr bytes.Buffer
	cmd.Stderr = &stdErr
	cmd.Stdout = &stdOut
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s: %s", err, stdErr.String())
	}
	return nil
}
