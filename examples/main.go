package main

import (
	"beer"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	srv := beer.New()

	beer.Loadini("./app.ini")

	srv.GET("/", Default)
	srv.GET("/detail/:id", Detail)
	if err := srv.Run(":8088"); err != nil {
		panic(err)
	}
}

func Default(c *beer.Context) {
	name := beer.Config().GetString("app_name")
	c.String(name)
}
func Detail(c *beer.Context) {
	id := c.Get("id")
	msg := fmt.Sprintf("id = %s",id)
	c.String(msg)
}