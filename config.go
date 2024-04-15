package requests

// BasicAuth is the basic authentication struct
type BasicAuth struct {
	Username string
	Password string
}

// Config is the configuration for a request
type Config struct {
	Url     string
	Method  string
	BaseUrl string
	Headers map[string]string
	Query   map[string]string
	Data    any
	Auth    BasicAuth
}
