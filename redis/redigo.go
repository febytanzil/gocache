package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type redigo struct {
	pool redis.Pool
	c    *config
}

func (r *redigo) Set(key, value string, expiry time.Duration) error {
	conn := r.pool.Get()
	defer conn.Close()

	var err error

	if expiry.Seconds() > 0 {
		_, err = conn.Do("SET", key, value, "EX", expiry.Seconds())
	} else {
		_, err = conn.Do("SET", key, value)
	}

	return err
}

func (r *redigo) SetNX(key, ttl time.Duration) (bool, error) {
	conn := r.pool.Get()
	defer conn.Close()

	exist := false
	_, err := conn.Do("SET", key, "1", "NX", "EX", ttl.Seconds())
	if redis.ErrNil == err {
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
