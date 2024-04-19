package requests

import "strings"

// methods allow methods
var methods = map[string]struct{}{
	"GET":     {},
	"POST":    {},
	"PUT":     {},
	"DELETE":  {},
	"HEAD":    {},
	"OPTIONS": {},
}

// BasicAuth is the basic authentication struct
type BasicAuth struct {
	Username string
	Password string
}

// Config is the configuration for a request
type Config struct {
	// Url is the URL to send the request to
	Url string
	// Method is the HTTP method to use for the request
	Method string
	// BaseUrl is the base URL to use for relative URLs in the request
	BaseUrl string
	// Headers is a map of headers to send with the request
	Headers map[string]string
	// Auth is the basic authentication credentials to use with the request
	Auth BasicAuth
	// Query is a map of query parameters to send with the request
	Query map[string]string
	// Form is a map of form parameters to send with the request
	Form map[string]any
	// Data is the data to send with the request (e.g. JSON payload)
	Data any
}

// Check checks the configuration for errors and returns an error if any
func (c Config) Check() error {
	if c.Url == "" && c.BaseUrl == "" {
		return ErrIllegalUrl
	}

	if _, ok := methods[strings.ToUpper(c.Method)]; !ok {
		return ErrMethodNotAllowed
	}

	return nil
}

// mergeConfig merges multiple Config objects into a single Config object
func mergeConfig(configs ...Config) Config {
	switch len(configs) {
	case 0:
		return Config{}
	case 1:
		return configs[0]
	default:
		merged := configs[0]
		for i := 1; i < len(configs); i++ {
			if configs[i].Url != "" {
				merged.Url = configs[i].Url
			}
			if configs[i].Method != "" {
				merged.Method = configs[i].Method
			}
			if configs[i].BaseUrl != "" {
				merged.BaseUrl = configs[i].BaseUrl
			}
			if len(configs[i].Headers) > 0 {
				merged.Headers = configs[i].Headers
			}
			if configs[i].Auth.Username != "" {
				merged.Auth.Username = configs[i].Auth.Username
			}
			if configs[i].Auth.Password != "" {
				merged.Auth.Password = configs[i].Auth.Password
			}
			if len(configs[i].Query) > 0 {
				merged.Query = configs[i].Query
			}
			if len(configs[i].Form) > 0 {
				merged.Form = configs[i].Form
			}
			if configs[i].Data != nil {
				merged.Data = configs[i].Data
			}
		}
		return merged
	}
}
