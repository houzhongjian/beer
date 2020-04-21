package beer

import (
	"net/http"
)

type Engine interface {
	Run(addr string) error
	GET(path string, handler beerFunc)
	POST(path string, handler beerFunc)
	DELETE(path string, handler beerFunc)
	PUT(path string, handler beerFunc)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Static(path string, dir string)
	SetTemplateDir(dir string)
	Use(middleware ...beerFunc)
}

type beerFunc func(*Context)

type beerHandler struct {
	Method string
	Path   string
}

type Handler struct {
	router           map[beerHandler]beerFunc
	fsRouter         map[string]string
	templateDir      string
	templateData      map[string]string
	middleware       []beerFunc
	middlewareRouter map[string][]beerFunc
}

func New() Engine {
	e := new(Handler)
	e.router = make(map[beerHandler]beerFunc)
	e.fsRouter = make(map[string]string)
	e.middlewareRouter = make(map[string][]beerFunc)
	e.templateData = make(map[string]string)
	return e
}
