//
// https://monobank.ua/api-docs/acquiring/metody/rozshcheplennia/get--api--merchant--statement
// https://monobank.ua/api-docs/acquiring/dodatkova-funktsionalnist/vypyska/get--api--merchant--statement
// https://monobank.ua/api-docs/acquiring/intehratory/marketpleisy-ta-ahenstka-skhema/get--api--merchant--statement
//

package monoacquiring

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"git.kbyte.app/mono/sdk/mono-acquiring-go/util"
	"github.com/pkg/errors"
)

type GetStatementRequest struct {
	To   *time.Time `validate:"omitempty"`
	Code *string    `validate:"omitempty"`
	From time.Time  `validate:"required"`
}

type GetStatementResponse struct {
	List []Statement `json:"list"`
}

func (c *Client) GetStatement(ctx context.Context, payload GetStatementRequest) (*GetStatementResponse, error) {
	err := c.validator.StructCtx(ctx, payload)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	query := make(map[string]string, 3)
	query["from"] = strconv.FormatInt(payload.From.Unix(), 10)

	if payload.Code != nil {
		query["code"] = util.PointerValue(payload.Code)
	}

	if payload.To != nil {
		query["to"] = strconv.FormatInt(util.PointerValue(payload.To).Unix(), 10)
	}

	req, err := c.newRequest(ctx, http.MethodGet, getStatementPath, query, nil)
	if err != nil {
		return nil, err
	}

	var result GetStatementResponse

	if err = c.doReq(req, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
