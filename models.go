package monoacquiring

const (
	PaymentTypeDebit = "debit"
	PaymentTypeHold  = "hold"
)

type errorData struct {
	Code    string `json:"errCode"`
	Message string `json:"errText"`
}

type SaveCardData struct {
	SaveCard bool    `json:"saveCard"`
	WalletID *string `json:"walletId,omitempty"`
}

type Discount struct {
	Type  string  `json:"type" validate:"required,oneof=DISCOUNT EXTRA_CHARGE"`
	Mode  string  `json:"mode" validate:"required,oneof=PERCENT VALUE"`
	Value float64 `json:"value" validate:"required,min=0.01"`
}

type BasketOrder struct {
	Name            string     `json:"name" validate:"required"`
	Qty             float64    `json:"qty" validate:"required"`
	Sum             int64      `json:"sum" validate:"required"`
	Total           *int64     `json:"total"`
	Icon            *string    `json:"icon,omitempty"`
	Unit            *string    `json:"unit,omitempty"`
	Code            string     `json:"code" validate:"required"`
	Barcode         *string    `json:"barcode,omitempty"`
	Header          *string    `json:"header,omitempty"`
	Footer          *string    `json:"footer,omitempty"`
	Tax             []int64    `json:"tax,omitempty"`
	Uktzed          *string    `json:"uktzed,omitempty"`
	SplitReceiverID *string    `json:"splitReceiverId,omitempty"`
	Discounts       []Discount `json:"discounts,omitempty" validate:"dive"`
}

type MerchantPaymentInfo struct {
	Reference      *string       `json:"reference,omitempty"`
	Destination    *string       `json:"destination,omitempty" validate:"max=280"`
	Comment        *string       `json:"comment,omitempty" validate:"max=280"`
	CustomerEmails []string      `json:"customerEmails" validate:"dive,email"`
	Discounts      []Discount    `json:"discounts,omitempty" validate:"dive"`
	BasketOrder    []BasketOrder `json:"basketOrder,omitempty" validate:"dive"`
}
