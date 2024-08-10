package iyzipay

import (
	"errors"
	"net/http"
	"strings"
)

type Option func(*iyzipayClient)

// Iyzipay, Iyzipay API ile etkileşim kurmak için kullanılan istemciyi temsil eder.
type iyzipayClient struct {
	// APIKey, Iyzipay'den aldığınız anahtardır.
	apiKey string

	// APISecret, Iyzipay'den aldığınız gizli anahtardır.
	apiSecret string

	// BaseURI, Iyzipay API'si için temel URI'dir.
	baseURI string

	// Istekler için kullanılacak client
	client *http.Client

	// 3DS işlemlerinde kart konrolü yapılmasını isteyip istemediğinizi belirten değer.
	binRequest bool

	// 3DS işlemlerinde html içeriğini decode etmek isteyip istemediğinizi belirten değer.
	htmlDecodeRequest bool
}

// WithBinRequest, 3DS işlemlerinde kart konrolü yapılmasını isteyip istemediğinizi belirtir.
func WithBinRequest(binRequest bool) Option {
	return func(c *iyzipayClient) {
		c.binRequest = binRequest
	}
}

// WithHtmlDecodeRequest, 3DS işlemlerinde html içeriğini decode etmek isteyip istemediğinizi belirtir.
func WithHtmlDecodeRequest(htmlDecodeRequest bool) Option {
	return func(c *iyzipayClient) {
		c.htmlDecodeRequest = htmlDecodeRequest
	}
}

// New, yeni bir Iyzipay clientı oluşturur.
func New(apiKey, apiSecret string, opts ...Option) (*iyzipayClient, error) {
	if apiKey == "" {
		return nil, errors.New("API key is required")
	}

	if apiSecret == "" {
		return nil, errors.New("API secret is required")
	}

	baseURI := "https://api.iyzipay.com"

	if strings.HasPrefix(apiKey, "sandbox-") {
		baseURI = "https://sandbox-api.iyzipay.com"
	}

	client := &iyzipayClient{
		apiKey:            apiKey,
		apiSecret:         apiSecret,
		baseURI:           baseURI,
		client:            &http.Client{},
		binRequest:        true, // Default olarak true
		htmlDecodeRequest: true, // Default olarak true
	}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}
