package monoacquiring

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type FinalizeHoldRequest struct {
	InvoiceID string        `json:"invoiceId" validate:"required"`
	Amount    *int64        `json:"amount,omitempty"`
	Items     []BasketOrder `json:"items,omitempty"`
}

type FinalizeHoldResponse struct {
	Status HoldFinalizationStatus `json:"status"`
}

func (c *Client) FinalizeHold(ctx context.Context, payload FinalizeHoldRequest) (*FinalizeHoldResponse, error) {
	err := c.validator.StructCtx(ctx, payload)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	buf := new(bytes.Buffer)
	if err = json.NewEncoder(buf).Encode(payload); err != nil {
		return nil, errors.Wrap(err, "failed to marshal finalize hold request")
	}

	req, err := c.newRequest(ctx, http.MethodPost, finalizeHoldPath, nil, buf)
	if err != nil {
		return nil, err
	}

	var result FinalizeHoldResponse

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
