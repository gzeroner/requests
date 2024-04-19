package requests

import (
	"fmt"
	"testing"
)

func TestHTTPRequestMethods(t *testing.T) {
	// Test cases for happy path
	t.Run("TestRequest", func(t *testing.T) {
		resp, err := Request(Config{Url: "https://bing.com", Method: "get"})
		if err != nil {
			t.Errorf("Request() returned error: %v", err)
		}
		if resp == nil {
			t.Error("Request() returned nil response")
		}
	})

	t.Run("TestGet", func(t *testing.T) {
		resp, err := Get("https://bing.com")
		if err != nil {
			t.Errorf("Get() returned error: %v", err)
		}
		if resp == nil {
			t.Error("Get() returned nil response")
		}
	})

	t.Run("TestDelete", func(t *testing.T) {
		resp, err := Delete("https://bing.com")
		if err != nil {
			t.Errorf("Delete() returned error: %v", err)
		}
		if resp == nil {
			t.Error("Delete() returned nil response")
		}
	})

	t.Run("TestHead", func(t *testing.T) {
		resp, err := Head("https://bing.com")
		if err != nil {
			t.Errorf("Head() returned error: %v", err)
		}
		if resp == nil {
			t.Error("Head() returned nil response")
		}
	})

	t.Run("TestOptions", func(t *testing.T) {
		resp, err := Options("https://bing.com")
		if err != nil {
			t.Errorf("Options() returned error: %v", err)
		}
		if resp == nil {
			t.Error("Options() returned nil response")
		}
	})

	t.Run("TestPost", func(t *testing.T) {
		resp, err := Post("https://bing.com", "data", Config{})
		if err != nil {
			t.Errorf("Post() returned error: %v", err)
		}
		if resp == nil {
			t.Error("Post() returned nil response")
		}
		fmt.Println(resp.String())
	})

	t.Run("TestPut", func(t *testing.T) {
		resp, err := Put("https://bing.com", "data", Config{})
		if err != nil {
			t.Errorf("Put() returned error: %v", err)
		}
		if resp == nil {
			t.Error("Put() returned nil response")
		}
	})
}
