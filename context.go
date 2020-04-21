package beer

import "C"
import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
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
	Layout      string
	IP          string
	step        int	//步长.
	Data        map[string]interface{}	//视图渲染的数据.
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
	tmpl, ok := templateData[htmlPath]
	if !ok {
		log.Printf("当前视图文件不存在:", htmlPath)
		return
	}

	if c.Layout != "" {
		layoutContent, ok := templateData[c.Layout]
		if !ok {
			log.Printf("当前视图文件不存在:", c.Layout)
			return
		}
		tmpl = fmt.Sprintf(`{{define  "LayoutContent"}}%s{{end}}{{template "LayoutContent" .}}`, tmpl)
		tmpl = strings.Replace(layoutContent, `{{template "LayoutContent" .}}`, tmpl, -1)
	}
	t := template.New(htmlPath)
	t, err := t.Parse(tmpl)
	if err != nil {
		log.Printf("err:%+v\n", err)
		return
	}
	c.Response.Header().Set("Content-Type","text/html;charset=utf-8")
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
