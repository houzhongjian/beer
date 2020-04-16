package beer

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