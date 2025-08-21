package webhook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignatureVerifier_Verify(t *testing.T) {
	sign := `MEQCIEaJMN/d0xcZoEgI1zya+yE6GYJb2f2osBZMPgjtXNUiAiAGVfUR9dxj2Ix7blF7MjMdAU2VZcpuyUuB6zncVoFadg==`
	publicKey := `LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUZrd0V3WUhLb1pJemowQ0FRWUlLb1pJemowREFRY0RRZ0FFK0UxRnBVZzczYmhGdmp2SzlrMlhJeTZtQkU1MQpib2F0RU1qU053Z1l5ZW55blpZQWh3Z3dyTGhNY0FpT25SYzNXWGNyMGRrY2NvVnFXcVBhWVQ5T3hRPT0KLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0tCg==`

	verifier, err := NewSignatureVerifier(publicKey)

	assert.NoError(t, err)
	assert.NotNil(t, verifier)

	tests := map[string]struct {
		Body string
		IsOk bool
	}{
		"valid": {
			Body: `{"invoiceId":"250811tUZjKAWjrnb9b","status":"success","payMethod":"wallet","amount":20200,"ccy":980,"finalAmount":20200,"createdDate":"2025-08-11T06:08:52Z","modifiedDate":"2025-08-11T06:08:54Z","reference":"ce223cb7-1c95-4f3b-8a3e-2a5fe21bce6c","destination":"Розрахунок за дату 2025-08-11 по картці {{masked_pan}} в торговій точці {{terminal_owner}} ({{terminal_retailer}})","paymentInfo":{"rrn":"061673331001","approvalCode":"117524","tranId":"19277588","terminal":"XPZ10001","bank":"Універсал Банк","paymentSystem":"visa","country":"804","fee":263,"paymentMethod":"wallet","maskedPan":"44440311******39"}}`,
			IsOk: true,
		},
		"invalid_1": {
			Body: `{}`,
			IsOk: false,
		},
		"invalid_2": {
			Body: `test`,
			IsOk: false,
		},
	}

	for name, val := range tests {
		t.Run(name, func(t *testing.T) {
			ok, err := verifier.Verify(sign, []byte(val.Body))

			assert.NoError(t, err)
			assert.Equal(t, val.IsOk, ok)
		})
	}
}

func TestNewSignatureVerifier_Error(t *testing.T) {
	tests := map[string]string{
		"empty":       "",
		"new line":    "\n",
		"numbers":     "123",
		"random text": "c29tZSByYW5kb20gdGV4dA==",
		"not ECDSA": `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAtv....
-----END PUBLIC KEY-----`,
	}

	for name, val := range tests {
		t.Run(name, func(t *testing.T) {
			verifier, err := NewSignatureVerifier(val)

			assert.Error(t, err)
			assert.Nil(t, verifier)
		})
	}
}
