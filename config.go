package requests

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
