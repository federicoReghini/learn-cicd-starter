package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		headers   http.Header
		expected  string
		expectErr error
	}{
		{
			name:      "Valid API Key",
			headers:   http.Header{"Authorization": {"ApiKey abc123"}},
			expected:  "abc123",
			expectErr: nil,
		},
		{
			name:      "Missing Authorization Header",
			headers:   http.Header{},
			expected:  "",
			expectErr: ErrNoAuthHeaderIncluded,
		},
		{
			name:      "Malformed Authorization Header",
			headers:   http.Header{"Authorization": {"Bearer abc123"}},
			expected:  "",
			expectErr: errors.New("malformed authorization header"),
		},
		{
			name:      "Empty Authorization Header",
			headers:   http.Header{"Authorization": {""}},
			expected:  "",
			expectErr: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tt.headers)
			if apiKey != tt.expected || (err != nil && err.Error() != tt.expectErr.Error()) {
				t.Errorf("expected (%v, %v), got (%v, %v)", tt.expected, tt.expectErr, apiKey, err)
			}
		})
	}
}
