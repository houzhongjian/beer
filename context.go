package beer

import (
	"io"
	"net/http"
)

//Context.
type Context struct {
	Method    string
	Request   *http.Request
	Response  http.ResponseWriter
	params    map[string]string
	UserAgent string
	Url       string
	Body      io.ReadCloser
}

func (c *Context) String(msg string) {
	_, _ = c.Response.Write([]byte(msg))
}

//Param.
func (c *Context) Param(key string) string {
	v, ok := c.params[key]
	if !ok {
		return ""
	}
	return v
}
