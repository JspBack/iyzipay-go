package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// İstek oluşturmak için kullanılan fonksiyon
func DoRequest(requestData []byte, client *http.Client, method, baseURI, apiKey, secretKey, uriPath string) (response []byte, err error) {
	randomStr := randString(8)
	authorizationString := generateAuthorizationString(apiKey, secretKey, string(requestData), uriPath)

	if !AcceptedRequestMethods[method] {
		return nil, fmt.Errorf("invalid request method: %s", method)
	}

	req, err := http.NewRequest(method, baseURI+uriPath, bytes.NewBuffer(requestData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authorizationString)
	req.Header.Set("x-iyzi-rnd", randomStr)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

// n uzunluğunda random string oluşturur
func randString(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	sb := strings.Builder{}
	sb.Grow(n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

// HMACSHA256 ile Iyzipay API için yetkilendirme stringi oluşturur (SHA1 artık desteklenmiyor) (bugün değil ama yakında :D)
func generateAuthorizationString(apiKey, secretKey, requestData, uriPath string) string {
	randomKey := randString(8)

	payload := randomKey + uriPath
	if requestData != "" {
		payload += requestData
	}

	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(payload))
	encryptedData := mac.Sum(nil)

	authorizationString := fmt.Sprintf("apiKey:%s&randomKey:%s&signature:%s",
		apiKey, randomKey, fmt.Sprintf("%x", encryptedData))

	base64EncodedAuthorization := base64.StdEncoding.EncodeToString([]byte(authorizationString))

	return "IYZWSv2 " + base64EncodedAuthorization
}

// Base64 decode işlemi (string alıp string döndürür)
func Base64Decode(data string) (res string, err error) {
	ba, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}

	return string(ba), nil
}
