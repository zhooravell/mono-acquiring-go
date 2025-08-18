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
	Rrn           string `json:"rrn"`
	TranID        string `json:"tranId"`
	Terminal      string `json:"terminal"`
	Bank          string `json:"bank"`
	PaymentSystem string `json:"paymentSystem"`
	PaymentMethod string `json:"paymentMethod"`
	Country       string `json:"country"`
	Fee           int64  `json:"fee"`
	AgentFee      int64  `json:"agentFee"`
}

type CancelListItem struct {
	Status       string `json:"status"`
	CreatedDate  string `json:"createdDate"`
	ModifiedDate string `json:"modifiedDate"`
	ApprovalCode string `json:"approvalCode"`
	Rrn          string `json:"rrn"`
	ExtRef       string `json:"extRef"`
	Amount       int64  `json:"amount"`
	Ccy          int    `json:"ccy"`
}
