package session

import (
	"github.com/go-redis/redis/v7"
)

//Drive session的驱动
type Drive int

//redisCli .
var redisCli *redis.Client

const (
	MemoryDrive Drive = iota
	RedisDrive
)

//SessionsOptions .
var SessionsOptions *Options

type Options struct {
	Drive    Drive
	Addr     string
	DB       int
	Password string
	Expired  int
}

type SessionStore interface {
	Get(key string) string
	Set(key string, val string)
}

type RedisSession struct {
	SessionID string
}

type RamSession struct {
	Data map[string]string
}

func (opt *Options) InitRedisDrive() {
	client := redis.NewClient(&redis.Options{
		Addr:     opt.Addr,
		Password: opt.Password,
		DB:       opt.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	redisCli = client
}

func (s *Session) Get(key string) SessionResult {
	result := SessionResult{
		val: s.data.Get(key),
	}
	return result
}

func (s *Session) Set(key string, val string) {
	s.data.Set(key, val)
}
