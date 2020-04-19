package filter

import (
	"github.com/houzhongjian/beer"
	"log"
)

func FilterLogin(c *beer.Context) {
	log.Println("执行中间件:FilterLogin")
	isLogin := c.Param("is_login")
	if isLogin != "yes" {
		c.Data["code"] = 1000
		c.Data["msg"] = "未登录"
		c.Json()
		c.MiddlewareReturn()
		return
	}
	c.Data["userid"] = 34567
	c.Data["name"] = "张三"
	c.Data["age"] = 20
	log.Println("FilterLogin")
}
