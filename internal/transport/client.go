package transport

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

var ErrNilRequest = errors.New("iyzipay: request is nil")

type Executor interface {
	Do(context.Context, string, string, any, any, ...http.Header) error
}
type Client struct {
	HTTPClient                 *http.Client
	BaseURL, APIKey, APISecret string
	Headers                    http.Header
}

func (c *Client) Do(ctx context.Context, method, path string, request, response any, extraHeaders ...http.Header) error {
	headers := c.Headers.Clone()
	for _, extra := range extraHeaders {
		if headers == nil {
			headers = make(http.Header)
		}
		for name, values := range extra {
			headers[name] = append([]string(nil), values...)
		}
	}

	var body []byte
	var err error
	if request != nil {
		body, err = json.Marshal(request)
		if err != nil {
			return fmt.Errorf("iyzipay: encode request: %w", err)
		}
	}
	data, err := Send(ctx, c.HTTPClient, method, c.BaseURL, c.APIKey, c.APISecret, path, body, headers)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(data, response); err != nil {
		return fmt.Errorf("iyzipay: decode response: %w", err)
	}
	return nil
}
