package webhook

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"

	"github.com/pkg/errors"
)

type SignatureVerifier struct {
	pubKey *ecdsa.PublicKey
}

func NewSignatureVerifier(pubKey string) (*SignatureVerifier, error) {
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	block, _ := pem.Decode(pubKeyBytes)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing the public key")
	}

	genericPubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var (
		ok bool
		sv = SignatureVerifier{}
	)

	if sv.pubKey, ok = genericPubKey.(*ecdsa.PublicKey); !ok {
		return nil, errors.New("failed to assert type of public key")
	}

	return &sv, nil
}

func (sv *SignatureVerifier) Verify(signature string, body []byte) (bool, error) {
	sign, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return false, errors.WithStack(err)
	}

	hash := sha256.Sum256(body)

	return ecdsa.VerifyASN1(sv.pubKey, hash[:], sign), nil
}
