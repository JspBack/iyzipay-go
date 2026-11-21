package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// Authorization, gönderilecek baytları verilen tek kullanımlık rastgele değerle imzalar.
func Authorization(key, secret, path string, body []byte, nonce string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(nonce + path))
	mac.Write(body)
	value := fmt.Sprintf("apiKey:%s&randomKey:%s&signature:%x", key, nonce, mac.Sum(nil))
	return "IYZWSv2 " + base64.StdEncoding.EncodeToString([]byte(value))
}
