package atgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wondenge/at-go/pkg/gen/africastalking"
	"io/ioutil"
	"net/http"
)

// Publishes messages to remote devices.
func (s *africastalkingsrvc) PublishIoT(ctx context.Context, p *africastalking.IoTPayload) (res *africastalking.IoTResponse, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	// Set Header Parameters
	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s", "https://iot.africastalking.com/data/publish"), bytes.NewReader(b))
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

	if err := s.logger.Log("info", fmt.Sprintf("africastalking.PublishIoT")); err != nil {
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

	res = &africastalking.IoTResponse{}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		err := fmt.Errorf("could not unmarshal response body: %w", err)
		fmt.Println(err.Error())
	}

	return res, err
}
