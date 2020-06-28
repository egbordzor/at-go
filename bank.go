package atgo

import (
	"context"
	"fmt"
	at "github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

type (
	Bank interface {

		// Collect money into your payment wallet.
		bankCheckout(ctx context.Context, p *at.BankCheckoutPayload) (res *at.BankCheckoutResponse, err error)

		// Validate a bank checkout charge request
		bankCheckoutValidate(ctx context.Context, p *at.BankCheckoutValidatePayload) (res *at.BankCheckoutValidateResponse, err error)

		// Initiate a bank transfer request.
		bankTransfer(ctx context.Context, p *at.BankTransferPayload) (res *at.BankTransferResponse, err error)
	}
)

// Collect money into your payment wallet.
func (c *ATClient) bankCheckout(ctx context.Context, p *at.BankCheckoutPayload) (res *at.BankCheckoutResponse, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/bank/checkout/charge"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.BankCheckoutResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Validate a bank checkout charge request
func (c *ATClient) bankCheckoutValidate(ctx context.Context, p *at.BankCheckoutValidatePayload) (res *at.BankCheckoutValidateResponse, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/bank/checkout/validate"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.BankCheckoutValidateResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Initiate a bank transfer request.
func (c *ATClient) bankTransfer(ctx context.Context, p *at.BankTransferPayload) (res *at.BankTransferResponse, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/bank/transfer"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.BankTransferResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
