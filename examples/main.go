package main

import (
	"beer_dmeo/router"
	"github.com/houzhongjian/beer"
	"github.com/houzhongjian/beer/log"
	"github.com/houzhongjian/beer/session"
)
func main() {
	//加载配置文件.
	beer.Config().Loadini("./public/conf/app.ini")

	cf := beer.Config()
	//初始化session.
	beer.Session().Options(&session.Options{
		Drive: session.RedisDrive,
		Addr:  cf.GetString("redis_addr"),
		DB:    0,
		Password:  "",
	})

	//初始化log.
	beer.Log().Options(&log.Options{})

	//初始化路由.
	srv := router.Init()
	if err := srv.Run(":8080"); err != nil {
		panic(err)
	}
}
