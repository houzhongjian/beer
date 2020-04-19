package beer

//Use 使用中间件.
func (srv *Handler) Use(middleware ...beerFunc) {
	srv.middleware = append(srv.middleware, middleware...)
}
