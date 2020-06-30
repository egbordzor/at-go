package atgo

import (
	"context"
	"encoding/json"
	"fmt"
	airtyme "github.com/wondenge/at-go/airtime"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Send Airtime.
// Payload attributes `p *airtyme.AirtimePayload` passed as key value args.
func (c *Client) SendAirtime(ctx context.Context, args map[string]string) (res *airtyme.AirtimeResponse, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.AirtimeEndpoint, "/version1/airtime/send"), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	req = req.WithContext(ctx)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		fmt.Errorf("could not load HTTP client: %w", err)
	}

	//  We're done reading from response body, lets close it.
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			fmt.Errorf("could not close response body: %w", err)
		}
	}()

	// Buffer the body
	// Read data from response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("bodyErr ", err.Error())
		fmt.Errorf("could not close response body: %w", err)

	}

	if resp.StatusCode != 200 {
		apiErr := &APIError{
			StatusCode: resp.StatusCode,
		}
		if err := json.Unmarshal(body, apiErr); err != nil {
			fmt.Fprintln(os.Stderr, "Invalid API response: "+string(body))
			return nil, fmt.Errorf("error unmarshaling %d error: %v", resp.StatusCode, err)
		}
		return nil, apiErr
	}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Errorf("could not unmarshal response body: %w", err)

	}

	return res, nil
}
