package monoacquiring

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type SyncPaymentRequest struct {
	MerchantPaymentInfo *MerchantPaymentInfo `json:"merchantPaymInfo,omitempty" validate:"omitempty"`
	GooglePay           *GooglePay           `json:"googlePay,omitempty" validate:"omitempty"`
	ApplePay            *ApplePay            `json:"applePay,omitempty" validate:"omitempty"`
	SyncPaymentCard     *SyncPaymentCard     `json:"cardData,omitempty" validate:"omitempty"`
	Currency            int                  `json:"ccy" validate:"required,iso4217_numeric"`
	Amount              int64                `json:"amount" validate:"required"`
}

type SyncPaymentResponse struct {
	FinalAmount   *int64            `json:"finalAmount,omitempty"`
	CreatedDate   *string           `json:"createdDate,omitempty"`
	FailureReason *string           `json:"failureReason,omitempty"`
	ErrCode       *string           `json:"errCode,omitempty"`
	TipsInfo      *TipsInfo         `json:"tipsInfo,omitempty"`
	WalletData    *WalletData       `json:"walletData,omitempty"`
	PaymentInfo   *PaymentInfo      `json:"paymentInfo,omitempty"`
	Destination   *string           `json:"destination,omitempty"`
	ModifiedDate  *string           `json:"modifiedDate,omitempty"`
	Reference     *string           `json:"reference,omitempty"`
	InvoiceID     string            `json:"invoiceId"`
	Status        SyncPaymentStatus `json:"status"`
	CancelList    []CancelListItem  `json:"cancelList"`
	Ccy           int               `json:"ccy"`
	Amount        int64             `json:"amount"`
}

func (c *Client) SyncPayment(ctx context.Context, payload SyncPaymentRequest) (*SyncPaymentResponse, error) {
	err := c.validator.StructCtx(ctx, payload)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	if err = json.NewEncoder(buf).Encode(payload); err != nil {
		return nil, errors.Wrap(err, "failed to marshal sync payment request")
	}

	req, err := c.newRequest(ctx, http.MethodPost, syncPaymentPath, nil, buf)
	if err != nil {
		return nil, err
	}

	var result SyncPaymentResponse

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
