package beer

import "C"
import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//Context.
type Context struct {
	Method      string
	Request     *http.Request
	Response    http.ResponseWriter
	params      map[string]string
	UserAgent   string
	Url         string
	Body        io.ReadCloser
	Header      http.Header
	templateDir string
	Layout      string
	IP          string
	step        int
	Data        map[string]interface{}
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

func (c *Context) Html(htmlPath string) {
	htmlPath = fmt.Sprintf("%s%s", c.templateDir, htmlPath)
	b, err := ioutil.ReadFile(htmlPath)
	if err != nil {
		log.Printf("err:%+v\n", err)
		return
	}
	var tmpl string
	if c.Layout == "" {
		tmpl = string(b)
	} else {
		layerByte, err := ioutil.ReadFile(fmt.Sprintf("%s%s", c.templateDir, c.Layout))
		if err != nil {
			log.Printf("err:%+v\n", err)
			return
		}
		tmpl = fmt.Sprintf(`{{define  "LayoutContent"}}%s{{end}}{{template "LayoutContent" .}}`, string(b))
		tmpl = strings.Replace(string(layerByte), `{{template "LayoutContent" .}}`, tmpl, -1)
	}
	t := template.New(htmlPath)
	t, err = t.Parse(tmpl)
	if err != nil {
		log.Printf("err:%+v\n", err)
		return
	}
	_ = t.Execute(c.Response, c.Data)
}

//Json.
func (c *Context) Json() {
	b, err := json.Marshal(c.Data)
	if err != nil {
		log.Printf("err:%+v\n", err)
		return
	}
	c.Response.Header().Set("Content-Type", "application/json")
	_, _ = c.Response.Write(b)
}

//MiddlewareReturn 直接中断当前中间件执行流程.
func (c *Context) MiddlewareReturn() {
	c.step -= 1
}

