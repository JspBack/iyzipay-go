package iyzipay

import (
	"errors"
	"net/http"
	"strings"
)

type Option func(*iyzipayClient)

// (TR) Iyzipay, Iyzipay API ile etkileşim kurmak için kullanılan istemciyi temsil eder.
//
// (EN) Iyzipay, creates a client to interact with IYzipay API.
type iyzipayClient struct {
	// (TR) APIKey, Iyzipay'den aldığınız API anahtarıdır.
	// 
	// (EN) APIKey, is the API key you get from Iyzipay.
	apiKey string

	// (TR) APISecret, Iyzipay'den aldığınız gizli anahtardır.
	// 
	// (EN) APISecret, is the secret key you get from Iyzipay.
	apiSecret string

	// (TR) BaseURI, Iyzipay API'si için temel URI'dir.
	// 
	// (EN) BaseURI, is the base URI for Iyzipay API.
	baseURI string

	// (TR) Istekler için kullanılacak client
	// 
	// (EN) The client used for requests
	client *http.Client

	// (TR) 3DS işlemlerinde kart konrolü yapılmasını isteyip istemediğinizi belirten değer.
	// 
	// (EN) Value indicating whether you want card control in 3DS transactions.
	binRequest bool

	// (TR) 3DS işlemlerinde html içeriğini decode etmek isteyip istemediğinizi belirten değer.
	// 
	// (EN) Value indicating whether you want to decode HTML contents in 3DS transactions.
	htmlDecodeRequest bool
}

// (TR) WithBinRequest, 3DS işlemlerinde kart konrolü yapılmasını isteyip istemediğinizi belirtir.
// 
// (EN) WithBinRequest, indicates whether you want card control in 3DS transactions.
func WithBinRequest(binRequest bool) Option {
	return func(c *iyzipayClient) {
		c.binRequest = binRequest
	}
}

// (TR) WithHtmlDecodeRequest, 3DS işlemlerinde html içeriğini decode etmek isteyip istemediğinizi belirtir.
// 
// (EN) WithHtmlDecodeRequest, indicates whether you want to decode html contents in 3DS transactions.
func WithHtmlDecodeRequest(htmlDecodeRequest bool) Option {
	return func(c *iyzipayClient) {
		c.htmlDecodeRequest = htmlDecodeRequest
	}
}

// (TR) New, yeni bir Iyzipay clientı oluşturur.
//
// (EN) New, creates a new Iyzipay client.
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
		binRequest:        true, // (TR) Varsayılan olarak true - (EN) True by default
		htmlDecodeRequest: true, // (TR) Varsayılan olarak true - (EN) True by default
	}

	for _, opt := range opts {
		opt(client)
	}

	return client, nil
}
