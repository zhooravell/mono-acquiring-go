package monoacquiring

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type CancelInvoiceItem struct {
	Barcode *string `json:"barcode,omitempty"`
	Header  *string `json:"header,omitempty"`
	Footer  *string `json:"footer,omitempty"`
	Uktzed  *string `json:"uktzed,omitempty"`
	Name    string  `json:"name" validate:"required"`
	Code    string  `json:"code" validate:"required"`
	Tax     []int64 `json:"tax,omitempty"`
	Qty     int     `json:"qty" validate:"required"`
	Sum     int64   `json:"sum" validate:"required"`
}

type CancelInvoiceRequest struct {
	InvoiceID string              `json:"invoiceId" validate:"required"`
	ExtRef    *string             `json:"extRef,omitempty"`
	Amount    *int64              `json:"amount,omitempty"`
	Items     []CancelInvoiceItem `json:"items"`
}

type CancelInvoiceResponse struct {
	Status       string `json:"status"`
	CreatedDate  string `json:"createdDate"`
	ModifiedDate string `json:"modifiedDate"`
}

func (c *Client) CancelInvoice(ctx context.Context, payload CancelInvoiceRequest) (*CancelInvoiceResponse, error) {
	var (
		err     error
		req     *http.Request
		reqBody []byte
		result  CancelInvoiceResponse

		path = "/api/merchant/invoice/cancel"
	)

	if err = c.validator.StructCtx(ctx, payload); err != nil {
		return nil, errors.WithStack(err)
	}

	if reqBody, err = json.Marshal(payload); err != nil {
		return nil, errors.Wrap(err, "failed to marshal cancel invoice request")
	}

	if req, err = c.newRequest(ctx, http.MethodPost, path, nil, bytes.NewBuffer(reqBody)); err != nil {
		return nil, err
	}

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
