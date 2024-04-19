package requests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	url2 "net/url"
	"strings"
)

type Client interface {
	Request(config Config) (*Response, error)
	Get(url string, config ...Config) (*Response, error)
	Delete(url string, config ...Config) (*Response, error)
	Head(url string, config ...Config) (*Response, error)
	Options(url string, config ...Config) (*Response, error)
	Post(url string, data any, config ...Config) (*Response, error)
	Put(url string, data any, config ...Config) (*Response, error)
	Core() *http.Client
}

type client struct {
	httpClient *http.Client
}

func (c *client) Request(config Config) (*Response, error) {
	return c.request(config)
}

func (c *client) Get(url string, config ...Config) (*Response, error) {
	config = append(config, Config{Method: http.MethodGet, Url: url})
	conf := mergeConfig(config...)
	return c.request(conf)
}

func (c *client) Delete(url string, config ...Config) (*Response, error) {
	config = append(config, Config{Method: http.MethodDelete, Url: url})
	conf := mergeConfig(config...)
	return c.request(conf)
}

func (c *client) Head(url string, config ...Config) (*Response, error) {
	config = append(config, Config{Method: http.MethodHead, Url: url})
	conf := mergeConfig(config...)
	return c.request(conf)
}

func (c *client) Options(url string, config ...Config) (*Response, error) {
	config = append(config, Config{Method: http.MethodOptions, Url: url})
	conf := mergeConfig(config...)
	return c.request(conf)
}

func (c *client) Post(url string, data any, config ...Config) (*Response, error) {
	config = append(config, Config{Method: http.MethodPost, Url: url, Data: data})
	conf := mergeConfig(config...)
	return c.request(conf)
}

func (c *client) Put(url string, data any, config ...Config) (*Response, error) {
	config = append(config, Config{Method: http.MethodPut, Url: url, Data: data})
	conf := mergeConfig(config...)
	return c.request(conf)
}

func (c *client) Core() *http.Client {
	return c.httpClient
}

func (c *client) request(config Config) (*Response, error) {
	var err error

	// check config
	if err = config.Check(); err != nil {
		return nil, err
	}

	// request address
	addr := config.Url
	if config.BaseUrl != "" {
		addr, err = url2.JoinPath(config.BaseUrl, addr)
		if err != nil {
			return nil, err
		}
	}

	// query parameters
	if config.Query != nil {
		addr, err = URLWithQuery(addr, config.Query)
	}

	// body
	body, err := c.handleBody(config)
	if err != nil {
		return nil, err
	}

	// create request
	config.Method = strings.ToUpper(config.Method)
	req, err := http.NewRequest(config.Method, addr, body)
	if err != nil {
		return nil, err
	}

	// request headers
	if config.Headers != nil {
		for k, v := range config.Headers {
			req.Header.Set(k, v)
		}
	}

	// auth
	if config.Auth.Username != "" && config.Auth.Password != "" {
		req.SetBasicAuth(config.Auth.Username, config.Auth.Password)
	}

	// send request
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	// response
	return &Response{Response: res}, nil
}

func (c *client) handleBody(config Config) (io.Reader, error) {
	if config.Data == nil {
		switch config.Data.(type) {
		case string:
			return strings.NewReader(config.Data.(string)), nil
		default:
			bs, err := json.Marshal(config.Data)
			if err != nil {
				return nil, err
			}
			return bytes.NewReader(bs), nil
		}
	}
	return nil, nil
}
