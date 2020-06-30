package atgo

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Set a text to be read out to the caller.
// Payload attributes `p *SayPayload` passed as key value args.
func (c *Client) Say(ctx context.Context, args map[string]string) (res string, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
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
			return "", fmt.Errorf("error unmarshaling %d error: %v", resp.StatusCode, err)
		}
		return "", apiErr
	}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Errorf("could not unmarshal response body: %w", err)

	}

	return res, nil
}

// Play back an audio file located anywhere on the web.
// Payload attributes `p *PlayPayload` passed as key value args.
func (c *Client) Play(ctx context.Context, args map[string]string) (res string, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
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
			return "", fmt.Errorf("error unmarshaling %d error: %v", resp.StatusCode, err)
		}
		return "", apiErr
	}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Errorf("could not unmarshal response body: %w", err)

	}

	return res, nil
}

// Get digits a user enters on their phone in response to a prompt from
// application
// Payload attributes `p *GetDigitsPayload` passed as key value args.
func (c *Client) GetDigits(ctx context.Context, args map[string]string) (res string, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
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
			return "", fmt.Errorf("error unmarshaling %d error: %v", resp.StatusCode, err)
		}
		return "", apiErr
	}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Errorf("could not unmarshal response body: %w", err)

	}

	return res, nil
}

// Connect the user who called your phone number to an external phone number.
// Payload attributes `p *DialPayload` passed as key value args.
func (c *Client) Dial(ctx context.Context, args map[string]string) (res string, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
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
			return "", fmt.Errorf("error unmarshaling %d error: %v", resp.StatusCode, err)
		}
		return "", apiErr
	}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Errorf("could not unmarshal response body: %w", err)

	}

	return res, nil
}

// Record a calls session into an mp3 file.
// Payload attributes `p *RecordPayload` passed as key value args.
func (c *Client) Record(ctx context.Context, args map[string]string) (res string, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
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
			return "", fmt.Errorf("error unmarshaling %d error: %v", resp.StatusCode, err)
		}
		return "", apiErr
	}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Errorf("could not unmarshal response body: %w", err)

	}

	return res, nil
}

// Pass an incoming calls to a queue to be handled later.
// Payload attributes `p *EnqueuePayload` passed as key value args.
func (c *Client) Enqueue(ctx context.Context, args map[string]string) (res string, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
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
			return "", fmt.Errorf("error unmarshaling %d error: %v", resp.StatusCode, err)
		}
		return "", apiErr
	}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Errorf("could not unmarshal response body: %w", err)

	}

	return res, nil
}

// Pass the calls enqueued to a separate number to be handled.
// Payload attributes `p *DequeuePayload` passed as key value args.
func (c *Client) Dequeue(ctx context.Context, args map[string]string) (res string, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
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
			return "", fmt.Errorf("error unmarshaling %d error: %v", resp.StatusCode, err)
		}
		return "", apiErr
	}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Errorf("could not unmarshal response body: %w", err)

	}

	return res, nil
}

// Transfer control of the calls to the script whose URL is passed in.
// Payload attributes `p *RedirectPayload` passed as key value args.
func (c *Client) Redirect(ctx context.Context, args map[string]string) (res string, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
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
			return "", fmt.Errorf("error unmarshaling %d error: %v", resp.StatusCode, err)
		}
		return "", apiErr
	}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Errorf("could not unmarshal response body: %w", err)

	}

	return res, nil
}

// Reject an incoming calls without incurring any usage charges.
// Payload attributes `p *RejectPayload` passed as key value args.
func (c *Client) Reject(ctx context.Context, args map[string]string) (res string, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), strings.NewReader(form.Encode()))
	if err != nil {
		return "", fmt.Errorf("could not make new http request: %w", err)
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
			return "", fmt.Errorf("error unmarshaling %d error: %v", resp.StatusCode, err)
		}
		return "", apiErr
	}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		fmt.Errorf("could not unmarshal response body: %w", err)

	}

	return res, nil
}
