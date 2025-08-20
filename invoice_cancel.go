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
	InvoiceID         string              `json:"invoiceId" validate:"required"`
	ExternalReference *string             `json:"extRef,omitempty"`
	Amount            *int64              `json:"amount,omitempty"`
	Items             []CancelInvoiceItem `json:"items"`
}

type CancelInvoiceResponse struct {
	Status       string `json:"status"`
	CreatedDate  string `json:"createdDate"`
	ModifiedDate string `json:"modifiedDate"`
}

func (c *Client) CancelInvoice(ctx context.Context, payload CancelInvoiceRequest) (*CancelInvoiceResponse, error) {
	err := c.validator.StructCtx(ctx, payload)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal cancel invoice request")
	}

	req, err := c.newRequest(ctx, http.MethodPost, invoiceCancelPath, nil, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	var result CancelInvoiceResponse

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
