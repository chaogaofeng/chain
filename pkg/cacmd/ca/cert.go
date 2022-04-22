package ca

import (
	"fmt"

	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/algo"
)

type Cert interface {
	WritePrivateKeytoMem() ([]byte, error)
	GetPubkeyFromCert() (crypto.PubKey, error)
	VerifyCertFromRoot(rootCert Cert) error
}

func ReadCertificateFromMem(data []byte) (Cert, error) {
	switch algo.Algo {
	case algo.SM2:
		return ReadSM2CertFromMem(data)

	default:
		return ReadX509CertFromMem(data)
	}
}

func UnexpectedPubKeyAlgo(expected string, pubkey interface{}) error {
	return fmt.Errorf(
		"x509: signature algorithm specifies an %s public key, but have public key of type %T",
		expected, pubkey,
	)
}
