package atgo

import (
	"context"
	"fmt"
	"github.com/hashicorp/errwrap"
	pay "github.com/wondenge/at-go/payments"
	"go.uber.org/zap"
)

// Fetch transactions of a particular payment product.
func (c *Client) FindTransaction(ctx context.Context, p *pay.FindTransactionPayload) (res *pay.FindTransactionResponse, err error) {

	if err := c.requestJSONBody(ctx, "GET", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/query/transaction/find"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Fetch transactions of a particular payment product.
func (c *Client) FetchProductTransactions(ctx context.Context, p *pay.ProductTransactionsPayload) (res *pay.ProductTransactionsResponse, err error) {

	if err := c.requestJSONBody(ctx, "GET", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/query/transaction/fetch"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Fetch your wallet transactions
func (c *Client) FetchWalletTransactions(ctx context.Context, p *pay.WalletTransactionsPayload) (res *pay.WalletTransactionsResponse, err error) {

	if err := c.requestJSONBody(ctx, "GET", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/query/wallet/fetch"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Fetch your wallet balance
func (c *Client) FetchWalletBalance(ctx context.Context, p *pay.WalletBalancePayload) (res *pay.WalletBalanceResponse, err error) {

	if err := c.requestJSONBody(ctx, "GET", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/query/wallet/balance"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}
