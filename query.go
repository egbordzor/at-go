package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
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

// Fetch transactions of a particular payment product.
func (c *Client) FindTransaction(ctx context.Context, p *africastalking.FindTransactionPayload) (res *africastalking.FindTransactionResponse, err error) {

	req, err := c.NewRequest("GET", "https://payments.sandbox.africastalking.com/query/transaction/find", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &africastalking.FindTransactionResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Fetch transactions of a particular payment product.
func (c *Client) FetchProductTransactions(ctx context.Context, p *africastalking.ProductTransactionsPayload) (res *africastalking.ProductTransactionsResponse, err error) {

	req, err := c.NewRequest("GET", "https://payments.sandbox.africastalking.com/query/transaction/fetch", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &africastalking.ProductTransactionsResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Fetch your wallet transactions
func (c *Client) FetchWalletTransactions(ctx context.Context, p *africastalking.WalletTransactionsPayload) (res *africastalking.WalletTransactionsResponse, err error) {

	req, err := c.NewRequest("GET", "https://payments.sandbox.africastalking.com/query/wallet/fetch", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &africastalking.WalletTransactionsResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Fetch your wallet balance
func (c *Client) FetchWalletBalance(ctx context.Context, p *africastalking.WalletBalancePayload) (res *africastalking.WalletBalanceResponse, err error) {

	req, err := c.NewRequest("GET", "https://payments.sandbox.africastalking.com/query/wallet/balance", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &africastalking.WalletBalanceResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
