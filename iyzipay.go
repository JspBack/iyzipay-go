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
