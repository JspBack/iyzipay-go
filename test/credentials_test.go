//go:build integration

package test

import (
	"os"
	"strings"
	"testing"
)

func sandboxCredentials(t *testing.T) (string, string) {
	t.Helper()
	key, secret := os.Getenv("IYZIPAY_API_KEY"), os.Getenv("IYZIPAY_API_SECRET")
	if key == "" || secret == "" {
		t.Skip("set IYZIPAY_API_KEY and IYZIPAY_API_SECRET for sandbox tests")
	}
	if !strings.HasPrefix(key, "sandbox-") {
		t.Fatal("integration tests require sandbox credentials")
	}
	return key, secret
}

func sandboxStoredCard(t *testing.T, prefix string) (string, string) {
	t.Helper()
	key := os.Getenv(prefix + "_USER_KEY")
	token := os.Getenv(prefix + "_TOKEN")
	if key == "" || token == "" {
		t.Skipf("set %s_USER_KEY and %s_TOKEN for this stored card test", prefix, prefix)
	}
	return key, token
}
