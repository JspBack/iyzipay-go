package transport

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/JspBack/iyzipay-go/internal/auth"
)

// APIError, HTTP veya ödeme sağlayıcısı hatasını tanımlar. Ham yanıt gövdesini içermez.
type APIError struct {
	HTTPStatus     int    `json:"-"`
	Code           string `json:"errorCode"`
	Message        string `json:"errorMessage"`
	Group          string `json:"errorGroup"`
	ConversationID string `json:"conversationId"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("iyzipay: HTTP %d: %s: %s", e.HTTPStatus, e.Code, e.Message)
}

// Send, imzalı isteği gönderir; HTTP durumunu ve ödeme sağlayıcısının işlem sonucunu denetler.
func Send(ctx context.Context, client *http.Client, method, baseURL, key, secret, path string, body []byte, headers http.Header) ([]byte, error) {
	nonce := rand.Text()
	req, err := http.NewRequestWithContext(ctx, method, baseURL+path, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("iyzipay: request: %w", err)
	}
	req.Header = headers.Clone()
	if req.Header == nil {
		req.Header = make(http.Header)
	}
	req.Header.Set("Content-Type", "application/json")
	signedPath, _, _ := strings.Cut(path, "?")
	req.Header.Set("Authorization", auth.Authorization(key, secret, signedPath, body, nonce))
	req.Header.Set("X-Iyzi-Rnd", nonce)
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("iyzipay: send request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()
	data, err := io.ReadAll(io.LimitReader(resp.Body, maxBody+1))
	if err != nil {
		return nil, fmt.Errorf("iyzipay: read response: %w", err)
	}
	if len(data) > maxBody {
		return nil, fmt.Errorf("iyzipay: response exceeds %d bytes", maxBody)
	}
	var envelope struct {
		Status string `json:"status"`
		APIError
	}
	decodeErr := json.Unmarshal(data, &envelope)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		envelope.HTTPStatus = resp.StatusCode
		if decodeErr != nil || envelope.Message == "" {
			envelope.Message = http.StatusText(resp.StatusCode)
		}
		return nil, &envelope.APIError
	}
	if decodeErr != nil {
		return nil, fmt.Errorf("iyzipay: decode response: %w", decodeErr)
	}
	if envelope.Status == "failure" || envelope.Code != "" {
		envelope.HTTPStatus = resp.StatusCode
		return nil, &envelope.APIError
	}
	if envelope.Status != "success" {
		return nil, fmt.Errorf("iyzipay: missing or invalid response status")
	}
	return data, nil
}
