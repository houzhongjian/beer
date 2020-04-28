package service

import (
	"github.com/houzhongjian/beer"
	"github.com/houzhongjian/beer/log"
)

func Rem(c *beer.Context)  {
	sess, _ := beer.Session().Start(c)
	beer.Session().Destroy(sess)
}

func Default(c *beer.Context) {
	log.Println(c.IP)
	c.Layout = "blog/layout.html"
	c.Data["title"] = "goBeer"
	c.Html("blog/index.html")
}

type User struct {
	Account string `json:"account"`
	Password string	`json:"password"`
}
func Login(c *beer.Context) {
	if c.Method == "POST" {
		user := User{}
		if err := c.BindJSON(&user); err != nil {
			log.Debug(err)
			return
		}
		log.Debug(user)
		return
	}
	sess, err := beer.Session().Start(c)
	if err != nil {
		log.Debug(err)
		c.Data["code"] = 1001
		c.Data["msg"] = "登录失败"
		c.Json()
		return
	}
	sess.Set("name","张三")
	//log.Debug(sess.Get("name"))
	c.Data["code"] = 1000
	c.Data["msg"] = "登录成功"
	c.Json()
}

func Detail(c *beer.Context) {
	log.Println(c.Param("id"))
	log.Println(c.Param("name"))
	log.Println(c.UserAgent)

	c.Layout = "blog/layout.html"
	sess,err := beer.Session().Start(c)
	if err != nil {
		log.Println(err.Error())
		return
	}
	c.Data["title"] = sess.Get("name")
	c.Html("blog/detail.html")
}
