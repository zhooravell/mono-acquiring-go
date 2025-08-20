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

func TestFiscalCheckStatus(t *testing.T) {
	n := FiscalCheckStatus("new")
	assert.Equal(t, "new", n.String())
	assert.True(t, n.IsNew())
	assert.False(t, n.IsProcess())
	assert.False(t, n.IsDone())
	assert.False(t, n.IsFailed())

	p := FiscalCheckStatus("process")
	assert.Equal(t, "process", p.String())
	assert.True(t, p.IsProcess())
	assert.False(t, p.IsNew())
	assert.False(t, p.IsDone())
	assert.False(t, p.IsFailed())

	d := FiscalCheckStatus("done")
	assert.Equal(t, "done", d.String())
	assert.True(t, d.IsDone())
	assert.False(t, d.IsNew())
	assert.False(t, d.IsProcess())
	assert.False(t, d.IsFailed())

	f := FiscalCheckStatus("failed")
	assert.Equal(t, "failed", f.String())
	assert.True(t, f.IsFailed())
	assert.False(t, f.IsNew())
	assert.False(t, f.IsProcess())
	assert.False(t, f.IsDone())
}

func TestFiscalCheckType(t *testing.T) {
	s := FiscalCheckType("sale")
	assert.Equal(t, "sale", s.String())
	assert.True(t, s.IsSale())
	assert.False(t, s.IsReturn())

	r := FiscalCheckType("return")
	assert.Equal(t, "return", r.String())
	assert.True(t, r.IsReturn())
	assert.False(t, r.IsSale())
}

func TestFiscalCheckSource(t *testing.T) {
	c := FiscalCheckSource("checkbox")
	assert.Equal(t, "checkbox", c.String())
	assert.True(t, c.IsCheckBox())
	assert.False(t, c.IsMonoPay())
	assert.False(t, c.IsVchasnoKasa())

	m := FiscalCheckSource("monopay")
	assert.Equal(t, "monopay", m.String())
	assert.True(t, m.IsMonoPay())
	assert.False(t, m.IsCheckBox())
	assert.False(t, m.IsVchasnoKasa())

	v := FiscalCheckSource("vchasnokasa")
	assert.Equal(t, "vchasnokasa", v.String())
	assert.True(t, v.IsVchasnoKasa())
	assert.False(t, v.IsCheckBox())
	assert.False(t, v.IsMonoPay())
}

func TestStatementStatus(t *testing.T) {
	h := StatementStatus("hold")
	assert.Equal(t, "hold", h.String())
	assert.True(t, h.IsHold())
	assert.False(t, h.IsProcessing())
	assert.False(t, h.IsSuccess())
	assert.False(t, h.IsProcessing())

	p := StatementStatus("processing")
	assert.Equal(t, "processing", p.String())
	assert.True(t, p.IsProcessing())
	assert.False(t, p.IsHold())
	assert.False(t, p.IsSuccess())
	assert.False(t, p.IsFailure())

	s := StatementStatus("success")
	assert.Equal(t, "success", s.String())
	assert.True(t, s.IsSuccess())
	assert.False(t, s.IsHold())
	assert.False(t, s.IsProcessing())
	assert.False(t, s.IsFailure())

	f := StatementStatus("failure")
	assert.Equal(t, "failure", f.String())
	assert.True(t, f.IsFailure())
	assert.False(t, f.IsHold())
	assert.False(t, f.IsProcessing())
	assert.False(t, f.IsSuccess())
}

func TestStatementPaymentScheme(t *testing.T) {
	b30 := StatementPaymentScheme("bnpl_later_30")
	assert.Equal(t, "bnpl_later_30", b30.String())
	assert.True(t, b30.IsBnplLater30())
	assert.False(t, b30.IsBnplParts4())
	assert.False(t, b30.IsFull())

	b4 := StatementPaymentScheme("bnpl_parts_4")
	assert.Equal(t, "bnpl_parts_4", b4.String())
	assert.True(t, b4.IsBnplParts4())
	assert.False(t, b4.IsBnplLater30())
	assert.False(t, b4.IsFull())

	f := StatementPaymentScheme("full")
	assert.Equal(t, "full", f.String())
	assert.True(t, f.IsFull())
	assert.False(t, f.IsBnplLater30())
	assert.False(t, f.IsBnplParts4())
}
