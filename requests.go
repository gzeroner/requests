package requests

import "net/http"

var DefaultClient = New()

func New() Client {
	return &client{
		httpClient: &http.Client{},
	}
}

func Request(config Config) (*Response, error) {
	return DefaultClient.Request(config)
}

func Get(url string, config ...Config) (*Response, error) {
	return DefaultClient.Get(url, config...)
}

func Delete(url string, config ...Config) (*Response, error) {
	return DefaultClient.Delete(url, config...)
}

func Head(url string, config ...Config) (*Response, error) {
	return DefaultClient.Head(url, config...)
}

func Options(url string, config ...Config) (*Response, error) {
	return DefaultClient.Options(url, config...)
}

func Post(url string, data any, config ...Config) (*Response, error) {
	return DefaultClient.Post(url, data, config...)
}

func Put(url string, data any, config ...Config) (*Response, error) {
	return DefaultClient.Put(url, data, config...)
}
