package test

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/JspBack/iyzipay-go"
)

func TestClientOptions(t *testing.T) {
	_, err := iyzipay.New("sandbox-key", "secret")
	if err != nil {
		t.Fatal(err)
	}
	for _, opt := range []iyzipay.Option{nil, iyzipay.WithHTTPClient(nil), iyzipay.WithBaseURL("://bad"), iyzipay.WithBaseURL("https://example.com?key=x")} {
		if _, err := iyzipay.New("key", "secret", opt); err == nil {
			t.Fatal("expected invalid option error")
		}
	}
	if _, err := iyzipay.New("", "secret"); err == nil {
		t.Fatal("empty key accepted")
	}
	if _, err := iyzipay.NewApp2AppClient("key", "secret", "", "", ""); err == nil {
		t.Fatal("empty App2App credentials accepted")
	}
}

func TestBINEndpoint(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/payment/bin/check" {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL)
		}
		var body iyzipay.BinRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Error(err)
		}
		if body.BinNumber != "454671" {
			t.Error("incorrect request body")
		}
		value, err := base64.StdEncoding.DecodeString(strings.TrimPrefix(r.Header.Get("Authorization"), "IYZWSv2 "))
		if err != nil {
			t.Error(err)
		}
		if !strings.Contains(string(value), "randomKey:"+r.Header.Get("X-Iyzi-Rnd")+"&signature:") {
			t.Error("nonce mismatch")
		}
		_, _ = io.WriteString(w, `{"status":"success","cardAssociation":"VISA"}`)
	}))
	defer server.Close()
	c, err := iyzipay.New("key", "secret", iyzipay.WithBaseURL(server.URL), iyzipay.WithHTTPClient(server.Client()))
	if err != nil {
		t.Fatal(err)
	}
	response, err := c.BinControlRequest(context.Background(), &iyzipay.BinRequest{BinNumber: "454671"})
	if err != nil || response.CardAssociation != "VISA" {
		t.Fatalf("response=%+v err=%v", response, err)
	}
	if _, err := c.BinControlRequest(context.Background(), nil); !errors.Is(err, iyzipay.ErrNilRequest) {
		t.Fatalf("nil request: %v", err)
	}
}

func TestResponseErrors(t *testing.T) {
	cases := []struct {
		name   string
		status int
		body   string
		api    bool
		code   string
	}{
		{"provider", 200, `{"status":"failure","errorCode":"100","errorMessage":"declined","conversationId":"order-1"}`, true, "100"},
		{"unauthorized", 401, `{"message":"unauthorized"}`, true, ""},
		{"server", 500, "<html>unavailable</html>", true, ""},
		{"malformed", 200, "{", false, ""},
		{"null", 200, "null", false, ""},
		{"missing status", 200, "{}", false, ""},
		{"wrong field type", 200, `{"status":"success","cardAssociation":123}`, false, ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tc.status)
				_, _ = io.WriteString(w, tc.body)
			}))
			defer server.Close()
			c, _ := iyzipay.New("key", "secret", iyzipay.WithBaseURL(server.URL), iyzipay.WithHTTPClient(server.Client()))
			_, err := c.BinControlRequest(context.Background(), &iyzipay.BinRequest{BinNumber: "454671"})
			if err == nil {
				t.Fatal("expected error")
			}
			var apiErr *iyzipay.APIError
			if errors.As(err, &apiErr) != tc.api {
				t.Fatalf("unexpected error type: %T: %v", err, err)
			}
			if tc.api && (apiErr.HTTPStatus != tc.status || apiErr.Code != tc.code) {
				t.Fatalf("incorrect API error: %+v", apiErr)
			}
		})
	}
}

func TestCancellation(t *testing.T) {
	release := make(chan struct{})
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
		case <-release:
		}
	}))
	defer server.Close()
	defer close(release)
	c, _ := iyzipay.New("key", "secret", iyzipay.WithBaseURL(server.URL), iyzipay.WithHTTPClient(server.Client()))
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel()
	_, err := c.BinControlRequest(ctx, &iyzipay.BinRequest{BinNumber: "454671"})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected deadline: %v", err)
	}
}

