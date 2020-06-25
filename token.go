package atgo

import (
	"context"
	"fmt"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

type (
	Africastalking struct {
		Username string
		APIKey   string
	}

	AccessTokenResponse struct {
		Token             string
		LifeTimeInSeconds int64
	}
)

// Generates a valid auth token
func (c *Client) generateToken(ctx context.Context, p *africastalking.GeneratePayload) (res *africastalking.AccessTokenResponse, err error) {

	// NewRequestWithContext returns a Request suitable for use with Client.Do or Transport.RoundTrip.
	// Contains a context for controlling the entire lifetime of a request and its response.
	req, err := c.NewRequest("POST", "https://api.africastalking.com/tlsauth-token/generate", p)
	if err != nil {
		return nil, fmt.Errorf("could not create generate tlsauth token request: %v", err)
	}

	// Set Header Parameters
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppApiKey")

	res = &africastalking.AccessTokenResponse{}

	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
