package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type redigo struct {
	pool *redis.Pool
	c    *config
}

func (r *redigo) Set(key, value string, ttl time.Duration) (string, error) {

}

func (r *redigo) SetNX(key, ttl time.Duration) (string, error) {

}

func (r *redigo) Get(key string) (string, error) {

}
