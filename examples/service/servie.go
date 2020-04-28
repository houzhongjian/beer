package service

import (
	"fmt"
	"github.com/houzhongjian/beer"
	"github.com/houzhongjian/beer/log"
)

func Rem(c *beer.Context)  {
	sess, _ := beer.Session().Start(c.Response, c.Request)
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
		account := c.Param("account")
		password := c.Param("password")
		if account == "zhangsan" && password == "123456" {
			session, err := beer.Session().Start(c.Response, c.Request)
			if err != nil {
				fmt.Println(err)
				c.Data["code"] = 1001
				c.Data["msg"] = "登录失败"
				c.Json()
				return
			}

			session.Set("name","张三")
			session.Set("age","20")
			session.Set("is_login","yes")

			c.Data["code"] = 1000
			c.Data["msg"] = "登录成功"
			c.Json()
			return
		}
		user := User{}
		if err := c.BindJSON(&user); err != nil {
			log.Debug(err)
			return
		}
		log.Debug(user)
		return
	}

	c.Html("admin/login.html")
}

func Detail(c *beer.Context) {
	c.Layout = "blog/layout.html"
	sess,err := beer.Session().Start(c.Response, c.Request)
	if err != nil {
		log.Println(err.Error())
		return
	}

	age, err := sess.Get("age").Int()
	if err != nil {
		log.Debug(err)
		return
	}
	fmt.Println(age)
	c.Data["title"] = sess.Get("name").String()
	c.Html("blog/detail.html")
}
