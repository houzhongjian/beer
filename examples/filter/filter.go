package filter

import (
	"github.com/houzhongjian/beer"
	"log"
)

func FilterLogin(c *beer.Context) {
	isLogin := c.Param("is_login")
	if isLogin != "yes" {
		c.Json(map[string]interface{}{"code":1001,"msg":"未登录"})
		c.MiddlewareReturn()
		return
	}
	log.Println("FilterLogin")
}
