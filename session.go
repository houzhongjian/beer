package beer

import (
	"log"
	"net/http"
	"strings"
	"sync"
	"github.com/google/uuid"
)

type sessionManager struct {
	lock sync.RWMutex
	item map[string]*session
}

var sess *sessionManager

type session struct {
	id string
	data map[string]interface{}
}

func init() {
	sess = &sessionManager{
		lock: sync.RWMutex{},
		item: make(map[string]*session),
	}
}

//Session 获取session管理器对象.
func Session() *sessionManager {
	return sess
}

func (*sessionManager) createSessionId(c *Context) *session {
	id := uuid.New().String()
	id = strings.Replace(id,"-","",-1)
	http.SetCookie(c.Response, &http.Cookie{
		Name:       "BEER_SESSION",
		Value:      id,
		Path:       "/",
		HttpOnly:   true,
	})
	s := &session{
		id:   id,
		data: make(map[string]interface{}),
	}
	sess.item[id] = s
	return s
}

func (sess *sessionManager) Start(c *Context) (s *session, err error) {
	sess.lock.Lock()
	defer sess.lock.Unlock()

	//判断是否有cookieid.
	cookie, err := c.Request.Cookie("BEER_SESSION")
	if err != nil && err != http.ErrNoCookie  {
		log.Printf("err:%+v\n",err)
		return s, err
	}
	//没有则创建，有则跳过.
	if cookie == nil {
		return sess.createSessionId(c),nil
	}

	//判断当前的cookid是否有session对象.
	s, ok := sess.item[cookie.Value]
	if !ok {
		s = sess.createSessionId(c)
		return s,nil
	}

	return s,nil
}

func (s *session) Get(key string) interface{} {
	v, ok := s.data[key]
	if !ok {
		return nil
	}
	return v
}

func (s *session) Set(key string, val interface{}) {
	s.data[key] = val
}

func (sm *sessionManager) Destroy(sessionId string) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	delete(sm.item, sessionId)
}
