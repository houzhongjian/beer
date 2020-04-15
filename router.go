package beer

func (srv *Handler) GET(path string, handler beerHandlerFunc) {
	_, ok := srv.router[path]
	if ok {
		panic("当前url已经存在")
	}
	srv.router[path] = handler
}

func (srv *Handler) POST(path string, handler beerHandlerFunc) {
	_, ok := srv.router[path]
	if ok {
		panic("当前url已经存在")
	}
	srv.router[path] = handler
}

func (srv *Handler) DELETE(path string, handler beerHandlerFunc) {
	_, ok := srv.router[path]
	if ok {
		panic("当前url已经存在")
	}
	srv.router[path] = handler
}

func (srv *Handler) PUT(path string, handler beerHandlerFunc) {
	_, ok := srv.router[path]
	if ok {
		panic("当前url已经存在")
	}
	srv.router[path] = handler
}
