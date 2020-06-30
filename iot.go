package atgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/errwrap"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/wondenge/at-go/edge"
)

// Publishes messages to remote devices.
func (c *Client) PublishIoT(ctx context.Context, p *edge.IoTPayload) (res *edge.IoTResponse, err error) {

	// Encode JSON from our payload instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		return nil, errwrap.Wrapf("could not marshall JSON: {{err}}", err)
	}

	// Set Header Parameters
	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.IoTEndpoint, "/data/publish"), bytes.NewReader(b))
	if err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	req = req.WithContext(ctx)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		err := errwrap.Wrapf("could not load HTTP client: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	//  We're done reading from response body, lets close it.
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			err := errwrap.Wrapf("could not close response body: {{err}}", err)
			c.Log.Info("error", zap.Error(err))
		}
	}()

	// Buffer the body
	// Read data from response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := errwrap.Wrapf("could not close response body: {{err}}", err)
		c.Log.Info("error", zap.Error(err))

	}

	if resp.StatusCode != 200 {
		apiErr := &APIError{
			StatusCode: resp.StatusCode,
		}
		if err := json.Unmarshal(body, apiErr); err != nil {

			_, err := fmt.Fprintln(os.Stderr, "Invalid API response: "+string(body))
			return nil, fmt.Errorf("error unmarshaling %d error: %v", resp.StatusCode, err)
		}
		return nil, apiErr
	}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		err := errwrap.Wrapf("could not unmarshal response body: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}
