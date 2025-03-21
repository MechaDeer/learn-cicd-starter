package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		apiKey  string
		wantErr string
	}{
		{
			name: "Malformed header",
			headers: http.Header{
				"Authorization": []string{"Bearer valid_token"},
			},
			apiKey:  "",
			wantErr: "malformed authorization header",
		},
		{
			name: "Missing auth header",
			headers: http.Header{
				"Authorizatio": []string{"Bearer valid_token"},
			},
			apiKey:  "",
			wantErr: "no authorization header included",
		},
		{
			name: "Valid auth token",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid_token"},
			},
			apiKey:  "valid_token",
			wantErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tt.headers)
			if err != nil {
				if strings.Contains(err.Error(), tt.wantErr) {
					return
				}
				t.Errorf("GetBearerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotKey != tt.apiKey {
				t.Errorf("GetBearerToken() gotToken = %v, want %v", gotKey, tt.apiKey)
			}
		})
	}
}
