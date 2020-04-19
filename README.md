<img src='./logo.jpeg'>

## Beer
Beer是一个使用golang开发的轻量级web框架，目的是希望使用这个框架就像夏天坐在路边喝着啤酒吃着烧烤一样的爽。由于参考了一部分的gin代码，所有很多方法与gin的类似。**目前还处于维护阶段不能用于生产环境**

## Examples

```go
package main

import (
	"github.com/houzhongjian/beer"
	"log"
)

func main() {
    //加载配置文件.
	beer.Loadini("./public/conf/app.ini")
	r := beer.New()

	r.Static("/img/", "./public/image")
	r.Static("/conf/", "./public/conf/")
	r.GET("/", Default)
	r.GET("/detail/:id", Detail)
	r.POST("/login", Login)

	if err := r.Run(":8088"); err != nil {
		panic(err)
	}
}

func Default(c *beer.Context) {
	log.Println(c.IP)
	c.Layout = "blog/layout.html"
	c.Data["name"] = "张三"
	c.Data["age"] = 20
	c.Data["title"] = "goBeer"
	c.Html("blog/index.html")
}

func Login(c *beer.Context) {
    session,err := beer.Session().Start(c)
	if err != nil {
		log.Printf("err:%+v\n",err)
		return
	}
	session.Set("name","张三")

	c.Data["code"] = 1000
	c.Data["msg"] = "登录成功"
	c.Json()
}

func Detail(c *beer.Context) {
	log.Println(c.Param("id"))
	log.Println(c.UserAgent)

	c.Layout = "blog/layout.html"
	c.Data["name"] = "zhangsan"
	c.Data["age"] = 20
	c.Data["title"] = "博客标题"
	c.Html("blog/detail.html")
}
```
