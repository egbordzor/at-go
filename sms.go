package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

const (
	SMSLiveURL    = "https://api.africastalking.com/version1/messaging"
	SMSSandboxURL = "https://api.sandbox.africastalking.com/version1/messaging"
)

// Send Bulk SMS
func (c *Client) sendBulkSMS(ctx context.Context, p *africastalking.BulkPayload) (res *africastalking.BulkResponse, err error) {

	req, err := c.NewRequest("POST", "https://api.sandbox.africastalking.com/version1/messaging", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", "MyAppApiKey")

	res = &africastalking.BulkResponse{}

	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Incrementally fetch messages from application inbox.
func (c *Client) fetchSMS(ctx context.Context, p *africastalking.FetchMsgPayload) (res *africastalking.FetchMsgResponse, err error) {

	req, err := c.NewRequest("GET", "https://api.sandbox.africastalking.com/version1/messaging?username=MyAppUsername&lastReceivedId=0", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Apikey", "MyAppApiKey")

	res = &africastalking.FetchMsgResponse{}

	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Send Premium SMS
func (c *Client) sendPremiumSMS(ctx context.Context, p *africastalking.PremiumPayload) (res *africastalking.PremiumSMSResponse, err error) {

	req, err := c.NewRequest("POST", "https://api.sandbox.africastalking.com/version1/messaging", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", "MyAppApiKey")

	res = &africastalking.PremiumSMSResponse{}

	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Generate a checkout token
func (c *Client) newCheckoutToken(ctx context.Context, p *africastalking.CheckoutTokenPayload) (res *africastalking.CheckoutTokenResponse, err error) {

	req, err := c.NewRequest("POST", "https://api.sandbox.africastalking.com/checkout/token/create", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res = &africastalking.CheckoutTokenResponse{}

	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Subscribe a phone number
func (c *Client) newPremiumSubscription(ctx context.Context, p *africastalking.NewSubPayload) (res *africastalking.NewSubResponse, err error) {

	req, err := c.NewRequest("POST", "https://api.sandbox.africastalking.com/version1/subscription/create", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", "MyAppApiKey")

	res = &africastalking.NewSubResponse{}

	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Incrementally fetch your premium sms subscriptions.
func (c *Client) fetchPremiumSubscription(ctx context.Context, p *africastalking.FetchSubPayload) (res *africastalking.FetchSubResponse, err error) {

	req, err := c.NewRequest("GET", "https://api.sandbox.africastalking.com/version1/subscription?username=MyAppUsername&shortCode=myPremiumShortCode&keyword=myPremiumKeyword&lastReceivedId=0", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Apikey", "MyAppApiKey")

	res = &africastalking.FetchSubResponse{}

	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Delete a Premium SMS Subscription
func (c *Client) purgePremiumSubscription(ctx context.Context, p *africastalking.PurgeSubPayload) (res *africastalking.PurgeSubResponse, err error) {

	req, err := c.NewRequest("POST", "https://api.sandbox.africastalking.com/version1/subscription/delete", p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", "MyAppApiKey")

	res = &africastalking.PurgeSubResponse{}

	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
