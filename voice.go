package atgo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/errwrap"
	"github.com/wondenge/at-go/voice"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Makes outbound calls.
// Payload attributes `p *MakeCallPayload` passed as key value args.
func (c *Client) MakeCall(ctx context.Context, args map[string]string) (res *voice.MakeCallResponse, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/calls"), strings.NewReader(form.Encode()))
	if err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

// Transfers calls to another number.
// Payload attributes `p *CallTransferPayload` passed as key value args.
func (c *Client) TransferCall(ctx context.Context, args map[string]string) (res *voice.CallTransferResponse, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/callTransfer"), strings.NewReader(form.Encode()))
	if err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

// Used when you have more calls than you can handle africastalking one time.
// Payload attributes `p *QueuedCallsPayload` passed as key value args.
func (c *Client) QueuedCall(ctx context.Context, args map[string]string) (res *voice.QueuedStatusResult, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/queueStatus"), strings.NewReader(form.Encode()))
	if err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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

// Uploads media or audio files to Africa'sTalking servers with the extension
// .mp3 or .wav
// Payload attributes `p *UploadMediaFile` passed as key value args.
func (c *Client) UploadMedia(ctx context.Context, args map[string]string) (res string, err error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", c.VoiceEndpoint, "/mediaUpload"), strings.NewReader(form.Encode()))
	if err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
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
			return "", fmt.Errorf("error unmarshaling %d error: %v", resp.StatusCode, err)
		}
		return "", apiErr
	}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		err := errwrap.Wrapf("could not unmarshal response body: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return res, nil
}
