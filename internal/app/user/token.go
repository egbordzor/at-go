package user

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/wondenge/at-go/internal/pkg"
	"github.com/wondenge/at-go/pkg/gen/africastalking"
	"io/ioutil"
	"net/http"
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
func (s *APIClient) Generate(ctx context.Context, p *africastalking.GeneratePayload) (res *africastalking.AccessTokenResponse, err error) {

	// Encode JSON from our instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := fmt.Errorf("could not marshall JSON: %w", err)
		fmt.Println(err.Error())
	}

	// NewRequestWithContext returns a Request suitable for use with Client.Do or Transport.RoundTrip.
	// Contains a context for controlling the entire lifetime of a request and its response.
	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.africastalking.com/tlsauth-token/generate", bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("could not create generate tlsauth token request: %v", err)
	}

	// Set Header Parameters
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", "MyAppApiKey")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not load default HTTP client: %w", err)
	}

	if err := s.logger.Log("info", fmt.Sprintf("africastalking.Generate")); err != nil {
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
		fmt.Errorf("could not read data from response body: %w", err)
		fmt.Println(err.Error())
	}

	res = &africastalking.AccessTokenResponse{}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		err := fmt.Errorf("could not unmarshal response body: %w", err)
		fmt.Println(err.Error())
	}

	// Print values of the object
	fmt.Printf("AccessTokenResponse: %v\n", res)

	return res, nil
}
