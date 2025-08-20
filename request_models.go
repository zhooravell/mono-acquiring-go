package monoacquiring

const (
	PaymentTypeDebit = "debit"
	PaymentTypeHold  = "hold"

	DisplayTypeIframe = "iframe"

	DiscountTypeDiscount    = "DISCOUNT"
	DiscountTypeExtraCharge = "EXTRA_CHARGE"

	DiscountModePercent = "PERCENT"
	DiscountModeValue   = "VALUE"
)

type errorData struct {
	Code    string `json:"errCode"`
	Message string `json:"errText"`
}

type SaveCardData struct {
	WalletID *string `json:"walletId,omitempty"`
	SaveCard bool    `json:"saveCard"`
}

type Discount struct {
	Type  string  `json:"type" validate:"required,oneof=DISCOUNT EXTRA_CHARGE"`
	Mode  string  `json:"mode" validate:"required,oneof=PERCENT VALUE"`
	Value float64 `json:"value" validate:"required,min=0.01"`
}

type BasketOrder struct {
	Footer          *string    `json:"footer,omitempty"`
	SplitReceiverID *string    `json:"splitReceiverId,omitempty"`
	Barcode         *string    `json:"barcode,omitempty"`
	Total           *int64     `json:"total"`
	Icon            *string    `json:"icon,omitempty"`
	Unit            *string    `json:"unit,omitempty"`
	Header          *string    `json:"header,omitempty"`
	Uktzed          *string    `json:"uktzed,omitempty"`
	Name            string     `json:"name" validate:"required"`
	Code            string     `json:"code" validate:"required"`
	Tax             []int64    `json:"tax,omitempty"`
	Discounts       []Discount `json:"discounts,omitempty" validate:"dive"`
	Qty             float64    `json:"qty" validate:"required"`
	Sum             int64      `json:"sum" validate:"required"`
}

type MerchantPaymentInfo struct {
	Reference      *string       `json:"reference,omitempty"`
	Destination    *string       `json:"destination,omitempty" validate:"max=280"`
	Comment        *string       `json:"comment,omitempty" validate:"max=280"`
	CustomerEmails []string      `json:"customerEmails" validate:"dive,email"`
	Discounts      []Discount    `json:"discounts,omitempty" validate:"dive"`
	BasketOrder    []BasketOrder `json:"basketOrder,omitempty" validate:"dive"`
}
