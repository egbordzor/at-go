package atgo

import (
	"context"
	"fmt"
	at "github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

type (
	ATAuth interface {

		// Generates a valid auth token
		generateToken(ctx context.Context, p *at.GeneratePayload) (res *at.AccessTokenResponse, err error)
	}
)

// Generates a valid auth token
func (c *Client) generateToken(ctx context.Context, p *at.GeneratePayload) (res *at.AccessTokenResponse, err error) {

	req, err := c.newRequest("POST", fmt.Sprintf("%s%s", c.AuthEndpoint, "/tlsauth-token/generate"), p)
	if err != nil {
		return nil, fmt.Errorf("could not create generate tlsauth token request: %v", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.AccessTokenResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
