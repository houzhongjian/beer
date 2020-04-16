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
	srv.GET("/detail", Detail)
	srv.GET("/login/:userid/:name/:age", Login)
	if err := srv.Run(":8088"); err != nil {
		panic(err)
	}
}

func Default(c *beer.Context) {
	name := beer.Config().GetString("app_name")
	c.String(name)
}

func Login(c *beer.Context) {
	session, err := beer.Session().Start(c)
	if err != nil {
		log.Printf("err:%+v\n",err)
		return
	}
	session.Set("userid",c.Get("userid"))
	session.Set("name",c.Get("name"))
	session.Set("age",c.Get("age"))
}
func Detail(c *beer.Context) {
	session, err := beer.Session().Start(c)
	if err != nil {
		log.Printf("err:%+v\n",err)
		return
	}
	uid := session.Get("userid").(string)
	name := session.Get("name").(string)
	age := session.Get("age").(string)

	msg := fmt.Sprintf("uid = %s, name = %s, age = %s", uid, name, age)
	c.String(msg)
}