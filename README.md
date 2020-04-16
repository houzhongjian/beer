<img src='./logo.jpeg'>

## Beer
Beer是一个使用golang开发的轻量级web框架，目的是希望使用这个框架的开发者就像夏天坐在路边喝着啤酒吃着烧烤一样的爽。由于参考了一部分的gin代码，所有很多方法与gin的类似。**目前还处于维护阶段不能用于生产环境**

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
```
