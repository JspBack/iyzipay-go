package iyzipay

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/JspBack/iyzipay-go/internal/transport"
	"github.com/JspBack/iyzipay-go/internal/validation"
)

// Iyzıco's documentation is whole another level. Words can't describe the struggle 💀
// (https://docs.iyzico.com)

type clientConfig struct {
	request   transport.Client
	validator *validation.Validator
}

// Client, Iyzipay API ile etkileşim kurmak için kullanılan istemciyi temsil eder.
type Client struct{ config clientConfig }

// App2AppClient, Iyzipay App2App API istemcisidir.
type App2AppClient struct {
	config       clientConfig
	A2AapiKey    string
	A2AsecretKey string
	merchantId   string
}

// New, yeni bir Iyzipay clientı oluşturur.
func New(apiKey, apiSecret string, opts ...Option) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New("API key is required")
	}

	if apiSecret == "" {
		return nil, errors.New("API secret is required")
	}

	baseURI := transport.ProductionURL
	if strings.HasPrefix(apiKey, "sandbox-") {
		baseURI = transport.SandboxURL
	}

	c := &Client{config: clientConfig{validator: validation.New(), request: transport.Client{
		APIKey: apiKey, APISecret: apiSecret, BaseURL: baseURI,
		HTTPClient: &http.Client{Timeout: 30 * time.Second},
	}}}

	for _, opt := range opts {
		if opt == nil {
			return nil, errors.New("iyzipay: nil option")
		}
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// NewApp2AppClient, yeni bir App2App clientı oluşturur.
func NewApp2AppClient(apiKey, apiSecret, a2AapiKey, a2AsecretKey, merchantId string, opts ...Option) (*App2AppClient, error) {
	c, err := New(apiKey, apiSecret, opts...)
	if err != nil {
		return nil, err
	}
	if a2AapiKey == "" || a2AsecretKey == "" || merchantId == "" {
		return nil, errors.New("iyzipay: App2App credentials and merchant ID are required")
	}
	return &App2AppClient{config: c.config, A2AapiKey: a2AapiKey, A2AsecretKey: a2AsecretKey, merchantId: merchantId}, nil
}

func (c *Client) transport() *transport.Client { return &c.config.request }
func (c *App2AppClient) transport() *transport.Client {
	request := c.config.request
	request.Headers = http.Header{"X-Api-Key": {c.A2AapiKey}, "X-Secret-Key": {c.A2AsecretKey}, "X-Merchant-Id": {c.merchantId}}
	return &request
}
