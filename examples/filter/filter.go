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
	log.Println("FilterLogin")
}
