package iyzipay

import (
	"errors"
	"net/http"
	"strings"

	"github.com/JspBack/iyzipay-go/utils"
)

// Iyzıco's documentation is whole another level. Words can't describe the struggle 💀
// (https://docs.iyzico.com)

// Iyzipay, Iyzipay API ile etkileşim kurmak için kullanılan istemciyi temsil eder.
type IyzipayClient struct {
	// APIKey, Iyzipay'den aldığınız anahtardır.
	apiKey string

	// APISecret, Iyzipay'den aldığınız gizli anahtardır.
	apiSecret string

	// BaseURI, Iyzipay API'si için temel URI'dir.
	baseURI string

	// Istekler için kullanılacak client
	client *http.Client
}

// App2AppClient, Iyzipay App2App API ile etkileşim kurmak için kullanılan istemciyi temsil eder.
type App2AppClient struct {
	// APIKey, Iyzipay'den aldığınız anahtardır.
	apiKey string

	// APISecret, Iyzipay'den aldığınız gizli anahtardır.
	apiSecret string

	// App2App APIKey, Iyzipay'den aldığınız App2App işlemler için özel anahtardır.
	A2AapiKey string

	// App2App APISecret, Iyzipay'den aldığınız App2App işlemler için özel gizli anahtardır.
	A2AsecretKey string

	// BaseURI, Iyzipay API'si için temel URI'dir.
	baseURI string

	// App2App işlemler için kullanılacak satıcı kimliği
	merchantId string

	// App2App istekleri için kullanılacak client
	client *http.Client
}

// New, yeni bir Iyzipay clientı oluşturur.
func New(apiKey, apiSecret string) (*IyzipayClient, error) {
	if apiKey == "" {
		return nil, errors.New("API key is required")
	}

	if apiSecret == "" {
		return nil, errors.New("API secret is required")
	}

	baseURI := utils.APIURL
	if strings.HasPrefix(apiKey, "sandbox-") {
		baseURI = utils.SandboxAPIURL
	}

	return &IyzipayClient{
		apiKey:    apiKey,
		apiSecret: apiSecret,
		baseURI:   baseURI,
		client:    &http.Client{},
	}, nil
}

// NewApp2AppClient, yeni bir App2App clientı oluşturur.
func NewApp2AppClient(apiKey, apiSecret, a2AapiKey, a2AsecretKey, merchantId string) (*App2AppClient, error) {
	if apiKey == "" {
		return nil, errors.New("API key is required")
	}

	if apiSecret == "" {
		return nil, errors.New("API secret is required")
	}

	baseURI := utils.APIURL
	if strings.HasPrefix(apiKey, "sandbox-") {
		baseURI = utils.SandboxAPIURL
	}

	return &App2AppClient{
		apiKey:       apiKey,
		apiSecret:    apiSecret,
		A2AapiKey:    a2AapiKey,
		A2AsecretKey: a2AsecretKey,
		merchantId:   merchantId,
		baseURI:      baseURI,
		client:       &http.Client{},
	}, nil
}
