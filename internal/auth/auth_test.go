package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type args struct {
		headers http.Header
	}
	tests := []struct {
		name string
		args args
		want string
		err  error
	}{
		{"ValidHeader", args{http.Header{"Authorization": []string{"ApiKey abc123"}}}, "abc123", nil},
		{"InvalidHeader", args{http.Header{"Authorization": []string{"Basic abc123"}}}, "", ErrMalformedAuthHeader},
		{"MalformedHeader", args{http.Header{"Authorization": []string{"ApiKey"}}}, "", ErrMalformedAuthHeader},
		{"EmptyHeader", args{http.Header{}}, "", ErrNoAuthHeaderIncluded},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.args.headers)
			if !errors.Is(err, tt.err) {
				t.Errorf("GetAPIKey() error = %v, wanted error = %v", err, tt.err)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = (%v, %v), want (%v, %v)", got, err, tt.want, tt.err)
			}
		})
	}
}
