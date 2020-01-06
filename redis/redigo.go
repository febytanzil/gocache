package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

type redigo struct {
	pool redis.Pool
	c    *config
}

func (r *redigo) Set(key, value string, ttl time.Duration) error {
	conn := r.pool.Get()
	defer conn.Close()
}

func (r *redigo) SetNX(key, ttl time.Duration) error {
	conn := r.pool.Get()
	defer conn.Close()
}

func (r *redigo) Get(key string) (string, error) {
	conn := r.pool.Get()
	defer conn.Close()
}

func (r *redigo) Del(key string) error {
	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}
