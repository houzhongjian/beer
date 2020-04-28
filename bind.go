package beer

import (
	"github.com/houzhongjian/beer/bind"
	"io/ioutil"
	"net/http"
)

func (c *Context) BindJSON(obj interface{}) error {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Response.WriteHeader(http.StatusBadRequest)
		return err
	}
	return bind.DecodeJSON(b, obj)
}

