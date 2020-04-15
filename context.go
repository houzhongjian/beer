package beer

import (
	"net/http"
)

//Context.
type Context struct {
	Request  *http.Request
	Response http.ResponseWriter
}

func (c *Context) String(msg string) {
	_, _ = c.Response.Write([]byte(msg))
}
