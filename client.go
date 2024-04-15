package requests

import (
	"net/http"
	neturl "net/url"
)

type Client interface {
	Request(config Config) (Response, error)
	Get(url string, config Config) (Response, error)
	Delete(url string, config Config) (Response, error)
	Head(url string, config Config) (Response, error)
	Options(url string, config Config) (Response, error)
	Post(url string, data any, config Config) (Response, error)
	Put(url string, data any, config Config) (Response, error)
	Core() *http.Client
}

type client struct {
	httpClient *http.Client
}

func (c *client) Core() *http.Client {
	return c.httpClient
}

func (c *client) Request(config Config) (Response, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) Get(url string, config Config) (Response, error) {
	if len(config.Query) > 0 {
		vs := neturl.Values{}
		for k, _ := range config.Query {
			vs.Add(k, config.Query[k])
		}
		url += "?" + vs.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Response{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Response{}, err
	}

	return Response{res}, nil
}

func (c *client) Delete(url string, config Config) (Response, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) Head(url string, config Config) (Response, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) Options(url string, config Config) (Response, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) Post(url string, data any, config Config) (Response, error) {
	//TODO implement me
	panic("implement me")
}

func (c *client) Put(url string, data any, config Config) (Response, error) {
	//TODO implement me
	panic("implement me")
}
