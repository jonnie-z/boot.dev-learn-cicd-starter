package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "ApiKey abc123", expected: "abc123"},
	}

	for _, tc := range tests {
		headerToTest := http.Header{}
		headerToTest.Set("Authorization", tc.input)
		expected := "abc123"

		actual, err := GetAPIKey(headerToTest)
		if err != nil {
			t.Errorf("GetAPIKey should not have thrown: %s\n", err)
		}

		if actual != expected {
			t.Errorf("Actual (%s) should be equivalent to Expected (%s)\n", actual, expected)
		}
	}
}
