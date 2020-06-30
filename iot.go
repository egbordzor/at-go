package atgo

import (
	"context"
	"fmt"
	"github.com/hashicorp/errwrap"
	"github.com/wondenge/at-go/edge"
	"go.uber.org/zap"
)

// Publishes messages to remote devices.
func (c *Client) PublishIoT(ctx context.Context, p *edge.IoTPayload) (res *edge.IoTResponse, err error) {

	if err := c.requestJSONBody(ctx, "POST", fmt.Sprintf("%s%s", c.IoTEndpoint, "/data/publish"), p, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}
