package atgo

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/errwrap"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func (c *Client) doHTTP(req *http.Request, res interface{}) error {

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
			return fmt.Errorf("error unmarshaling %d error: %v", resp.StatusCode, err)
		}
		return apiErr
	}

	// Parse the JSON-encoded data from response body.
	// The data is stored in the value pointed by response.
	if err := json.Unmarshal(body, &res); err != nil {
		err := errwrap.Wrapf("could not unmarshal response body: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return nil
}

func (c *Client) requestFormBody(ctx context.Context, method, path string, args map[string]string, response interface{}) error {

	var req *http.Request

	var err error

	if method == "POST" {
		req, err = c.postFormBody(ctx, method, c.BaseURL+path, args)
	} else {
		req, err = c.requestQueryString(ctx, c.BaseURL+path, args)
	}

	if err != nil {
		err := errwrap.Wrapf("error: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return c.doHTTP(req, response)
}

func (c *Client) checkoutTokenBody(ctx context.Context, method, path string, args map[string]string, response interface{}) error {

	var req *http.Request

	var err error

	if method == "POST" {
		req, err = c.postTokenBody(ctx, method, c.BaseURL+path, args)
	} else {
		req, err = c.requestQueryString(ctx, c.BaseURL+path, args)
	}

	if err != nil {
		err := errwrap.Wrapf("error: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return c.doHTTP(req, response)
}

func (c *Client) requestJSONBody(ctx context.Context, method, path string, payload, response interface{}) error {

	var req *http.Request

	var err error

	if method == "POST" {
		req, err = c.postJSONBody(ctx, method, c.BaseURL+path, payload)
	} else {
		req, err = c.getJSONBody(ctx, method, c.BaseURL+path, payload)
	}
	if err != nil {
		err := errwrap.Wrapf("error: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	return c.doHTTP(req, response)
}

func (c *Client) postFormBody(ctx context.Context, method, uri string, args map[string]string) (*http.Request, error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequestWithContext(ctx, method, uri, strings.NewReader(form.Encode()))
	if err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Apikey", c.APIKey)

	return req, nil
}

func (c *Client) postTokenBody(ctx context.Context, method, uri string, args map[string]string) (*http.Request, error) {

	form := url.Values{}
	for k, v := range args {
		form.Set(k, v)
	}

	req, err := http.NewRequestWithContext(ctx, method, uri, strings.NewReader(form.Encode()))
	if err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return req, nil
}

func (c *Client) postJSONBody(ctx context.Context, method, uri string, payload interface{}) (*http.Request, error) {

	// Encode JSON from our payload instance, using marshall.
	b, err := json.Marshal(payload)
	if err != nil {
		err := errwrap.Wrapf("could not marshall JSON: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	req, err := http.NewRequestWithContext(ctx, method, uri, bytes.NewReader(b))
	if err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", c.APIKey)

	return req, nil
}

func (c *Client) requestQueryString(ctx context.Context, uri string, args map[string]string) (*http.Request, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", uri, nil)
	if err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}
	if len(args) > 0 {
		query := req.URL.Query()
		for k, v := range args {
			query.Add(k, v)
		}
		req.URL.RawQuery = query.Encode()
	}
	return req, nil
}

func (c *Client) getJSONBody(ctx context.Context, method, uri string, payload interface{}) (*http.Request, error) {

	// Encode JSON from our payload instance, using marshall.
	b, err := json.Marshal(payload)
	if err != nil {
		err := errwrap.Wrapf("could not marshall JSON: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	req, err := http.NewRequestWithContext(ctx, method, uri, bytes.NewReader(b))
	if err != nil {
		err := errwrap.Wrapf("could not make new http request: {{err}}", err)
		c.Log.Info("error", zap.Error(err))
	}

	// Set Header Parameters.
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Apikey", c.APIKey)

	return req, nil
}
