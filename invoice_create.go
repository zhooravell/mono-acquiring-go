//
// https://monobank.ua/api-docs/acquiring/metody/internet-ekvairynh/post--api--merchant--invoice--create
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
	Amount              int64                `json:"amount" validate:"required"`
	Ccy                 *int64               `json:"ccy,omitempty"`
	MerchantPaymentInfo *MerchantPaymentInfo `json:"merchantPaymInfo,omitempty"`
	RedirectURL         *string              `json:"redirectUrl,omitempty" validate:"omitempty,http_url"`
	WebHookURL          *string              `json:"webHookUrl,omitempty" validate:"omitempty,http_url"`
	Validity            *int64               `json:"validity,omitempty"`
	PaymentType         string               `json:"paymentType" validate:"required,oneof=debit hold"`
	QrID                *string              `json:"qrId"`
	Code                *string              `json:"code"`
	SaveCardData        *SaveCardData        `json:"saveCardData,omitempty"`
	AgentFeePercent     *float64             `json:"agentFeePercent,omitempty"`
	TipsEmployeeID      *string              `json:"tipsEmployeeId,omitempty"`
	DisplayType         *string              `json:"displayType,omitempty"`
}

type InvoiceCreateResponse struct {
	InvoiceID string `json:"invoiceId"`
	PageURL   string `json:"pageUrl"`
}

func (c *Client) CreateInvoice(ctx context.Context, payload InvoiceCreateRequest) (*InvoiceCreateResponse, error) {
	var (
		err     error
		req     *http.Request
		reqBody []byte
		result  *InvoiceCreateResponse

		path = "/api/merchant/invoice/create"
	)

	if payload.PaymentType == "" {
		payload.PaymentType = PaymentTypeDebit
	}

	if err = c.validator.StructCtx(ctx, payload); err != nil {
		return nil, errors.WithStack(err)
	}

	if reqBody, err = json.Marshal(payload); err != nil {
		return nil, errors.WithStack(err)
	}

	if req, err = c.newRequest(ctx, http.MethodPost, path, nil, bytes.NewBuffer(reqBody)); err != nil {
		return nil, err
	}

	if err = c.doReq(req, result); err != nil {
		return nil, errors.WithStack(err)
	}

	return result, nil
}
