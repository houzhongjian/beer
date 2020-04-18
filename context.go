package beer

import "C"
import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"html/template"
	"strings"
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
	Header    http.Header
	templateDir string
	Layout   string
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

func (c *Context) Html(htmlPath string, data interface{}) {
	htmlPath = fmt.Sprintf("%s%s",c.templateDir, htmlPath)
	b, err := ioutil.ReadFile(htmlPath)
	if err != nil {
		log.Printf("err:%+v\n",err)
		return
	}
	var tmpl string
	if c.Layout == "" {
		tmpl = string(b)
	} else {
		layerByte, err := ioutil.ReadFile(fmt.Sprintf("%s%s",c.templateDir, c.Layout))
		if err != nil {
			log.Printf("err:%+v\n",err)
			return
		}
		tmpl = fmt.Sprintf(`{{define  "LayoutContent"}}%s{{end}}`, string(b))
		tmpl = strings.Replace(string(layerByte), `{{template "LayoutContent" .}}`, tmpl, -1)
	}
	t := template.New(htmlPath)
	t, err = t.Parse(tmpl)
	if err != nil {
		log.Printf("err:%+v\n",err)
		return
	}
	_ = t.Execute(c.Response, data)
}
