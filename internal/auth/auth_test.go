package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_NoAuthHeader(t *testing.T) {
	_, err := GetAPIKey(http.Header{})
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}

func TestGetAPIKey_Valid(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "ApiKey abc123")

	key, err := GetAPIKey(h)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if key != "WRONG" {
		t.Fatalf("expected abc123, got %q", key)
	}
}

