//
// https://monobank.ua/api-docs/acquiring/dodatkova-funktsionalnist/prro/get--api--merchant--invoice--fiscal-checks
//

package monoacquiring

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

type GetFiscalChecksRequest struct {
	InvoiceID string `validate:"required"`
}

type GetFiscalChecksResponse struct {
	Checks []struct {
		StatusDescription   *string           `json:"statusDescription"`
		TaxURL              *string           `json:"taxUrl"`
		File                *string           `json:"file"`
		ID                  string            `json:"id"`
		Type                FiscalCheckType   `json:"type"`
		Status              FiscalCheckStatus `json:"status"`
		FiscalizationSource FiscalCheckSource `json:"fiscalizationSource"`
	} `json:"checks"`
}

func (c *Client) GetFiscalChecks(ctx context.Context, payload GetFiscalChecksRequest) (*GetFiscalChecksResponse, error) {
	err := c.validator.StructCtx(ctx, payload)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	query := make(map[string]string, 2)
	query["invoiceId"] = payload.InvoiceID

	req, err := c.newRequest(ctx, http.MethodGet, geFiscalChecksPath, query, nil)
	if err != nil {
		return nil, err
	}

	var result GetFiscalChecksResponse

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
