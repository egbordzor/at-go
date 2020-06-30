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

// Mobile Checkout APIs allow you to initiate Customer to Business (C2B) payments
// on a mobile subscriber’s device.
// This allows for a smoother checkout experience, since the client will no longer
// need to remember the amount or an account number to complete the transaction.
func (c *Client) MobileCheckout(ctx context.Context, p *pay.MobileCheckoutPayload) (res *pay.MobileCheckoutResponse, err error) {

	// Encode JSON from our payload instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		return nil, errwrap.Wrapf("could not marshall JSON: {{err}}", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/mobile/checkout/request"), bytes.NewReader(b))
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

// Mobile Business To Consumer (B2C) APIs allow you to send payments to mobile
// subscribers from your Payment Wallet.
func (c *Client) MobileB2C(ctx context.Context, p *pay.MobileB2CPayload) (res *pay.MobileB2CResponse, err error) {

	// Encode JSON from our payload instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		return nil, errwrap.Wrapf("could not marshall JSON: {{err}}", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/mobile/b2c/request"), bytes.NewReader(b))
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

// Mobile Business To Business (B2B) APIs allow you to send payments to businesses
// e.g banks from your Payment Wallet.
func (c *Client) MobileB2B(ctx context.Context, p *pay.MobileB2BPayload) (res *pay.MobileB2BResponse, err error) {

	// Encode JSON from our payload instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		return nil, errwrap.Wrapf("could not marshall JSON: {{err}}", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/mobile/b2b/request"), bytes.NewReader(b))
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
