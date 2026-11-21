package test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/JspBack/iyzipay-go"
)

func TestIyzilinkDeleteRequest(t *testing.T) {
	const path = "/v2/iyzilink/products/demo%2Ftoken"
	httpClient := &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodDelete || req.URL.EscapedPath() != path {
			t.Errorf("unexpected request: %s %s", req.Method, req.URL)
		}
		if query := req.URL.Query(); query.Get("locale") != "tr" || query.Get("conversationId") != "order + 1" {
			t.Errorf("unexpected query: %v", query)
		}
		if req.Body != nil && req.Body != http.NoBody {
			t.Error("delete request must have an empty body")
		}
		nonce := req.Header.Get("X-Iyzi-Rnd")
		mac := hmac.New(sha256.New, []byte("secret"))
		_, _ = mac.Write([]byte(nonce + path))
		payload := fmt.Sprintf("apiKey:key&randomKey:%s&signature:%x", nonce, mac.Sum(nil))
		want := "IYZWSv2 " + base64.StdEncoding.EncodeToString([]byte(payload))
		if req.Header.Get("Authorization") != want {
			t.Error("delete signature does not match the request path")
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"status":"success"}`)),
		}, nil
	})}
	client, err := iyzipay.New("key", "secret", iyzipay.WithHTTPClient(httpClient))
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.IyzilinkDelete(t.Context(), &iyzipay.IyzilinkDeleteRequest{
		Token: "demo/token", Locale: "tr", ConversationID: "order + 1",
	})
	if err != nil || response.Status != "success" {
		t.Fatalf("delete returned %+v, %v", response, err)
	}
}

func TestIyzilinkUpdateRequest(t *testing.T) {
	httpClient := &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPut || req.URL.EscapedPath() != "/v2/iyzilink/products/demo%2Ftoken" {
			t.Errorf("unexpected request: %s %s", req.Method, req.URL)
		}
		var payload map[string]any
		if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
			return nil, err
		}
		if _, ok := payload["token"]; ok {
			t.Error("token must be sent only in the URL path")
		}
		if _, ok := payload["encodedImageFile"]; ok {
			t.Error("an omitted image should not be sent")
		}
		if payload["name"] != "Updated product" || payload["addressIgnorable"] != false {
			t.Errorf("unexpected update body: %v", payload)
		}
		if _, ok := payload["installmentRequested"]; ok {
			t.Error("unspecified installment setting should not be sent")
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{"status":"success"}`))}, nil
	})}
	client, err := iyzipay.New("key", "secret", iyzipay.WithHTTPClient(httpClient))
	if err != nil {
		t.Fatal(err)
	}
	addressIgnorable := false
	_, err = client.IyzilinkUpdate(t.Context(), &iyzipay.IyzilinkUpdateRequest{
		Token: "demo/token", Name: "Updated product", Price: "2.00",
		Description: "Example", CurrencyCode: "TRY", AddressIgnorable: &addressIgnorable,
	})
	if err != nil {
		t.Fatal(err)
	}
}

func TestIyzilinkGetDetailRequest(t *testing.T) {
	httpClient := &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet || req.URL.EscapedPath() != "/v2/iyzilink/products/demo%2Ftoken" {
			t.Errorf("unexpected request: %s %s", req.Method, req.URL)
		}
		if query := req.URL.Query(); query.Get("locale") != "tr" || query.Get("conversationId") != "order + 1" {
			t.Errorf("unexpected query: %v", query)
		}
		if req.Body != nil && req.Body != http.NoBody {
			t.Error("detail request must have an empty body")
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"status":"success","conversationId":"order + 1","data":{"token":"demo/token","price":"2.50","presetPriceValues":["1.25",2]}}`)),
		}, nil
	})}
	client, err := iyzipay.New("key", "secret", iyzipay.WithHTTPClient(httpClient))
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.IyzilinkGetDetail(t.Context(), &iyzipay.IyzilinkGetDetailRequest{
		Token: "demo/token", Locale: "tr", ConversationID: "order + 1",
	})
	if err != nil {
		t.Fatal(err)
	}
	prices := response.Data.PresetPriceValues
	if response.ConversationID != "order + 1" || response.Data.Token != "demo/token" ||
		response.Data.Price != 2.5 || len(prices) != 2 || prices[0] != 1.25 || prices[1] != 2 {
		t.Fatalf("unexpected detail response: %+v", response)
	}
}

func TestIyzilinkListRequest(t *testing.T) {
	httpClient := &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodGet || req.URL.Path != "/v2/iyzilink/products" {
			t.Errorf("unexpected request: %s %s", req.Method, req.URL)
		}
		query := req.URL.Query()
		if query.Get("page") != "2" || query.Get("count") != "10" ||
			query.Get("locale") != "tr" || query.Get("conversationId") != "order + 1" {
			t.Errorf("unexpected query: %v", query)
		}
		if req.Body != nil && req.Body != http.NoBody {
			t.Error("list request must have an empty body")
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"status":"success","data":{"currentPage":2}}`)),
		}, nil
	})}
	client, err := iyzipay.New("key", "secret", iyzipay.WithHTTPClient(httpClient))
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.IyzilinkGetList(t.Context(), &iyzipay.IyzilinkGetListRequest{
		Locale: "tr", ConversationID: "order + 1", Page: 2, Count: 10,
	})
	if err != nil || response.Data.CurrentPage != 2 {
		t.Fatalf("list returned %+v, %v", response, err)
	}
}

func TestIyzilinkCreateFastLinkRequest(t *testing.T) {
	httpClient := &http.Client{Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
		if req.Method != http.MethodPost || req.URL.Path != "/v2/iyzilink/fast-link/products" {
			t.Errorf("unexpected request: %s %s", req.Method, req.URL)
		}
		var payload map[string]string
		if err := json.NewDecoder(req.Body).Decode(&payload); err != nil {
			return nil, err
		}
		if payload["price"] != "2.50" || payload["currencyCode"] != "TRY" || len(payload) != 2 {
			t.Errorf("unexpected Fast Link payload: %v", payload)
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Body:       io.NopCloser(strings.NewReader(`{"status":"success","data":{"token":"fast-token"}}`)),
		}, nil
	})}
	client, err := iyzipay.New("key", "secret", iyzipay.WithHTTPClient(httpClient))
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.IyzilinkCreateFastLink(t.Context(), &iyzipay.IyzilinkCreateFastLinkRequest{
		Price: "2.50", CurrencyCode: "TRY",
	})
	if err != nil || response.Data.Token != "fast-token" {
		t.Fatalf("Fast Link returned %+v, %v", response, err)
	}
}
