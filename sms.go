package atgo

import (
	"context"
	"fmt"
	at "github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

type (
	SMS interface {

		// Send Bulk SMS
		sendBulkSMS(ctx context.Context, p *at.BulkPayload) (res *at.BulkResponse, err error)

		// Incrementally fetch messages from application inbox.
		fetchSMS(ctx context.Context, p *at.FetchMsgPayload) (res *at.FetchMsgResponse, err error)

		// Send Premium SMS
		sendPremiumSMS(ctx context.Context, p *at.PremiumPayload) (res *at.PremiumSMSResponse, err error)

		// Generate a checkout token
		newCheckoutToken(ctx context.Context, p *at.CheckoutTokenPayload) (res *at.CheckoutTokenResponse, err error)

		// Subscribe a phone number
		newPremiumSubscription(ctx context.Context, p *at.NewSubPayload) (res *at.NewSubResponse, err error)

		// Incrementally fetch your premium sms subscriptions.
		fetchPremiumSubscription(ctx context.Context, p *at.FetchSubPayload) (res *at.FetchSubResponse, err error)

		// Delete a Premium SMS Subscription
		purgePremiumSubscription(ctx context.Context, p *at.PurgeSubPayload) (res *at.PurgeSubResponse, err error)
	}
)

// Send Bulk SMS
func (c *ATClient) sendBulkSMS(ctx context.Context, p *at.BulkPayload) (res *at.BulkResponse, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/messaging"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.BulkResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Incrementally fetch messages from application inbox.
func (c *ATClient) fetchSMS(ctx context.Context, p *at.FetchMsgPayload) (res *at.FetchMsgResponse, err error) {

	req, err := c.newRequest("GET", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/messaging"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.FetchMsgResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Send Premium SMS
func (c *ATClient) sendPremiumSMS(ctx context.Context, p *at.PremiumPayload) (res *at.PremiumSMSResponse, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/messaging"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.PremiumSMSResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Generate a checkout token
func (c *ATClient) newCheckoutToken(ctx context.Context, p *at.CheckoutTokenPayload) (res *at.CheckoutTokenResponse, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.SMSEndpoint, "/checkout/token/create"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res = &at.CheckoutTokenResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Subscribe a phone number
func (c *ATClient) newPremiumSubscription(ctx context.Context, p *at.NewSubPayload) (res *at.NewSubResponse, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/subscription/create"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.NewSubResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Incrementally fetch your premium sms subscriptions.
func (c *ATClient) fetchPremiumSubscription(ctx context.Context, p *at.FetchSubPayload) (res *at.FetchSubResponse, err error) {

	req, err := c.newRequest("GET", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/subscription"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.FetchSubResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}

// Delete a Premium SMS Subscription
func (c *ATClient) purgePremiumSubscription(ctx context.Context, p *at.PurgeSubPayload) (res *at.PurgeSubResponse, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/subscription/delete"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.PurgeSubResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
