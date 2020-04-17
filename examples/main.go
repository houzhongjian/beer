package main

import (
	"github.com/houzhongjian/beer"
	"io/ioutil"
	"log"
)

func main() {
	srv := beer.New()

	srv.Static("/img/", "./public/image")
	srv.Static("/conf/", "./public/conf/")

	beer.Loadini("./public/conf/app.ini")

	srv.GET("/", Default)
	srv.POST("/detail/:id", Detail)
	srv.GET("/login", Login)
	srv.GET("/rem", Rem)
	if err := srv.Run(":8088"); err != nil {
		panic(err)
	}
}

func Rem(c *beer.Context)  {
	sess, _ := beer.Session().Start(c)
	beer.Session().Destroy(sess)
}

func Default(c *beer.Context) {

}

func Login(c *beer.Context) {
	log.Println(c.Param("id"))
	log.Println(c.Param("name"))
	log.Println(c.UserAgent)
}

func Detail(c *beer.Context) {
	b, err := ioutil.ReadAll(c.Body)
	if err != nil {
		log.Printf("err:%+v\n", err)
		return
	}
	log.Println(string(b))
}