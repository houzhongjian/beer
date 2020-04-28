package main

import (
	_ "beer_dmeo/model"
	"beer_dmeo/router"
	"github.com/houzhongjian/beer"
	"github.com/houzhongjian/beer/session"
)
func main() {

	//加载配置文件.
	beer.Loadini("./public/conf/app.ini")

	//初始化session.
	beer.Session().Options(&session.Options{
		Drive: session.RedisDrive,
		Addr:  "127.0.0.1:6379",
		DB:    0,
		Password:  "",
	})

	//初始化路由.
	srv := router.Init()
	if err := srv.Run(":8080"); err != nil {
		panic(err)
	}
}
