package test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestSharedValidationRules(t *testing.T) {
	client, err := iyzipay.New("key", "secret", iyzipay.WithHTTPClient(&http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		t.Error("invalid request reached HTTP transport")
		return nil, errors.New("unexpected request")
	})}))
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name, field, tag string
		mutate           func(*iyzipay.InitTDSRequest)
	}{
		{"date", "RegistrationDate", "iyziDate", func(r *iyzipay.InitTDSRequest) { r.Buyer.RegistrationDate = "2026-99-99" }},
		{"card", "CardNumber", "creditcard", func(r *iyzipay.InitTDSRequest) { r.PaymentCard.CardNumber = "5528790000000009" }},
		{"callback", "CallbackUrl", "httpsurl", func(r *iyzipay.InitTDSRequest) { r.CallbackUrl = "http://example.com/callback" }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			request := &iyzipay.InitTDSRequest{PaymentRequest: *paymentFixture(), CallbackUrl: "https://example.com/callback"}
			tc.mutate(request)
			_, _, err := client.InitializeTDSPayment(context.Background(), request)
			if err == nil || !strings.Contains(err.Error(), "field '"+tc.field+"' with tag '"+tc.tag+"'") {
				t.Fatalf("unexpected validation error: %v", err)
			}
			if errors.Unwrap(err) == nil {
				t.Fatal("validation cause was lost")
			}
		})
	}
}

func TestConcurrentValidationAndClientReuse(t *testing.T) {
	var calls atomic.Int32
	client, err := iyzipay.New("key", "secret", iyzipay.WithHTTPClient(&http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		calls.Add(1)
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{"status":"success"}`)), Header: make(http.Header)}, nil
	})}))
	if err != nil {
		t.Fatal(err)
	}
	const workers = 16
	const iterations = 4
	var wg sync.WaitGroup
	for range workers {
		wg.Go(func() {
			for range iterations {
				if _, err := client.CreatePayment(context.Background(), paymentFixture()); err != nil {
					t.Error(err)
				}
				if _, err := client.BinControlRequest(context.Background(), &iyzipay.BinRequest{BinNumber: "invalid"}); err == nil || !strings.HasPrefix(err.Error(), "validation errors:") {
					t.Errorf("unexpected validation result: %v", err)
				}
			}
		})
	}
	wg.Wait()
	if got := calls.Load(); got != workers*iterations {
		t.Fatalf("HTTP calls=%d want=%d", got, workers*iterations)
	}
}

func TestCheckoutFormTypes(t *testing.T) {
	for _, formType := range []string{"responsive", "popup", "iframe", "invalid"} {
		t.Run(formType, func(t *testing.T) {
			calls := 0
			const pageURL = "https://example.com/payment?token=1"
			client, err := iyzipay.New("key", "secret", iyzipay.WithHTTPClient(&http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
				calls++
				return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{"status":"success","paymentPageUrl":"` + pageURL + `"}`)), Header: make(http.Header)}, nil
			})}))
			if err != nil {
				t.Fatal(err)
			}
			fixture := paymentFixture()
			request := &iyzipay.CFRequest{InitPWIRequest: iyzipay.InitPWIRequest{
				Price: fixture.Price, PaidPrice: fixture.PaidPrice, Currency: fixture.Currency,
				Buyer: fixture.Buyer, BillingAddress: fixture.BillingAddress, BasketItems: fixture.BasketItems,
				CallbackUrl: "https://example.com/callback",
			}}
			response, err := client.CheckoutFormPaymentRequest(context.Background(), request, formType)
			if formType == "invalid" {
				if err == nil || calls != 0 {
					t.Fatalf("invalid form: calls=%d err=%v", calls, err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			want := pageURL
			if formType == "iframe" {
				want += "&iframe=true"
			}
			if response.PaymentPageUrl != want || calls != 1 {
				t.Fatalf("URL=%q calls=%d", response.PaymentPageUrl, calls)
			}
		})
	}
}
