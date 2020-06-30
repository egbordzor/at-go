package atgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/wondenge/at-go/sms"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Send Bulk SMS
// Payload attributes `p *BulkPayload` passed as key value args.
func (c *Client) SendBulkSMS(ctx context.Context, args map[string]string) (res *sms.BulkResponse, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/messaging"), strings.NewReader(form.Encode()))
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

// Incrementally fetch messages from application inbox.
func (c *Client) FetchSMS(ctx context.Context, p *sms.FetchMsgPayload) (res *sms.FetchMsgResponse, err error) {

	// Encode JSON from our ReversalRequest instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Errorf("could not marshall JSON: %w", err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/messaging"), bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
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

// Send Premium SMS
// Payload attributes `p *PremiumPayload` passed as key value args.
func (c *Client) SendPremiumSMS(ctx context.Context, args map[string]string) (res *sms.PremiumSMSResponse, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/messaging"), strings.NewReader(form.Encode()))
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

// Generate a checkout token
// Payload attributes `p *CheckoutTokenPayload` passed as key value args.
func (c *Client) NewCheckoutToken(ctx context.Context, args map[string]string) (res *sms.CheckoutTokenResponse, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.SMSEndpoint, "/checkout/token/create"), strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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

// Subscribe a phone number
// Payload attributes `p *NewSubPayload` passed as key value args.
func (c *Client) NewPremiumSubscription(ctx context.Context, args map[string]string) (res *sms.NewSubResponse, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/subscription/create"), strings.NewReader(form.Encode()))
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

// Incrementally fetch your premium sms subscriptions.
func (c *Client) FetchPremiumSubscription(ctx context.Context, p *sms.FetchSubPayload) (res *sms.FetchSubResponse, err error) {

	// Encode JSON from our ReversalRequest instance, using marshall.
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Errorf("could not marshall JSON: %w", err)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/subscription"), bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("could not make new http request: %w", err)
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
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

// Delete a Premium SMS Subscription
// Payload attributes `p *PurgeSubPayload` passed as key value args.
func (c *Client) PurgePremiumSubscription(ctx context.Context, args map[string]string) (res *sms.PurgeSubResponse, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.SMSEndpoint, "/version1/subscription/delete"), strings.NewReader(form.Encode()))
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
