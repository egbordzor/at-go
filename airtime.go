package atgo

import (
	"context"
	"fmt"
	at "github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

const (
	AirtimeLiveURL    = "https://api.africastalking.com/version1/airtime/send"
	AirtimeSandboxURL = "https://api.sandbox.africastalking.com/version1/airtime/send"
)

type (
	Airtime interface {
		// Send Airtime.
		sendAirtime(ctx context.Context, p *at.AirtimePayload) (res *at.AirtimeResponse, err error)
	}
)

// Send Airtime.
func (c *Client) sendAirtime(ctx context.Context, p *at.AirtimePayload) (res *at.AirtimeResponse, err error) {

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", "", "/version1/airtime/send"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.apiKey)

	res = &at.AirtimeResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
