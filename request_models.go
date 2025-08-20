package monoacquiring

const (
	PaymentTypeDebit = "debit"
	PaymentTypeHold  = "hold"

	DisplayTypeIframe = "iframe"

	DiscountTypeDiscount    = "DISCOUNT"
	DiscountTypeExtraCharge = "EXTRA_CHARGE"

	DiscountModePercent = "PERCENT"
	DiscountModeValue   = "VALUE"

	SyncPaymentCardTypeFPAN = "FPAN"
	SyncPaymentCardTypeDPAN = "DPAN"

	InitiationKindClient   = "client"
	InitiationKindMerchant = "merchant"
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
	Destination    *string       `json:"destination,omitempty" validate:"omitempty,max=280"`
	Comment        *string       `json:"comment,omitempty" validate:"omitempty,max=280"`
	CustomerEmails []string      `json:"customerEmails" validate:"dive,email"`
	Discounts      []Discount    `json:"discounts,omitempty" validate:"dive"`
	BasketOrder    []BasketOrder `json:"basketOrder,omitempty" validate:"dive"`
}

type GooglePay struct {
	Cryptogram   *string `json:"cryptogram" validate:"omitempty"`
	Token        string  `json:"token" validate:"required"`
	Expiration   string  `json:"exp" validate:"required,card_exp"`
	EciIndicator string  `json:"eciIndicator" validate:"required"`
}

type ApplePay struct {
	Cryptogram   *string `json:"cryptogram" validate:"omitempty"`
	Token        string  `json:"token" validate:"required"`
	Expiration   string  `json:"exp" validate:"required,card_exp"`
	EciIndicator string  `json:"eciIndicator" validate:"required"`
}

type SyncPaymentCard struct {
	CVV              *string `json:"cvv" validate:"omitempty"`
	CAVV             *string `json:"cavv" validate:"omitempty"`
	TAVV             *string `json:"tavv" validate:"omitempty"`
	DSTranID         *string `json:"dsTranId" validate:"omitempty"`
	TokenRequestorID *string `json:"tReqID" validate:"omitempty"`
	MIT              *string `json:"mit" validate:"omitempty"`
	SST              *string `json:"sst" validate:"omitempty"`
	TraceID          *string `json:"tid" validate:"omitempty"`
	PAN              string  `json:"pan" validate:"required"`
	Type             string  `json:"type" validate:"required,oneof=FPAN DPAN"`
	Expiration       string  `json:"exp" validate:"required,card_exp"`
	EciIndicator     string  `json:"eciIndicator" validate:"required"`
}

type DirectPaymentCard struct {
	PAN        string `json:"pan" validate:"required"`
	Expiration string `json:"exp" validate:"required,card_exp"`
	CVV        string `json:"cvv" validate:"required"`
}
