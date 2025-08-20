//
// https://monobank.ua/api-docs/acquiring/metody/internet-ekvairynh/post--api--merchant--invoice--create
// https://monobank.ua/api-docs/acquiring/metody/qr-ekvairynh/post--api--merchant--invoice--create
//

package monoacquiring

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type InvoiceCreateRequest struct {
	SaveCardData        *SaveCardData        `json:"saveCardData,omitempty"`
	Currency            *int64               `json:"ccy,omitempty" validate:"omitempty,iso4217_numeric"`
	MerchantPaymentInfo *MerchantPaymentInfo `json:"merchantPaymInfo,omitempty"`
	RedirectURL         *string              `json:"redirectUrl,omitempty" validate:"omitempty,http_url"`
	WebHookURL          *string              `json:"webHookUrl,omitempty" validate:"omitempty,http_url"`
	Validity            *int64               `json:"validity,omitempty"`
	QrID                *string              `json:"qrId,omitempty"`
	Code                *string              `json:"code,omitempty"`
	AgentFeePercent     *float64             `json:"agentFeePercent,omitempty"`
	TipsEmployeeID      *string              `json:"tipsEmployeeId,omitempty"`
	DisplayType         *string              `json:"displayType,omitempty" validate:"omitempty,oneof=iframe"`
	PaymentType         string               `json:"paymentType" validate:"required,oneof=debit hold"`
	Amount              int64                `json:"amount" validate:"required"`
}

type InvoiceCreateResponse struct {
	InvoiceID string `json:"invoiceId"`
	PageURL   string `json:"pageUrl"`
}

func (c *Client) CreateInvoice(ctx context.Context, payload InvoiceCreateRequest) (*InvoiceCreateResponse, error) {
	if payload.PaymentType == "" {
		payload.PaymentType = PaymentTypeDebit
	}

	var err error

	if err = c.validator.StructCtx(ctx, payload); err != nil {
		return nil, errors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	if err = json.NewEncoder(buf).Encode(payload); err != nil {
		return nil, errors.Wrap(err, "failed to marshal create invoice request")
	}

	req, err := c.newRequest(ctx, http.MethodPost, invoiceCreatePath, nil, buf)
	if err != nil {
		return nil, err
	}

	var result InvoiceCreateResponse

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
