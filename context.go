package beer

import (
	"net/http"
)

//Context.
type Context struct {
	Method   string
	Request  *http.Request
	Response http.ResponseWriter
	Params   map[string]string
}

func (c *Context) String(msg string) {
	_, _ = c.Response.Write([]byte(msg))
}

//Get.
func (c *Context) Get(key string) string {
	v, ok := c.Params[key]
	if !ok {
		return ""
	}
	return v
}
