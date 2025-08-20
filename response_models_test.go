package monoacquiring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQRAmountType(t *testing.T) {
	m := QRAmountType(qrAmountTypeMerchant)

	assert.Equal(t, "merchant", m.String())
	assert.True(t, m.IsMerchant())
	assert.False(t, m.IsClient())
	assert.False(t, m.IsFix())

	c := QRAmountType(qrAmountTypeClient)

	assert.Equal(t, "client", c.String())
	assert.True(t, c.IsClient())
	assert.False(t, c.IsMerchant())
	assert.False(t, c.IsFix())

	f := QRAmountType(qrAmountTypeFix)

	assert.Equal(t, "fix", f.String())
	assert.True(t, f.IsFix())
	assert.False(t, f.IsClient())
	assert.False(t, f.IsMerchant())
}
