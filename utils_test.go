package requests

import "testing"

func TestURLWithQuery(t *testing.T) {
	// Happy path
	t.Run("HappyPath", func(t *testing.T) {
		url := "http://example.com"
		query := map[string]string{
			"key1": "value1",
			"key2": "value2",
		}
		expectedResult := "http://example.com?key1=value1&key2=value2"
		result, err := URLWithQuery(url, query)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
		if result != expectedResult {
			t.Errorf("Expected %s, but got %s", expectedResult, result)
		}
	})

	// Edge case: URL with no query parameters
	t.Run("NoQueryParams", func(t *testing.T) {
		url := "http://example.com"
		query := map[string]string{}
		expectedResult := "http://example.com"
		result, err := URLWithQuery(url, query)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
		if result != expectedResult {
			t.Errorf("Expected %s, but got %s", expectedResult, result)
		}
	})

	// Edge case: Nil query parameters
	t.Run("NilQueryParams", func(t *testing.T) {
		url := "http://example.com"
		var query map[string]string
		expectedResult := "http://example.com"
		result, err := URLWithQuery(url, query)
		if err != nil {
			t.Errorf("Expected no error, but got %v", err)
		}
		if result != expectedResult {
			t.Errorf("Expected %s, but got %s", expectedResult, result)
		}
	})

	// Edge case: Malformed URL
	t.Run("MalformedURL", func(t *testing.T) {
		url := ":"
		query := map[string]string{
			"key1": "value1",
		}
		_, err := URLWithQuery(url, query)
		if err == nil {
			t.Error("Expected an error, but got nil")
		}
	})
}
