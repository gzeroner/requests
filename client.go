package requests

import (
	"net/http"
	url2 "net/url"
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
	//TODO implement me
	panic("implement me")
}

func (c *client) Get(url string, config ...Config) (*Response, error) {
	if len(config) > 0 {
		if config[0].Query != nil {
			var err error
			if url, err = c.autoCompleteQuery(url, config[0].Query); err != nil {
				return nil, err
			}
		}
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	return &Response{Response: res}, nil
}

func (c *client) Delete(url string, config ...Config) (*Response, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) Head(url string, config ...Config) (*Response, error) {
	if len(config) > 0 {
		if config[0].Query != nil {
			var err error
			if url, err = c.autoCompleteQuery(url, config[0].Query); err != nil {
				return nil, err
			}
		}
	}

	res, err := c.httpClient.Head(url)
	if err != nil {
		return nil, err
	}
	return &Response{Response: res}, nil
}

func (c *client) Options(url string, config ...Config) (*Response, error) {
	if len(config) > 0 {
		if config[0].Query != nil {
			var err error
			if url, err = c.autoCompleteQuery(url, config[0].Query); err != nil {
				return nil, err
			}
		}
	}

	req, err := http.NewRequest(http.MethodOptions, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return &Response{Response: res}, nil
}

func (c *client) Post(url string, data any, config ...Config) (*Response, error) {
	if len(config) > 0 {
		if config[0].Query != nil {
			var err error
			if url, err = c.autoCompleteQuery(url, config[0].Query); err != nil {
				return nil, err
			}
		}
	}

	if data != nil {

	}

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return nil, err
	}
	// default content-type
	req.Header.Set("Content-Type", "application/json")
	if len(config) > 0 && config[0].Headers != nil {
		for k, v := range config[0].Headers {
			req.Header.Set(k, v)
		}
	}

	return nil, err
}

func (c *client) Put(url string, data any, config ...Config) (*Response, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) Core() *http.Client {
	//TODO implement me
	panic("implement me")
}

func (c *client) request() (*Response, error) {
	return nil, nil
}

// autoCompleteQuery 自动补全 url 查询参数
func (c *client) autoCompleteQuery(url string, query map[string]string) (string, error) {
	addr, err := url2.Parse(url)
	if err != nil {
		return "", err
	}

	if query == nil {
		return url, nil
	}

	vs := addr.Query()
	for k, v := range query {
		vs.Set(k, v)
	}
	addr.RawQuery = vs.Encode()
	return addr.String(), nil
}
