package session

import (
	"fmt"
	"log"
)

var sessionKey = "session:"

func (r *RedisSession) Get(key string) string {
	key = fmt.Sprintf("%s%s:%s", sessionKey, r.SessionID, key)
	return redisCli.Get(key).Val()
}

func (r *RedisSession) Set(key string, val string) {
	key = fmt.Sprintf("%s%s:%s", sessionKey, r.SessionID, key)
	if err := redisCli.Set(key, val, 0).Err(); err != nil {
		log.Printf("err:%+v\n", err)
		return
	}
}
