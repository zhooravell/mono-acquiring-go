package monoacquiring

type TipsInfo struct {
	EmployeeID string `json:"employeeId"`
	Amount     int    `json:"amount"`
}

const (
	walletDataStatusNew     = "new"
	walletDataStatusCreated = "created"
	walletDataStatusFailed  = "failed"
)

type WalletDataStatus string

func (wds WalletDataStatus) String() string {
	return string(wds)
}

func (wds WalletDataStatus) IsNew() bool {
	return wds.String() == walletDataStatusNew
}

func (wds WalletDataStatus) IsCreated() bool {
	return wds.String() == walletDataStatusCreated
}

func (wds WalletDataStatus) IsFailed() bool {
	return wds.String() == walletDataStatusFailed
}

type WalletData struct {
	CardToken string           `json:"cardToken"`
	WalletID  string           `json:"walletId"`
	Status    WalletDataStatus `json:"status"`
}

const (
	paymentSystemMasterCard = "mastercard"
	paymentSystemVisa       = "visa"
)

type PaymentSystem string

func (ps PaymentSystem) String() string {
	return string(ps)
}

func (ps PaymentSystem) IsMasterCard() bool {
	return ps.String() == paymentSystemMasterCard
}

func (ps PaymentSystem) IsVisa() bool {
	return ps.String() == paymentSystemVisa
}

const (
	paymentMethodPAN      = "pan"
	paymentMethodApple    = "apple"
	paymentMethodGoogle   = "google"
	paymentMethodMonobank = "monobank"
	paymentMethodWallet   = "wallet"
	paymentMethodDirect   = "direct"
)

type PaymentMethod string

func (pm PaymentMethod) String() string {
	return string(pm)
}

func (pm PaymentMethod) IsPAN() bool {
	return pm.String() == paymentMethodPAN
}

func (pm PaymentMethod) IsApple() bool {
	return pm.String() == paymentMethodApple
}

func (pm PaymentMethod) IsGoogle() bool {
	return pm.String() == paymentMethodGoogle
}

func (pm PaymentMethod) IsMonobank() bool {
	return pm.String() == paymentMethodMonobank
}

func (pm PaymentMethod) IsWallet() bool {
	return pm.String() == paymentMethodWallet
}

func (pm PaymentMethod) IsDirect() bool {
	return pm.String() == paymentMethodDirect
}

type PaymentInfo struct {
	MaskedPan     string        `json:"maskedPan"`
	ApprovalCode  string        `json:"approvalCode"`
	RRN           string        `json:"rrn"`
	TransactionID string        `json:"tranId"`
	Terminal      string        `json:"terminal"`
	Bank          string        `json:"bank"`
	PaymentSystem PaymentSystem `json:"paymentSystem"`
	PaymentMethod PaymentMethod `json:"paymentMethod"`
	Country       string        `json:"country"`
	Fee           int64         `json:"fee"`
	AgentFee      int64         `json:"agentFee"`
}

const (
	cancelListItemStatusProcessing = "processing"
	cancelListItemStatusSuccess    = "success"
	cancelListItemStatusFailure    = "failure"
)

type CancelListItemStatus string

func (cli CancelListItemStatus) String() string {
	return string(cli)
}

func (cli CancelListItemStatus) IsProcessing() bool {
	return cli.String() == cancelListItemStatusProcessing
}

func (cli CancelListItemStatus) IsSuccess() bool {
	return cli.String() == cancelListItemStatusSuccess
}

func (cli CancelListItemStatus) IsFailure() bool {
	return cli.String() == cancelListItemStatusFailure
}

