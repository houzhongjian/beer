package beer

import (
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
	sessi "github.com/houzhongjian/beer/session"
)

const SESSION_NAME = "BEER_SESSION"

type sessionManager struct {
	lock sync.RWMutex
	item map[string]*session
}

var sess *sessionManager

type session struct {
	id       string
	data     sessi.Session
	response http.ResponseWriter
	request  *http.Request
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

func (s *sessionManager) Options(opt *sessi.Options) {
	opt.InitRedisDrive()
	sessi.SessionsOptions = opt
}

func (*sessionManager) createSessionId(c *Context) *session {
	id := uuid.New().String()
	id = strings.Replace(id, "-", "", -1)
	http.SetCookie(c.Response, &http.Cookie{
		Name:     SESSION_NAME,
		Value:    id,
		Path:     "/",
		HttpOnly: true,
	})
	s := &session{
		id:       id,
		request:  c.Request,
		response: c.Response,
	}

	if sessi.SessionsOptions == nil {
		sessi.SessionsOptions = &sessi.Options{
			Drive:    sessi.MemoryDrive,
		}
	}

	if sessi.SessionsOptions.Drive == sessi.RedisDrive {
		s.data = new(sessi.RedisSession)
	} else {
		ram := new(sessi.RamSession)
		ram.Data = make(map[string]string)
		s.data = ram
	}
	sess.item[id] = s
	return s
}

func (sess *sessionManager) Start(c *Context) (s *session, err error) {
	sess.lock.Lock()
	defer sess.lock.Unlock()

	//判断是否有cookieid.
	cookie, err := c.Request.Cookie("BEER_SESSION")
	if err != nil && err != http.ErrNoCookie {
		log.Printf("err:%+v\n", err)
		return s, err
	}
	//没有则创建，有则跳过.
	if cookie == nil {
		return sess.createSessionId(c), nil
	}

	//判断当前的cookid是否有session对象.
	s, ok := sess.item[cookie.Value]
	if !ok {
		s = sess.createSessionId(c)
		return s, nil
	}

	return s, nil
}

func (s *session) Get(key string) string {
	return s.data.Get(key)
}

func (s *session) Set(key string, val string) {
	s.data.Set(key, val)
}

func (sm *sessionManager) Destroy(s *session) {
	sm.lock.Lock()
	defer sm.lock.Unlock()
	delete(sm.item, s.id)

	//删除cookie.
	http.SetCookie(s.response, &http.Cookie{
		Name:    SESSION_NAME,
		MaxAge:  -1,
		Expires: time.Now().Add(-100 * time.Hour), // Set expires for older versions of IE
		Path:    "/",
	})
}
