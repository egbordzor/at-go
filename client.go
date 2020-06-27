package atgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/log"
	"io"
	"io/ioutil"
	"net/http"
)



type Client struct {
	username         string
	smsEndpoint      string
	voiceEndpoint    string
	paymentsEndpoint string
	airtimeEndpoint  string
	iotEndpoint      string
	userEndpoint     string
	apiKey           string
	HTTPClient       *http.Client
	logger           log.Logger // log the requests.
}

// NewClient returns new Production Client struct
func NewClient(Username, APIKey string) (*Client, error) {
	return &Client{
		username:         Username,
		smsEndpoint:      "",
		voiceEndpoint:    "",
		paymentsEndpoint: "",
		airtimeEndpoint:  "",
		iotEndpoint:      "",
		userEndpoint:     "",
		apiKey:           "",
		HTTPClient:       &http.Client{},
		logger:           nil,
	}, nil
}

// NewSandboxClient returns new Sandbox Client struct
func NewSandboxClient(APIKey string) (*Client, error) {
	return &Client{
		username:         "sandbox",
		smsEndpoint:      "",
		voiceEndpoint:    "",
		paymentsEndpoint: "",
		airtimeEndpoint:  "",
		iotEndpoint:      "",
		userEndpoint:     "",
		apiKey:           APIKey,
		HTTPClient:       &http.Client{},
		logger:           nil,
	}, nil
}

// NewRequest constructs a request
// Convert payload to a JSON
func (c *Client) NewRequest(method, url string, payload interface{}) (*http.Request, error) {

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
