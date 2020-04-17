package router

import (
	"beer_dmeo/service"
	"github.com/houzhongjian/beer"
)

func Init() beer.Engine {
	srv := beer.New()

	srv.Static("/img/", "./public/image")
	srv.Static("/conf/", "./public/conf/")

	beer.Loadini("./public/conf/app.ini")

	srv.GET("/", service.Default)
	srv.POST("/detail/:id", service.Detail)
	srv.GET("/login", service.Login)
	srv.GET("/rem", service.Rem)
	return srv
}
