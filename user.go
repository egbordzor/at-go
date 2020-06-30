package atgo

import (
	"context"
	"fmt"
	"github.com/hashicorp/errwrap"
	"github.com/wondenge/at-go/user"
	"go.uber.org/zap"
)

// Generates a valid auth token
func (c *Client) GenerateToken(ctx context.Context, p *user.GeneratePayload) (res *user.AccessTokenResponse, err error) {

	if err := c.requestJSONBody(ctx, "POST", fmt.Sprintf("%s%s", c.AuthEndpoint, "/tlsauth-token/generate"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}

// Initiate an application data requestFormBody.
func (c *Client) InitiateAppData(ctx context.Context, p string) (res *user.UserResponse, err error) {

	if err := c.requestJSONBody(ctx, "GET", fmt.Sprintf("%s%s", c.UserEndpoint, "/version1/user"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}