func TestApp2App(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Api-Key") != "a2akey" || r.Header.Get("X-Secret-Key") != "a2asecret" || r.Header.Get("X-Merchant-Id") != "merchant" {
			t.Error("missing App2App headers")
		}
		if r.Method != http.MethodGet {
			t.Error("expected GET")
		}
		_, _ = io.WriteString(w, `{"status":"success"}`)
	}))
	defer server.Close()
	c, err := iyzipay.NewApp2AppClient("key", "secret", "a2akey", "a2asecret", "merchant", iyzipay.WithBaseURL(server.URL), iyzipay.WithHTTPClient(server.Client()))
	if err != nil {
		t.Fatal(err)
	}
	if _, err = c.App2AppListUsers(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func TestPaymentValidation(t *testing.T) {
	c, _ := iyzipay.New("key", "secret")
	if _, err := c.CreatePayment(context.Background(), &iyzipay.PaymentRequest{}); err == nil {
		t.Fatal("empty payment accepted")
	}
	if _, err := c.CreatePayment(context.Background(), nil); !errors.Is(err, iyzipay.ErrNilRequest) {
		t.Fatal(err)
	}
}

func TestCreatePayment(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/payment/auth" {
			t.Errorf("unexpected payment path: %s", r.URL.Path)
		}
		_, _ = io.WriteString(w, `{"status":"success","paymentId":"123"}`)
	}))
	defer server.Close()
	c, _ := iyzipay.New("key", "secret", iyzipay.WithBaseURL(server.URL), iyzipay.WithHTTPClient(server.Client()))
	request := paymentFixture()
	response, err := c.CreatePayment(context.Background(), request)
	if err != nil || response.PaymentID != "123" {
		t.Fatalf("payment=%+v err=%v", response, err)
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestClientDefaults(t *testing.T) {
	original := http.DefaultTransport
	t.Cleanup(func() { http.DefaultTransport = original })
	http.DefaultTransport = roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.URL.String() != iyzipay.SandboxURL+"/payment/bin/check" {
			t.Errorf("unexpected sandbox URL: %s", req.URL)
		}
		deadline, ok := req.Context().Deadline()
		if remaining := time.Until(deadline); !ok || remaining <= 29*time.Second || remaining > 30*time.Second {
			t.Error("expected default timeout of 30 seconds")
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{"status":"success"}`)), Header: make(http.Header)}, nil
	})
	client, err := iyzipay.New("sandbox-key", "secret")
	if err != nil {
		t.Fatal(err)
	}
	if _, err = client.BinControlRequest(context.Background(), &iyzipay.BinRequest{BinNumber: "454671"}); err != nil {
		t.Fatal(err)
	}
}

func paymentFixture() *iyzipay.PaymentRequest {
	return &iyzipay.PaymentRequest{
		Locale:         "tr",
		ConversationID: "123456789",
		Price:          "0.5",
		PaidPrice:      "1.0",
		Installment:    1,
		PaymentChannel: "WEB",
		BasketID:       "B67832",
		PaymentGroup:   "PRODUCT",
		PaymentCard: iyzipay.PaymentCard{
			CardHolderName: "John Doe",
			CardNumber:     "5528790000000008",
			ExpireYear:     "2030",
			ExpireMonth:    "12",
			CVC:            "123",
			RegisterCard:   0,
		},
		Buyer: iyzipay.Buyer{
			ID:                  "BY789",
			Name:                "John",
			Surname:             "Doe",
			IdentityNumber:      "74300864791",
			Email:               "email@email.com",
			GSMNumber:           "+905350000000",
			RegistrationDate:    "2013-04-21 15:12:09",
			LastLoginDate:       "2015-10-05 12:43:35",
			RegistrationAddress: "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
			City:                "Istanbul",
			Country:             "Turkey",
			ZipCode:             "34732",
			IP:                  "85.34.78.112",
		},
		BillingAddress: iyzipay.BillingAddress{
			Address:     "Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1",
			ContactName: "Jane Doe",
			City:        "Istanbul",
			Country:     "Turkey",
			ZipCode:     "34742",
		},
		BasketItems: []iyzipay.BasketItem{
			{
				ID:        "BI102",
				Name:      "Game code",
				Category1: "Game",
				Category2: "Online Game Items",
				ItemType:  "VIRTUAL",
				Price:     "0.5",
			},
		},
		Currency: "TRY",
	}
}
