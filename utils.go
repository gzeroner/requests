package requests

import url2 "net/url"

// URLWithQuery adds query parameters to the given URL.
func URLWithQuery(url string, query map[string]string) (string, error) {
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
