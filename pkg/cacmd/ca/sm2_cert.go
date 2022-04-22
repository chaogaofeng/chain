package ca

import (
	"crypto/ecdsa"
	"errors"

	"github.com/tendermint/tendermint/crypto"
	tmsm2 "github.com/tendermint/tendermint/crypto/sm2"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

// Sm2Cert defines sm2 signed X509 certificate
type Sm2Cert struct {
	*x509.Certificate
	*sm2.PrivateKey
}

func ReadSM2CertFromMem(data []byte) (Cert, error) {
	cert, err := x509.ReadCertificateFromPem(data)
	return Sm2Cert{cert, nil}, err
}

func (sm2c Sm2Cert) WritePrivateKeytoMem() ([]byte, error) {
	return x509.WritePrivateKeyToPem(sm2c.PrivateKey, nil)
}

func (sm2c Sm2Cert) VerifyCertFromRoot(rootCert Cert) error {
	if rc, ok := rootCert.(Sm2Cert); ok {
		return sm2c.Certificate.CheckSignatureFrom(rc.Certificate)
	}
	return errors.New("can not verify sm2 certificate by other algorithm certificate")
}

func (sm2c Sm2Cert) GetPubkeyFromCert() (crypto.PubKey, error) {
	expectedPubKeyAlgo := sm2c.Certificate.PublicKeyAlgorithm
	pub, ok := sm2c.Certificate.PublicKey.(*ecdsa.PublicKey)
	if !ok || expectedPubKeyAlgo != x509.ECDSA {
		return nil, UnexpectedPubKeyAlgo("ECDSA", sm2c.Certificate.PublicKey)
	}
	switch pub.Curve {
	case sm2.P256Sm2():
		sm2Pub := sm2.PublicKey{
			Curve: pub.Curve,
			X:     pub.X,
			Y:     pub.Y,
		}

		compPubkey := sm2.Compress(&sm2Pub)
		var pubKey tmsm2.PubKeySm2
		copy(pubKey[:], compPubkey)
		return pubKey, nil
	default:
		return nil, UnexpectedPubKeyAlgo("SM2", sm2c.Certificate.PublicKey)
	}
}
