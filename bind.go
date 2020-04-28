package beer

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (c *Context) BindJSON(obj interface{}) error {
	b, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.Response.WriteHeader(http.StatusBadRequest)
		return err
	}
	if err := decodeJSON(b, obj); err != nil {
		c.Response.WriteHeader(http.StatusBadRequest)
		return err
	}
	return nil
}

func decodeJSON(body []byte, obj interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(body))
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}
