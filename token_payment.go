//
// https://monobank.ua/api-docs/acquiring/dodatkova-funktsionalnist/tokenizatsiia/post--api--merchant--wallet--payment
//

package monoacquiring

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type TokenPaymentRequest struct {
	RedirectURL         *string              `json:"redirectUrl,omitempty" validate:"omitempty,http_url"`
	WebHookURL          *string              `json:"webHookUrl,omitempty" validate:"omitempty,http_url"`
	MerchantPaymentInfo *MerchantPaymentInfo `json:"merchantPaymInfo,omitempty" validate:"omitempty"`
	CardToken           string               `json:"cardToken" validate:"required"`
	InitiationKind      string               `json:"initiationKind" validate:"required,oneof=merchant client"`
	PaymentType         string               `json:"paymentType" validate:"required,oneof=debit hold"`
	Currency            int                  `json:"ccy" validate:"required,iso4217_numeric"`
	Amount              int64                `json:"amount" validate:"required"`
}

type TokenPaymentResponse struct {
	InvoiceID     string             `json:"invoiceId"`
	TdsURL        string             `json:"tdsUrl"`
	Status        TokenPaymentStatus `json:"status"`
	FailureReason *string            `json:"failureReason,omitempty"`
	CreatedDate   string             `json:"createdDate"`
	ModifiedDate  string             `json:"modifiedDate"`
	Amount        int64              `json:"amount"`
	Currency      int                `json:"ccy"`
}

func (c *Client) TokenPayment(ctx context.Context, payload TokenPaymentRequest) (*TokenPaymentResponse, error) {
	if payload.PaymentType == "" {
		payload.PaymentType = PaymentTypeDebit
	}

	err := c.validator.StructCtx(ctx, payload)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	if err = json.NewEncoder(buf).Encode(payload); err != nil {
		return nil, errors.Wrap(err, "failed to marshal token payment request")
	}

	req, err := c.newRequest(ctx, http.MethodPost, tokenPaymentPath, nil, buf)
	if err != nil {
		return nil, err
	}

	var result TokenPaymentResponse

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
