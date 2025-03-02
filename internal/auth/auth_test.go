package auth

import (
	"net/http"
	"testing"
)

func TestApiKey(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "ApiKey efiwufiwefi3jfi2fi",
			expected: "efiwufiwefi3jfi2fi",
		},
		{
			input:    "",
			expected: ErrNoAuthHeaderIncluded.Error(),
		},
	}
	for _, c := range testCases {
		h := http.Header{}
		if c.input != "" {
			h.Add("Authorization", c.input)
		}
		key, err := GetAPIKey(h)
		if err != nil && err.Error() != c.expected {
			t.Error(err)
		}
		if key != c.expected {
			t.Errorf("expected: %s, actual: %s", c.expected, key)
		}
	}
}
