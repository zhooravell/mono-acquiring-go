package monoacquiring

type TipsInfo struct {
	EmployeeID string `json:"employeeId"`
	Amount     int    `json:"amount"`
}

type WalletData struct {
	CardToken string `json:"cardToken"`
	WalletID  string `json:"walletId"`
	Status    string `json:"status"`
}

type PaymentInfo struct {
	MaskedPan     string `json:"maskedPan"`
	ApprovalCode  string `json:"approvalCode"`
	RRN           string `json:"rrn"`
	TransactionID string `json:"tranId"`
	Terminal      string `json:"terminal"`
	Bank          string `json:"bank"`
	PaymentSystem string `json:"paymentSystem"`
	PaymentMethod string `json:"paymentMethod"`
	Country       string `json:"country"`
	Fee           int64  `json:"fee"`
	AgentFee      int64  `json:"agentFee"`
}

type CancelListItem struct {
	Status            string `json:"status"`
	CreatedDate       string `json:"createdDate"`
	ModifiedDate      string `json:"modifiedDate"`
	ApprovalCode      string `json:"approvalCode"`
	RRN               string `json:"rrn"`
	ExternalReference string `json:"extRef"`
	Amount            int64  `json:"amount"`
	Currency          int    `json:"ccy"`
}

const (
	qrAmountTypeMerchant = "merchant"
	qrAmountTypeClient   = "client"
	qrAmountTypeFix      = "fix"
)

type QRAmountType string

func (qat QRAmountType) String() string {
	return string(qat)
}

func (qat QRAmountType) IsMerchant() bool {
	return qat.String() == qrAmountTypeMerchant
}

func (qat QRAmountType) IsClient() bool {
	return qat.String() == qrAmountTypeClient
}

func (qat QRAmountType) IsFix() bool {
	return qat.String() == qrAmountTypeFix
}
