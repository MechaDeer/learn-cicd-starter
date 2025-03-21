package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		apiKey  string
		wantErr bool
	}{
		{
			name: "Malformed header",
			headers: http.Header{
				"Authorization": []string{"Bearer valid_token"},
			},
			apiKey:  "",
			wantErr: true,
		},
		{
			name: "Missing auth header",
			headers: http.Header{
				"Authorizatio": []string{"Bearer valid_token"},
			},
			apiKey:  "",
			wantErr: true,
		},
		{
			name: "Valid auth token",
			headers: http.Header{
				"Authorization": []string{"ApiKey valid_token"},
			},
			apiKey:  "valid_token",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBearerToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotKey != tt.apiKey {
				t.Errorf("GetBearerToken() gotToken = %v, want %v", gotKey, tt.apiKey)
			}
		})
	}
}
