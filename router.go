package beer

import (
	"os"
	"path/filepath"
)

func (srv *Handler) Static(path string, dir string) {
	//判断dir路径是否存在.
	_, err := os.Stat(dir)
	if err != nil {
		panic(err)
	}
	srv.fsHandle(path, dir)
}

func (srv *Handler) fsHandle(fpath string, dir string) {
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

func (srv *Handler) GET(path string, handler beerFunc) {
	srv.handle("GET", path, handler)
}

func (srv *Handler) POST(path string, handler beerFunc) {
	srv.handle("POST", path, handler)
}

func (srv *Handler) DELETE(path string, handler beerFunc) {
	srv.handle("DELETE", path, handler)
}

func (srv *Handler) PUT(path string, handler beerFunc) {
	srv.handle("PUT", path, handler)
}

func (srv *Handler) handle(method string, path string, handler beerFunc) {
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
