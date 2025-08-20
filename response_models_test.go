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

func TestSyncPaymentStatus(t *testing.T) {
	c := SyncPaymentStatus("created")
	assert.Equal(t, "created", c.String())
	assert.True(t, c.IsCreated())
	assert.False(t, c.IsProcessing())
	assert.False(t, c.IsSuccess())
	assert.False(t, c.IsFailure())
	assert.False(t, c.IsHold())
	assert.False(t, c.IsReversed())
	assert.False(t, c.IsExpired())

	p := SyncPaymentStatus("processing")
	assert.Equal(t, "processing", p.String())
	assert.True(t, p.IsProcessing())
	assert.False(t, p.IsCreated())
	assert.False(t, p.IsSuccess())
	assert.False(t, p.IsFailure())
	assert.False(t, p.IsHold())
	assert.False(t, p.IsReversed())
	assert.False(t, p.IsExpired())

	h := SyncPaymentStatus("hold")
	assert.Equal(t, "hold", h.String())
	assert.True(t, h.IsHold())
	assert.False(t, h.IsCreated())
	assert.False(t, h.IsSuccess())
	assert.False(t, h.IsFailure())
	assert.False(t, h.IsProcessing())
	assert.False(t, h.IsReversed())
	assert.False(t, h.IsExpired())

	s := SyncPaymentStatus("success")
	assert.Equal(t, "success", s.String())
	assert.True(t, s.IsSuccess())
	assert.False(t, s.IsCreated())
	assert.False(t, s.IsFailure())
	assert.False(t, s.IsHold())
	assert.False(t, s.IsProcessing())
	assert.False(t, s.IsReversed())
	assert.False(t, s.IsExpired())

	f := SyncPaymentStatus("failure")
	assert.Equal(t, "failure", f.String())
	assert.True(t, f.IsFailure())
	assert.False(t, f.IsCreated())
	assert.False(t, f.IsSuccess())
	assert.False(t, f.IsHold())
	assert.False(t, f.IsProcessing())
	assert.False(t, f.IsReversed())
	assert.False(t, f.IsExpired())

	r := SyncPaymentStatus("reversed")
	assert.Equal(t, "reversed", r.String())
	assert.True(t, r.IsReversed())
	assert.False(t, r.IsCreated())
	assert.False(t, r.IsSuccess())
	assert.False(t, r.IsFailure())
	assert.False(t, r.IsHold())
	assert.False(t, r.IsProcessing())
	assert.False(t, r.IsExpired())

	e := SyncPaymentStatus("expired")
	assert.Equal(t, "expired", e.String())
	assert.True(t, e.IsExpired())
	assert.False(t, e.IsCreated())
	assert.False(t, e.IsSuccess())
	assert.False(t, e.IsFailure())
	assert.False(t, e.IsHold())
	assert.False(t, e.IsProcessing())
	assert.False(t, e.IsReversed())
}

func TestWalletDataStatus(t *testing.T) {
	n := WalletDataStatus("new")
	assert.Equal(t, "new", n.String())
	assert.True(t, n.IsNew())
	assert.False(t, n.IsCreated())
	assert.False(t, n.IsFailed())

	c := WalletDataStatus("created")
	assert.Equal(t, "created", c.String())
	assert.True(t, c.IsCreated())
	assert.False(t, c.IsNew())
	assert.False(t, c.IsFailed())

	f := WalletDataStatus("failed")
	assert.Equal(t, "failed", f.String())
	assert.True(t, f.IsFailed())
	assert.False(t, f.IsNew())
	assert.False(t, f.IsCreated())
}

func TestPaymentSystem(t *testing.T) {
	v := PaymentSystem("visa")
	assert.Equal(t, "visa", v.String())
	assert.True(t, v.IsVisa())
	assert.False(t, v.IsMasterCard())

	m := PaymentSystem("mastercard")
	assert.Equal(t, "mastercard", m.String())
	assert.True(t, m.IsMasterCard())
	assert.False(t, m.IsVisa())
}

func TestPaymentMethod(t *testing.T) {
	p := PaymentMethod("pan")
	assert.Equal(t, "pan", p.String())
	assert.True(t, p.IsPAN())
	assert.False(t, p.IsApple())
	assert.False(t, p.IsGoogle())
	assert.False(t, p.IsMonobank())
	assert.False(t, p.IsWallet())
	assert.False(t, p.IsDirect())

	a := PaymentMethod("apple")
	assert.Equal(t, "apple", a.String())
	assert.True(t, a.IsApple())
	assert.False(t, a.IsPAN())
	assert.False(t, a.IsGoogle())
	assert.False(t, a.IsMonobank())
	assert.False(t, a.IsWallet())
	assert.False(t, a.IsDirect())

	g := PaymentMethod("google")
	assert.Equal(t, "google", g.String())
	assert.True(t, g.IsGoogle())
	assert.False(t, g.IsPAN())
	assert.False(t, g.IsApple())
	assert.False(t, g.IsMonobank())
	assert.False(t, g.IsWallet())
	assert.False(t, g.IsDirect())

	m := PaymentMethod("monobank")
	assert.Equal(t, "monobank", m.String())
	assert.True(t, m.IsMonobank())
	assert.False(t, m.IsPAN())
	assert.False(t, m.IsApple())
	assert.False(t, m.IsGoogle())
	assert.False(t, m.IsWallet())
	assert.False(t, m.IsDirect())

	w := PaymentMethod("wallet")
	assert.Equal(t, "wallet", w.String())
	assert.True(t, w.IsWallet())
	assert.False(t, w.IsPAN())
	assert.False(t, w.IsApple())
	assert.False(t, w.IsGoogle())
	assert.False(t, w.IsMonobank())
	assert.False(t, w.IsDirect())

	d := PaymentMethod("direct")
	assert.Equal(t, "direct", d.String())
	assert.True(t, d.IsDirect())
	assert.False(t, d.IsPAN())
	assert.False(t, d.IsApple())
	assert.False(t, d.IsGoogle())
	assert.False(t, d.IsMonobank())
	assert.False(t, d.IsWallet())
}

func TestCancelListItemStatus(t *testing.T) {
	p := CancelListItemStatus("processing")
	assert.Equal(t, "processing", p.String())
	assert.True(t, p.IsProcessing())
	assert.False(t, p.IsSuccess())
	assert.False(t, p.IsFailure())

	s := CancelListItemStatus("success")
	assert.Equal(t, "success", s.String())
	assert.True(t, s.IsSuccess())
	assert.False(t, s.IsProcessing())
	assert.False(t, s.IsFailure())

	f := CancelListItemStatus("failure")
	assert.Equal(t, "failure", f.String())
	assert.True(t, f.IsFailure())
	assert.False(t, f.IsProcessing())
	assert.False(t, f.IsSuccess())
}