type CancelListItem struct {
	Status            CancelListItemStatus `json:"status"`
	CreatedDate       string               `json:"createdDate"`
	ModifiedDate      string               `json:"modifiedDate"`
	ApprovalCode      string               `json:"approvalCode"`
	RRN               string               `json:"rrn"`
	ExternalReference string               `json:"extRef"`
	Amount            int64                `json:"amount"`
	Currency          int                  `json:"ccy"`
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

const (
	fiscalCheckStatusNew     = "new"
	fiscalCheckStatusProcess = "process"
	fiscalCheckStatusDone    = "done"
	fiscalCheckStatusFailed  = "failed"
)

type FiscalCheckStatus string

func (fcs FiscalCheckStatus) String() string {
	return string(fcs)
}

func (fcs FiscalCheckStatus) IsNew() bool {
	return fcs.String() == fiscalCheckStatusNew
}

func (fcs FiscalCheckStatus) IsProcess() bool {
	return fcs.String() == fiscalCheckStatusProcess
}

func (fcs FiscalCheckStatus) IsDone() bool {
	return fcs.String() == fiscalCheckStatusDone
}

func (fcs FiscalCheckStatus) IsFailed() bool {
	return fcs.String() == fiscalCheckStatusFailed
}

const (
	fiscalCheckTypeSale   = "sale"
	fiscalCheckTypeReturn = "return"
)

type FiscalCheckType string

func (fct FiscalCheckType) String() string {
	return string(fct)
}

func (fct FiscalCheckType) IsSale() bool {
	return fct.String() == fiscalCheckTypeSale
}

func (fct FiscalCheckType) IsReturn() bool {
	return fct.String() == fiscalCheckTypeReturn
}

const (
	fiscalCheckSourceCheckBox    = "checkbox"
	fiscalCheckSourceMonoPay     = "monopay"
	fiscalCheckSourceVchasnoKasa = "vchasnokasa"
)

type FiscalCheckSource string

func (fcs FiscalCheckSource) String() string {
	return string(fcs)
}

func (fcs FiscalCheckSource) IsCheckBox() bool {
	return fcs.String() == fiscalCheckSourceCheckBox
}

func (fcs FiscalCheckSource) IsMonoPay() bool {
	return fcs.String() == fiscalCheckSourceMonoPay
}

func (fcs FiscalCheckSource) IsVchasnoKasa() bool {
	return fcs.String() == fiscalCheckSourceVchasnoKasa
}

const (
	statementStatusHold       = "hold"
	statementStatusProcessing = "processing"
	statementStatusSuccess    = "success"
	statementStatusFailure    = "failure"
)

type StatementStatus string

func (ss StatementStatus) String() string {
	return string(ss)
}

func (ss StatementStatus) IsHold() bool {
	return ss.String() == statementStatusHold
}

func (ss StatementStatus) IsProcessing() bool {
	return ss.String() == statementStatusProcessing
}

func (ss StatementStatus) IsSuccess() bool {
	return ss.String() == statementStatusSuccess
}

func (ss StatementStatus) IsFailure() bool {
	return ss.String() == statementStatusFailure
}

const (
	statementPaymentSchemeBnplLater30 = "bnpl_later_30"
	statementPaymentSchemeBnplParts4  = "bnpl_parts_4"
	statementPaymentSchemeFull        = "full"
)

type StatementPaymentScheme string

func (ssp StatementPaymentScheme) String() string {
	return string(ssp)
}

func (ssp StatementPaymentScheme) IsBnplLater30() bool {
	return ssp.String() == statementPaymentSchemeBnplLater30
}

func (ssp StatementPaymentScheme) IsBnplParts4() bool {
	return ssp.String() == statementPaymentSchemeBnplParts4
}

func (ssp StatementPaymentScheme) IsFull() bool {
	return ssp.String() == statementPaymentSchemeFull
}

type StatementCancel struct {
	ApprovalCode *string `json:"approvalCode,omitempty"`
	RRN          *string `json:"rrn,omitempty"`
	MaskedPan    string  `json:"maskedPan"`
	Date         string  `json:"date"`
	Amount       int64   `json:"amount"`
	Currency     int     `json:"ccy"`
}

type Statement struct {
	InvoiceID     string                 `json:"invoiceId"`
	MaskedPan     string                 `json:"maskedPan"`
	Date          string                 `json:"date"`
	Status        StatementStatus        `json:"status"`
	PaymentScheme StatementPaymentScheme `json:"paymentScheme"`
	ApprovalCode  *string                `json:"approvalCode,omitempty"`
	RRN           *string                `json:"rrn,omitempty"`
	Reference     *string                `json:"reference,omitempty"`
	ShortQrID     *string                `json:"shortQrId,omitempty"`
	Destination   *string                `json:"destination,omitempty"`
	ProfitAmount  *int64                 `json:"profitAmount,omitempty"`
	CancelList    []StatementCancel      `json:"cancelList,omitempty"`
	Amount        int64                  `json:"amount"`
	Currency      int                    `json:"ccy"`
}

const (
	holdStatusSuccess = "success"
)

type HoldFinalizationStatus string

func (hfs HoldFinalizationStatus) String() string {
	return string(hfs)
}

func (hfs HoldFinalizationStatus) IsSuccess() bool {
	return hfs.String() == holdStatusSuccess
}

const (
	SyncPaymentStatusCreated    = "created"
	SyncPaymentStatusProcessing = "processing"
	SyncPaymentStatusHold       = "hold"
	SyncPaymentStatusSuccess    = "success"
	SyncPaymentStatusFailure    = "failure"
	SyncPaymentStatusReversed   = "reversed"
	SyncPaymentStatusExpired    = "expired"
)

type SyncPaymentStatus string

func (sps SyncPaymentStatus) String() string {
	return string(sps)
}

func (sps SyncPaymentStatus) IsCreated() bool {
	return sps.String() == SyncPaymentStatusCreated
}

func (sps SyncPaymentStatus) IsProcessing() bool {
	return sps.String() == SyncPaymentStatusProcessing
}

func (sps SyncPaymentStatus) IsHold() bool {
	return sps.String() == SyncPaymentStatusHold
}

func (sps SyncPaymentStatus) IsSuccess() bool {
	return sps.String() == SyncPaymentStatusSuccess
}

func (sps SyncPaymentStatus) IsFailure() bool {
	return sps.String() == SyncPaymentStatusFailure
}

func (sps SyncPaymentStatus) IsReversed() bool {
	return sps.String() == SyncPaymentStatusReversed
}

func (sps SyncPaymentStatus) IsExpired() bool {
	return sps.String() == SyncPaymentStatusExpired
}

const (
	tokenPaymentStatusProcessing = "processing"
	tokenPaymentStatusSuccess    = "success"
	tokenPaymentStatusFailure    = "failure"
)

type TokenPaymentStatus string

func (tps TokenPaymentStatus) String() string {
	return string(tps)
}

func (tps TokenPaymentStatus) IsProcessing() bool {
	return tps.String() == tokenPaymentStatusProcessing
}

func (tps TokenPaymentStatus) IsSuccess() bool {
	return tps.String() == tokenPaymentStatusSuccess
}

func (tps TokenPaymentStatus) IsFailure() bool {
	return tps.String() == tokenPaymentStatusFailure
}

const (
	directPaymentStatusProcessing = "processing"
	directPaymentStatusSuccess    = "success"
	directPaymentStatusFailure    = "failure"
)

type DirectPaymentStatus string

func (dps DirectPaymentStatus) String() string {
	return string(dps)
}

func (dps DirectPaymentStatus) IsProcessing() bool {
	return dps.String() == directPaymentStatusProcessing
}

func (dps DirectPaymentStatus) IsSuccess() bool {
	return dps.String() == directPaymentStatusSuccess
}

func (dps DirectPaymentStatus) IsFailure() bool {
	return dps.String() == directPaymentStatusFailure
}
