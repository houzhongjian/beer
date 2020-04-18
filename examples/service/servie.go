package service

import (
	"github.com/houzhongjian/beer"
	"io/ioutil"
	"log"
)

func Rem(c *beer.Context)  {
	sess, _ := beer.Session().Start(c)
	beer.Session().Destroy(sess)
}

func Default(c *beer.Context) {
	c.Layout = "blog/layer.html"
	data := map[string]interface{}{
		"name":"张三",
		"age":20,
		"title":"goBeer",
	}
	c.Html("blog/index.html", data)
}

func Login(c *beer.Context) {
	log.Println(c.Param("id"))
	log.Println(c.Param("name"))
	log.Println(c.UserAgent)

	c.Html("./views/admin/login.html",nil)
}

func Detail(c *beer.Context) {
	b, err := ioutil.ReadAll(c.Body)
	if err != nil {
		log.Printf("err:%+v\n", err)
		return
	}
	log.Println(string(b))
}
