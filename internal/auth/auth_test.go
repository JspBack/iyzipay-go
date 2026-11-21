package auth

import "testing"

func TestAuthorizationVector(t *testing.T) {
	got := Authorization("key", "secret", "/payment/bin/check", []byte(`{"binNumber":"454671"}`), "nonce")
	const want = "IYZWSv2 YXBpS2V5OmtleSZyYW5kb21LZXk6bm9uY2Umc2lnbmF0dXJlOjlhZDk5MDI2OTBmMWZmMWZkZTgyYTExYWQ4OWRiMGI4MDY5N2Q3MjY5YWQxZWRkOWZmNGFmZTIyNjhkNWFkMDA="
	if got != want {
		t.Fatalf("got %s; want %s", got, want)
	}
}
