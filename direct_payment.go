//
// https://monobank.ua/api-docs/acquiring/metody/internet-ekvairynh/post--api--merchant--invoice--payment-direct
//

package monoacquiring

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type DirectPaymentRequest struct {
	MerchantPaymentInfo *MerchantPaymentInfo `json:"merchantPaymInfo,omitempty" validate:"omitempty"`
	SaveCardData        *SaveCardData        `json:"saveCardData,omitempty" validate:"omitempty"`
	Currency            *int                 `json:"ccy" validate:"omitempty,iso4217_numeric"`
	InitiationKind      *string              `json:"initiationKind" validate:"omitempty,oneof=merchant client"`
	Card                DirectPaymentCard    `json:"cardData" validate:"required"`
	PaymentType         string               `json:"paymentType" validate:"required,oneof=debit hold"`
	Amount              int64                `json:"amount" validate:"required"`
}

type DirectPaymentResponse struct {
	InvoiceID     string              `json:"invoiceId"`
	TdsURL        string              `json:"tdsUrl"`
	Status        DirectPaymentStatus `json:"status"`
	FailureReason string              `json:"failureReason"`
	CreatedDate   string              `json:"createdDate"`
	ModifiedDate  string              `json:"modifiedDate"`
	Amount        int                 `json:"amount"`
	Ccy           int                 `json:"ccy"`
}

func (c *Client) DirectPayment(ctx context.Context, payload DirectPaymentRequest) (*DirectPaymentResponse, error) {
	if payload.PaymentType == "" {
		payload.PaymentType = PaymentTypeDebit
	}

	err := c.validator.StructCtx(ctx, payload)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	if err = json.NewEncoder(buf).Encode(payload); err != nil {
		return nil, errors.Wrap(err, "failed to marshal direct payment request")
	}

	req, err := c.newRequest(ctx, http.MethodPost, directPaymentPath, nil, buf)
	if err != nil {
		return nil, err
	}

	var result DirectPaymentResponse

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
