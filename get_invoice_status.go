//
// https://monobank.ua/api-docs/acquiring/metody/internet-ekvairynh/get--api--merchant--invoice--status
// https://monobank.ua/api-docs/acquiring/dodatkova-funktsionalnist/tokenizatsiia/get--api--merchant--invoice--status
// https://monobank.ua/api-docs/acquiring/metody/rozshcheplennia/get--api--merchant--invoice--status
//

package monoacquiring

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

type GetInvoiceStatusRequest struct {
	InvoiceID string `json:"invoiceId" validate:"required"`
}

type GetInvoiceStatusResponse struct {
	Destination   *string          `json:"destination,omitempty"`
	TipsInfo      *TipsInfo        `json:"tipsInfo,omitempty"`
	FinalAmount   *int             `json:"finalAmount,omitempty"`
	CreatedDate   *string          `json:"createdDate,omitempty"`
	ModifiedDate  *string          `json:"modifiedDate,omitempty"`
	Reference     *string          `json:"reference,omitempty"`
	ErrCode       *string          `json:"errCode,omitempty"`
	PaymentInfo   *PaymentInfo     `json:"paymentInfo"`
	FailureReason *string          `json:"failureReason,omitempty"`
	WalletData    *WalletData      `json:"walletData,omitempty"`
	InvoiceID     string           `json:"invoiceId"`
	Status        string           `json:"status"`
	CancelList    []CancelListItem `json:"cancelList,omitempty"`
	Amount        int64            `json:"amount"`
	Currency      int              `json:"ccy"`
}

func (c *Client) GetInvoiceStatus(
	ctx context.Context,
	payload GetInvoiceStatusRequest,
) (*GetInvoiceStatusResponse, error) {
	err := c.validator.StructCtx(ctx, payload)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	query := make(map[string]string, 1)
	query["invoiceId"] = payload.InvoiceID

	req, err := c.newRequest(ctx, http.MethodGet, invoiceStatusPath, query, nil)
	if err != nil {
		return nil, err
	}

	var result GetInvoiceStatusResponse
	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
