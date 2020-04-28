package router

import (
	"beer_dmeo/filter"
	"beer_dmeo/service"
	"github.com/houzhongjian/beer"
)

func Init() beer.Engine {
	srv := beer.New()
	srv.Use(filter.Log)

	srv.Static("/img/", "./public/image")
	srv.Static("/conf/", "./public/conf/")
	srv.Static("/css/", "./public/css/")
	srv.Static("/js/", "./public/js/")

	srv.SetTemplateDir("./views/")

	srv.GET("/login", service.Login)
	srv.POST("/login", service.Login)

	srv.Use(filter.FilterLogin)
	srv.GET("/detail/:id", service.Detail)
	srv.GET("/", service.Default)

	srv.GET("/rem", service.Rem)
	return srv
}
