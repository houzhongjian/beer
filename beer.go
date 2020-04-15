package beer

import (
	"net/http"
)

type Engine interface {
	Run(addr string) error
	GET(path string, handler beerHandlerFunc)
	POST(path string, handler beerHandlerFunc)
	DELETE(path string, handler beerHandlerFunc)
	PUT(path string, handler beerHandlerFunc)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type beerHandlerFunc func(*Context)

type Handler struct {
	router map[string]beerHandlerFunc
}

func New() Engine {
	e := new(Handler)
	e.router = make(map[string]beerHandlerFunc)
	return e
}

func (srv *Handler) Run(addr string) error {
	return http.ListenAndServe(addr, srv)
}

func (srv *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	funcName, ok := srv.router[r.RequestURI]
	if !ok {
		_, _ = w.Write([]byte("not found"))
		return
	}
	ctx := &Context{
		Request:  r,
		Response: w,
	}
	funcName(ctx)
}
