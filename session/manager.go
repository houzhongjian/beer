package session

import (
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"
)

const SESSION_NAME = "BEER_SESSION"

var Object *Manager

type Manager struct {
	lock sync.RWMutex
	item map[string]*Session
}

type Session struct {
	id       string
	data     SessionStore
	response http.ResponseWriter
	request  *http.Request
}

func init() {
	sess := &Manager{
		lock: sync.RWMutex{},
		item: make(map[string]*Session),
	}

	Object = sess
}

func (m *Manager) Options(opt *Options) {
	opt.InitRedisDrive()
	SessionsOptions = opt
}

func (m *Manager) createSessionId(w http.ResponseWriter, r *http.Request) *Session {
	id := uuid.New().String()
	id = strings.Replace(id, "-", "", -1)
	http.SetCookie(w, &http.Cookie{
		Name:     SESSION_NAME,
		Value:    id,
		Path:     "/",
		HttpOnly: true,
	})
	s := &Session{
		id:       id,
		response: w,
		request:  r,
	}

	if SessionsOptions == nil {
		SessionsOptions = &Options{
			Drive: MemoryDrive,
		}
	}

	if SessionsOptions.Drive == RedisDrive {
		rs := new(RedisSession)
		rs.SessionID = id
		s.data = rs
	} else {
		ram := new(RamSession)
		ram.Data = make(map[string]string)
		s.data = ram
	}
	m.item[id] = s
	return s
}

func (m *Manager) Start(w http.ResponseWriter, r *http.Request) (s *Session, err error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	//判断是否有cookieid.
	cookie, err := r.Cookie("BEER_SESSION")
	if err != nil && err != http.ErrNoCookie {
		log.Printf("err:%+v\n", err)
		return s, err
	}
	//没有则创建，有则跳过.
	if cookie == nil {
		return m.createSessionId(w,r), nil
	}

	//判断当前的cookid是否有session对象.
	s, ok := m.item[cookie.Value]
	if !ok {
		s = m.createSessionId(w,r)
		return s, nil
	}

	return s, nil
}

func (m *Manager) Destroy(s *Session) {
	m.lock.Lock()
	defer m.lock.Unlock()
	delete(m.item, s.id)

	//删除cookie.
	http.SetCookie(s.response, &http.Cookie{
		Name:    SESSION_NAME,
		MaxAge:  -1,
		Expires: time.Now().Add(-100 * time.Hour), // Set expires for older versions of IE
		Path:    "/",
	})
}
