package atgo

import (
	"context"
	"fmt"
	at "github.com/wondenge/at-go/internal/pkg/gen/africastalking"
)

type (
	User interface {

		// Initiate an application data request.
		initiateAppData(ctx context.Context, p string) (res *at.UserResponse, err error)
	}
)

// Initiate an application data request.
func (c *ATClient) initiateAppData(ctx context.Context, p string) (res *at.UserResponse, err error) {

	req, err := c.newRequest("GET", fmt.Sprintf("%s%s", c.UserEndpoint, "/version1/user"), p)
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	res = &at.UserResponse{}
	if err := c.sendRequest(ctx, req, res); err != nil {
		return nil, err
	}

	return res, nil
}
