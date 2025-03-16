package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		headers := http.Header{}
		headers.Add("Authorization", "ApiKey 1234567890")
		apiKey, err := GetAPIKey(headers)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if apiKey != "1234567890" {
			t.Errorf("unexpected api key: %s", apiKey)
		}
	})

	t.Run("no auth header", func(t *testing.T) {
		headers := http.Header{}
		apiKey, err := GetAPIKey(headers)
		if err == nil {
			t.Errorf("expected error, got: %v", apiKey)
		}
		if err != ErrNoAuthHeaderIncluded {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
