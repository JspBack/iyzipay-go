package iyzipay

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/JspBack/iyzipay-go/internal/transport"
)

const (
	SandboxURL    = transport.SandboxURL
	ProductionURL = transport.ProductionURL
)

// Option, istemci oluşturulurken uygulanacak yapılandırmayı tanımlar.
type Option func(*Client) error

// WithHTTPClient, bağlantı ve zaman aşımı ayarlarıyla birlikte kullanılacak HTTP istemcisini belirler.
func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) error {
		if client == nil {
			return fmt.Errorf("iyzipay: HTTP client is nil")
		}
		c.config.request.HTTPClient = client
		return nil
	}
}

// WithBaseURL, yerel testler dahil otomatik ortam seçimi yerine kullanılacak temel URL adresini belirler.
func WithBaseURL(raw string) Option {
	return func(c *Client) error {
		u, err := url.Parse(raw)
		if err != nil || u.Host == "" || (u.Scheme != "https" && u.Scheme != "http") || u.User != nil || u.RawQuery != "" || u.Fragment != "" {
			return fmt.Errorf("iyzipay: invalid base URL")
		}
		c.config.request.BaseURL = strings.TrimRight(raw, "/")
		return nil
	}
}
