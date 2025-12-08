package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	cases := []struct {
		headers        http.Header
		expectedResult string
		wantErr        bool
	}{
		{
			headers:        http.Header{},
			expectedResult: "",
			wantErr:        true,
		},
		{
			headers: http.Header{
				"Authorization": {"this is a malformed api key"},
			},
			expectedResult: "",
			wantErr:        true,
		},
		{
			headers: http.Header{
				"Authorization": {"ApiKey thisismyapikey"},
			},
			expectedResult: "thisismyapikey",
			wantErr:        false,
		},
	}

	for _, tc := range cases {
		actualResult, actualErr := GetAPIKey(tc.headers)
		if actualResult != tc.expectedResult {
			t.Errorf("expected result: %s, got %s", actualResult, tc.expectedResult)
		}
		if tc.wantErr != (actualErr != nil) {
			t.Errorf("expected error: %v, got %v", actualErr, tc.wantErr)
		}
	}
}
