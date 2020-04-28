package filter

import (
	"github.com/houzhongjian/beer"
	"github.com/houzhongjian/beer/log"
)

func FilterLogin(c *beer.Context) {
	session, err := beer.Session().Start(c.Response, c.Request)
	if err != nil {
		c.Data["code"] = 1001
		c.Data["msg"] = "服务器错误"
		c.Json()
		c.MiddlewareReturn()
		return
	}
	if session.Get("is_login").String() != "yes" {
		c.Data["code"] = 1001
		c.Data["msg"] = "未登录"
		c.Json()
		c.MiddlewareReturn()
		return
	}

	log.Println("继续往下走")
	c.Data["userid"] = 34567
	c.Data["name"] = "张三"
	c.Data["age"] = 20
}
