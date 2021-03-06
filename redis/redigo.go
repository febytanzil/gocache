package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type redigo struct {
	pool redis.Pool
	c    *config
}

func (r *redigo) Set(key, value string, expire time.Duration) error {
	conn := r.pool.Get()
	defer conn.Close()

	var err error

	if expire.Seconds() > 0 {
		_, err = conn.Do("SET", key, value, "EX", int64(expire.Seconds()))
	} else {
		_, err = conn.Do("SET", key, value)
	}

	return err
}

func (r *redigo) SetNX(key, value string, expire time.Duration) (bool, error) {
	conn := r.pool.Get()
	defer conn.Close()

	exist := false
	reply, err := conn.Do("SET", key, value, "NX", "EX", int64(expire.Seconds()))
	if redis.ErrNil == err || nil == reply {
		exist = true
	}

	return exist, err
}

func (r *redigo) Get(key string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()

	reply, err := conn.Do("GET", key)
	return redis.String(reply, err)
}

func (r *redigo) Del(key string) error {
	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}
