package atgo

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-kit/kit/log"
	"io"
	"io/ioutil"
	"net/http"
)

// Sandbox Endpoints
const (
	SMSTestURL     = "https://api.sandbox.africastalking.com"
	VoiceTestURL   = "https://voice.sandbox.africastalking.com"
	PaymentTestURL = "https://payments.sandbox.africastalking.com"
	AirtimeTestURL = "https://api.sandbox.africastalking.com"
	UserTestURL    = "https://api.sandbox.africastalking.com"
)

// Production Endpoints
const (
	SMSBaseURL     = "https://api.africastalking.com"
	VoiceBaseURL   = "https://voice.africastalking.com"
	PaymentBaseURL = "https://payments.africastalking.com"
	AirtimeBaseURL = "https://api.africastalking.com"
	IoTBaseURL     = "https://iot.africastalking.com"
	UserBaseURL    = "https://api.africastalking.com"
	AuthBaseURL    = "https://api.africastalking.com"
)

type (
	Client struct {
		Username        string
		SMSEndpoint     string
		VoiceEndpoint   string
		PaymentEndpoint string
		AirtimeEndpoint string
		IoTEndpoint     string
		UserEndpoint    string
		AuthEndpoint    string
		APIKey          string
		HTTPClient      *http.Client
		logger          log.Logger // log the requests.
	}
)

// NewClient returns new Production Client struct
// Use "test" for Sandbox Environment and "prod" for Production Environment
func NewClient(username string, apiKey string, Sandbox bool) (*Client, error) {

	if username == "" || apiKey == "" {
		return nil, errors.New("username, apiKey are required to create a Client")
	} else if username == "sandbox" {
		return &Client{
			Username:        "sandbox",
			SMSEndpoint:     SMSTestURL,
			VoiceEndpoint:   VoiceTestURL,
			PaymentEndpoint: PaymentTestURL,
			AirtimeEndpoint: AirtimeTestURL,
			IoTEndpoint:     "", // No Testing Environment for IoT
			UserEndpoint:    UserTestURL,
			AuthEndpoint:    "", // No Testing Environment for Auth
			APIKey:          apiKey,
			HTTPClient:      &http.Client{},
			logger:          nil,
		}, nil
	} else {
		return &Client{
			Username:        username,
			SMSEndpoint:     SMSBaseURL,
			VoiceEndpoint:   VoiceBaseURL,
			PaymentEndpoint: PaymentBaseURL,
			AirtimeEndpoint: AirtimeBaseURL,
			IoTEndpoint:     IoTBaseURL,
			UserEndpoint:    UserBaseURL,
			AuthEndpoint:    AuthBaseURL,
			APIKey:          apiKey,
			HTTPClient:      &http.Client{},
			logger:          nil,
		}, nil
	}
}

// NewRequest constructs a request
// Convert payload to a JSON
func (c *Client) newRequest(method, url string, payload interface{}) (*http.Request, error) {

	var buf io.Reader

	if payload != nil {

		// Encode JSON from our instance, using marshall.
		b, err := json.Marshal(&payload)
		if err != nil {
			err := fmt.Errorf("could not marshall JSON: %w", err)
			fmt.Println(err.Error())

			return nil, err
		}

		//buf = bytes.NewBuffer(b)
		buf = bytes.NewReader(b)
	}

	return http.NewRequest(method, url, buf)
}

func (c *Client) sendRequest(ctx context.Context, req *http.Request, v interface{}) (err error) {

	var resp *http.Response

	req = req.WithContext(ctx)

	resp, err = c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("could not load default HTTP client: %w", err)
	}

	if err := c.logger.Log("info", fmt.Sprintf("")); err != nil {
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
		err := fmt.Errorf("error: could not read data from response body: %v", err)
		fmt.Println(err.Error())
	}

	v = resp

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &v); err != nil {
		err := fmt.Errorf("could not unmarshal response body: %w", err)
		fmt.Println(err.Error())
	}

	// Print values of the object
	fmt.Printf("Response: %v\n", v)

	return nil
}
