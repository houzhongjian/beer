package main

import (
	"beer"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	srv := beer.New()

	beer.Loadini("./app.ini")

	srv.GET("/", Default)
	if err := srv.Run(":8088"); err != nil {
		panic(err)
	}
}

func Default(c *beer.Context) {
	name := beer.Config().GetString("app_name")
	c.String(name)
}