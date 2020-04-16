<img src='./logo.jpeg'>

## Beer
Beer是一个使用golang开发的轻量级web框架，目的是希望使用这个框架就像夏天坐在路边喝着啤酒吃着烧烤一样的爽。由于参考了一部分的gin代码，所有很多方法与gin的类似。**目前还处于维护阶段不能用于生产环境**

## Examples

```go
package main

import (
	"beer"
	"fmt"
	"log"
)

func main() {
	srv := beer.New()

	srv.Static("/img/", "./public/image")
	srv.Static("/conf/", "./public/conf/")

	beer.Loadini("./public/conf/app.ini")

	srv.GET("/", Default)
	srv.GET("/detail", Detail)
	srv.GET("/login/:userid/:name/:age", Login)
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
```
