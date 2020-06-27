package atgo

import (
	"context"
	"fmt"
	at "github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

const (
	FindTransactionLiveURL    = "https://payments.africastalking.com/query/transaction/find"
	FindTransactionSandboxURL = "https://payments.sandbox.africastalking.com/query/transaction/find"

	FetchProductTransactionsLiveURL    = "https://payments.africastalking.com/query/transaction/fetch"
	FetchProductTransactionsSandboxURL = "https://payments.sandbox.africastalking.com/query/transaction/fetch"

	FetchWalletTransactionsLiveURL    = "https://payments.africastalking.com/query/wallet/fetch"
	FetchWalletTransactionsSandboxURL = "https://payments.sandbox.africastalking.com/query/wallet/fetch"

	FetchWalletBalanceLiveURL    = "https://payments.africastalking.com/query/wallet/balance"
	FetchWalletBalanceSandboxURL = "https://payments.sandbox.africastalking.com/query/wallet/balance"
)

type (
	PaymentQueries interface {

		// Fetch transactions of a particular payment product.
		findTransaction(ctx context.Context, p *at.FindTransactionPayload) (res *at.FindTransactionResponse, err error)
		// Fetch transactions of a particular payment product.
		fetchProductTransactions(ctx context.Context, p *at.ProductTransactionsPayload) (res *at.ProductTransactionsResponse, err error)

		// Fetch your wallet transactions
		fetchWalletTransactions(ctx context.Context, p *at.WalletTransactionsPayload) (res *at.WalletTransactionsResponse, err error)
		// Fetch your wallet balance
		fetchWalletBalance(ctx context.Context, p *at.WalletBalancePayload) (res *at.WalletBalanceResponse, err error)
	}
)

// Fetch transactions of a particular payment product.
func (c *Client) findTransaction(ctx context.Context, p *at.FindTransactionPayload) (res *at.FindTransactionResponse, err error) {

	req, err := c.NewRequest("GET", "https://payments.sandbox.africastalking.com/query/transaction/find", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &at.FindTransactionResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Fetch transactions of a particular payment product.
func (c *Client) fetchProductTransactions(ctx context.Context, p *at.ProductTransactionsPayload) (res *at.ProductTransactionsResponse, err error) {

	req, err := c.NewRequest("GET", "https://payments.sandbox.africastalking.com/query/transaction/fetch", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &at.ProductTransactionsResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Fetch your wallet transactions
func (c *Client) fetchWalletTransactions(ctx context.Context, p *at.WalletTransactionsPayload) (res *at.WalletTransactionsResponse, err error) {

	req, err := c.NewRequest("GET", "https://payments.sandbox.africastalking.com/query/wallet/fetch", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &at.WalletTransactionsResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Fetch your wallet balance
func (c *Client) fetchWalletBalance(ctx context.Context, p *at.WalletBalancePayload) (res *at.WalletBalanceResponse, err error) {

	req, err := c.NewRequest("GET", "https://payments.sandbox.africastalking.com/query/wallet/balance", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &at.WalletBalanceResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
