//
// https://monobank.ua/api-docs/acquiring/metody/qr-ekvairynh/get--api--merchant--qr--details
//

package monoacquiring

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

type GetQrDetailsRequest struct {
	QrID string `validate:"required"`
}

type GetQrDetailsResponse struct {
	ShortQrID string `json:"shortQrId"`
	InvoiceID string `json:"invoiceId"`
	Amount    int    `json:"amount"`
	Currency  int    `json:"ccy"`
}

func (c *Client) GetQRDetails(ctx context.Context, payload GetQrDetailsRequest) (*GetQrDetailsResponse, error) {
	err := c.validator.StructCtx(ctx, payload)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	query := make(map[string]string, 1)
	query["qrId"] = payload.QrID

	req, err := c.newRequest(ctx, http.MethodGet, getQrDetailsPath, query, nil)
	if err != nil {
		return nil, err
	}

	var result GetQrDetailsResponse

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
