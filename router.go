package beer

import (
	"os"
	"path/filepath"
)

func (srv *Handler) Static(path string, dir string) {
	srv.createFileServer(path, dir)
}

func (srv *Handler) createFileServer(fpath string, dir string) {
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			srv.fsRouter[fpath+info.Name()] = path
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func (srv *Handler) GET(path string, handler hFunc) {
	srv.handle("GET", path, handler)
}

func (srv *Handler) POST(path string, handler hFunc) {
	srv.handle("POST", path, handler)
}

func (srv *Handler) DELETE(path string, handler hFunc) {
	srv.handle("DELETE", path, handler)
}

func (srv *Handler) PUT(path string, handler hFunc) {
	srv.handle("PUT", path, handler)
}

func (srv *Handler) handle(method string, path string, handler hFunc) {
	h := beerHandler{
		Method: method,
		Path:   path,
	}
	_, ok := srv.router[h]
	if ok {
		panic("当前url已经存在")
	}
	srv.router[h] = handler
}
