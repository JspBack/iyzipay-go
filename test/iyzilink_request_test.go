package test

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
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
