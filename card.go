package atgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/errwrap"
	pay "github.com/wondenge/at-go/payments"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Collect money into your Payment Wallet by initiating transactions that deduct
// money from a customers Debit or Credit Card.
// These APIs are currently only available in Nigeria on MasterCard and Verve cards.
func (c *Client) CardCheckout(ctx context.Context, p *pay.CardCheckoutPayload) (res *pay.CardCheckoutResponse, err error) {

	// Encode JSON from our payload instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		err := errwrap.Wrapf("could not marshall JSON: {{err}}", err)
		log.Println(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/card/checkout/charge"), bytes.NewReader(b))
	if err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
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

// Allows your application to validate card checkout charge requests.
func (c *Client) CardCheckoutValidate(ctx context.Context, p *pay.CardCheckoutValidatePayload) (res *pay.CardCheckoutValidateResponse, err error) {

	// Encode JSON from our payload instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		return nil, errwrap.Wrapf("could not marshall JSON: {{err}}", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.PaymentEndpoint, "/card/checkout/validate"), bytes.NewReader(b))
	if err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
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
