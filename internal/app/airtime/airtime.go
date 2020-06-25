package airtime

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/wondenge/at-go/internal/pkg/gen/africastalking"
	"io/ioutil"
	"net/http"
	"time"
)

const (

	// APIBaseSandBox points to the sandbox (for testing) version of the API
	APIBaseSandBox = ""

	// APIBaseLive points to the live version of the API
	APIBaseLive = ""
)

type (
	Client struct {
		Client         *http.Client
		APIKey         string
		Username       string
		Sandbox        string
		Live           string
		Token          *africastalking.AccessTokenResponse
		tokenExpiresAt time.Time
		logger         log.Logger // log the requests.
	}
)

// Send Airtime.
func (c *Client) SendAirtime(ctx context.Context, p *africastalking.AirtimePayload) (res *africastalking.AirtimeResponse, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s%s", "", "/version1/airtime/send"), bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", "MyAppAPIKey")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not load default HTTP client: %w", err)
	}

	if err := c.logger.Log("info", fmt.Sprintf("africastalking.SendAirtime")); err != nil {
		err := fmt.Errorf("could not log to stdout: %w", err)
		fmt.Println(err.Error())
	}

	// We're done reading from response body, lets close it.
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			err := fmt.Errorf("could not close response body: %w", err)
			fmt.Println(err.Error())
		}
	}()

	// Read data from response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("oauth2: cannot fetch token: %v", err)
		fmt.Println(err.Error())
	}

	res = &africastalking.AirtimeResponse{}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		err := fmt.Errorf("could not unmarshal response body: %w", err)
		fmt.Println(err.Error())
	}

	// Print values of the object
	fmt.Printf("AirtimeResponse: %v\n", res)

	return res, nil
}
