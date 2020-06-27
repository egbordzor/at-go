package atgo

import (
	"context"
	"fmt"
	at "github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

const (
	MobileCheckoutLiveURL    = "https://payments.africastalking.com/mobile/checkout/request"
	MobileCheckoutSandboxURL = "https://payments.sandbox.africastalking.com/mobile/checkout/request"

	MobileB2CLiveURL = "https://payments.africastalking.com/mobile/b2c/request"
	MobileB2CSandbox = "https://payments.sandbox.africastalking.com/mobile/b2c/request"

	MobileB2BLiveURL    = "https://payments.africastalking.com/mobile/b2b/request"
	MobileB2BSandboxURL = "https://payments.sandbox.africastalking.com/mobile/b2b/request"
)

type (
	Mobile interface {
		// Initiate C2B payments on a mobile subscriber’s device.
		mobileCheckout(ctx context.Context, p *at.MobileCheckoutPayload) (res *at.MobileCheckoutResponse, err error)

		// Send payments to mobile subscribers from your Payment Wallet.
		mobileB2C(ctx context.Context, p *at.MobileB2CPayload) (res *at.MobileB2CResponse, err error)

		// Send payments to businesses e.g banks from your Payment Wallet.
		mobileB2B(ctx context.Context, p *at.MobileB2BPayload) (res at.MobileB2BResponse, err error)
	}
)

// Mobile Checkout APIs allow you to initiate Customer to Business (C2B) payments
// on a mobile subscriber’s device.
// This allows for a smoother checkout experience, since the client will no longer
// need to remember the amount or an account number to complete the transaction.
func (c *Client) mobileCheckout(ctx context.Context, p *at.MobileCheckoutPayload) (res *at.MobileCheckoutResponse, err error) {

	req, err := c.NewRequest("POST", "https://payments.sandbox.africastalking.com/mobile/checkout/request", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &at.MobileCheckoutResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Mobile Business To Consumer (B2C) APIs allow you to send payments to mobile
// subscribers from your Payment Wallet.
func (c *Client) mobileB2C(ctx context.Context, p *at.MobileB2CPayload) (res *at.MobileB2CResponse, err error) {

	req, err := c.NewRequest("POST", "https://payments.sandbox.africastalking.com/mobile/b2c/request", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &at.MobileB2CResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Mobile Business To Business (B2B) APIs allow you to send payments to businesses
// e.g banks from your Payment Wallet.
func (c *Client) mobileB2B(ctx context.Context, p *at.MobileB2BPayload) (res *at.MobileB2BResponse, err error) {

	req, err := c.NewRequest("POST", "https://payments.sandbox.africastalking.com/mobile/b2b/request", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &at.MobileB2BResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
