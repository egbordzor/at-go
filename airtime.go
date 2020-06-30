package airtyme

import (
	"context"
	"fmt"
)

type (
	Airtime interface {

		// Send Airtime.
		SendAirtime(ctx context.Context, p *AirtimePayload) (res *AirtimeResponse, err error)
	}
)

// Send Airtime.
func (c *at.ATClient) SendAirtime(ctx context.Context, p *AirtimePayload) (res *AirtimeResponse, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.AirtimeEndpoint, "/version1/airtyme/send"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
