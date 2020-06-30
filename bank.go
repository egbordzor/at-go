package atgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/errwrap"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	pay "github.com/wondenge/at-go/payments"
)

// Collect money into your payment wallet.
func (c *Client) BankCheckout(ctx context.Context, p *pay.BankCheckoutPayload) (res *pay.BankCheckoutResponse, err error) {

	// Encode JSON from our payload instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		return nil, errwrap.Wrapf("could not marshall JSON: {{err}}", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/bank/checkout/charge"), bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
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

// Validate a bank checkout charge request
func (c *Client) BankCheckoutValidate(ctx context.Context, p *pay.BankCheckoutValidatePayload) (res *pay.BankCheckoutValidateResponse, err error) {

	// Encode JSON from our payload instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		return nil, errwrap.Wrapf("could not marshall JSON: {{err}}", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/bank/checkout/validate"), bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
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

// Initiate a bank transfer request.
func (c *Client) BankTransfer(ctx context.Context, p *pay.BankTransferPayload) (res *pay.BankTransferResponse, err error) {

	// Encode JSON from our payload instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		return nil, errwrap.Wrapf("could not marshall JSON: {{err}}", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/bank/transfer"), bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
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
