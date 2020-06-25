package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

// Publishes messages to remote devices.
func (c *Client) publishIoT(ctx context.Context, p *africastalking.IoTPayload) (res *africastalking.IoTResponse, err error) {

	// Set Header Parameters
	req, err := c.NewRequest("POST", fmt.Sprintf("%s", "https://iot.africastalking.com/data/publish"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", "MyAppAPIKey")

	res = &africastalking.IoTResponse{}

	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
