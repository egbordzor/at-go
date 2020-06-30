package atgo

import (
	"context"
	"fmt"
	"github.com/hashicorp/errwrap"
	airtyme "github.com/wondenge/at-go/airtime"
	"go.uber.org/zap"
)

// Send Airtime.
// Payload attributes `p *airtyme.AirtimePayload` passed as key value args.
func (c *Client) SendAirtime(ctx context.Context, args map[string]string) (res *airtyme.AirtimeResponse, err error) {

	if err := c.requestFormBody(ctx, "POST", fmt.Sprintf("%s%s", c.AirtimeEndpoint, "/version1/airtime/send"), args, res); err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}
