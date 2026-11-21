package test

import (
	"context"
	"encoding/base64"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestTDSWithBIN(t *testing.T) {
	for _, binEnabled := range []bool{true, false} {
		name := "without BIN"
		if binEnabled {
			name = "with BIN"
		}
		t.Run(name, func(t *testing.T) {
			var mu sync.Mutex
			var paths []string
			const html = "<html>3DS</html>"
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				mu.Lock()
				paths = append(paths, r.URL.Path)
				mu.Unlock()
				switch r.URL.Path {
				case "/payment/bin/check":
					_, _ = io.WriteString(w, `{"status":"success","cardAssociation":"VISA"}`)
				case "/payment/3dsecure/initialize":
					_, _ = io.WriteString(w, `{"status":"success","threeDSHtmlContent":"`+base64.StdEncoding.EncodeToString([]byte(html))+`"}`)
				default:
					t.Errorf("unexpected path: %s", r.URL.Path)
					w.WriteHeader(http.StatusNotFound)
				}
			}))
			defer server.Close()
			client, err := iyzipay.New("key", "secret", iyzipay.WithBaseURL(server.URL), iyzipay.WithHTTPClient(server.Client()))
			if err != nil {
				t.Fatal(err)
			}
			request := &iyzipay.InitTDSRequest{PaymentRequest: *paymentFixture(), CallbackUrl: "https://example.com/callback"}
			response, decoded, err := client.InitializeTDSPayment(context.Background(), request, iyzipay.WithBinRequest(binEnabled), iyzipay.WithHtmlDecodeRequest(true))
			if err != nil || response.Status != "success" || decoded != html {
				t.Fatalf("response=%+v decoded=%q err=%v", response, decoded, err)
			}
			want := []string{"/payment/3dsecure/initialize"}
			if binEnabled {
				want = append([]string{"/payment/bin/check"}, want...)
			}
			mu.Lock()
			defer mu.Unlock()
			if !reflect.DeepEqual(paths, want) {
				t.Fatalf("paths=%v want=%v", paths, want)
			}
		})
	}
}

func TestApp2AppCallback(t *testing.T) {
	const callback = "https://example.com/callback"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/in-store/payment" || r.Method != http.MethodPost {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		if r.Header.Get("X-Callback-Url") != callback {
			t.Error("missing callback header")
		}
		if r.Header.Get("X-Api-Key") != "updated-key" {
			t.Error("App2App credentials not forwarded")
		}
		_, _ = io.WriteString(w, `{"status":"success"}`)
	}))
	defer server.Close()
	client, err := iyzipay.NewApp2AppClient("key", "secret", "a2akey", "a2asecret", "merchant", iyzipay.WithBaseURL(server.URL), iyzipay.WithHTTPClient(server.Client()))
	if err != nil {
		t.Fatal(err)
	}
	client.A2AapiKey = "updated-key"
	_, err = client.App2CreatePayment(context.Background(), &iyzipay.A2ACreatePaymentRequest{Amount: "1.00", Email: "buyer@example.com", CallBackUrl: callback})
	if err != nil {
		t.Fatal(err)
	}
}
