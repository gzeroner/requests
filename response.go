package requests

import (
	"encoding/json"
	"io"
	"net/http"
	"sync"
)

type Response struct {
	*http.Response
	body []byte
	once sync.Once
}

func (r *Response) readBody() {
	r.once.Do(func() {
		r.body, _ = io.ReadAll(r.Body)
		defer func() {
			_ = r.Body.Close()
		}()
	})
}

func (r *Response) String() string {
	r.readBody()

	return string(r.body)
}

func (r *Response) Unmarshal(v any) error {
	return json.Unmarshal(r.body, v)
}
