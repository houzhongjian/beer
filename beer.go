package beer

import (
	"github.com/houzhongjian/beer/conf"
	"github.com/houzhongjian/beer/log"
	"github.com/houzhongjian/beer/session"
)

//Session 获取session管理器对象.
func Session() *session.Manager {
	return session.Object
}

//Config 获取config管理器对象..
func Config() *conf.ConfigManager {
	return conf.Manager
}

func Log() *log.LogManager  {
	return log.Manager
}