package airtime

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

// Send Airtime.
func (c *atgo.Client) sendAirtime(ctx context.Context, p *africastalking.AirtimePayload) (res *africastalking.AirtimeResponse, err error) {

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", "", "/version1/airtime/send"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.apiKey)

	res = &africastalking.AirtimeResponse{}

	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
