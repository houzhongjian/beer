package beer

import (
	"net/http"
	"sync"
)

type sessionManager struct {
	lock sync.RWMutex
	data map[string]interface{}
}

var session *sessionManager

func init() {
	session = &sessionManager{
		lock: sync.RWMutex{},
		data: make(map[string]interface{}),
	}
}

//Session 获取session管理器对象.
func Session() *sessionManager {
	return session
}

func (sess *sessionManager) Start(r *http.Request) {

}

func (sess *sessionManager) Get(key string) {

}

func (sess *sessionManager) Save() {

}
