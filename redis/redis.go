package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type Cache interface {
	Basic
}

type Basic interface {
	Set(key, value string, ttl time.Duration) (string, error)
	SetNX(key, ttl time.Duration) (string, error)
	Get(key string) (string, error)
}

func NewBasic(address string, opts ...Option) Basic {
	c := &config{}
	for _, o := range opts {
		o(c)
	}

	return &redigo{c: c, pool: &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address)
			if nil != err {
				return nil, err
			}

			return c, nil
		},
		TestOnBorrow:    nil,
		MaxIdle:         c.maxIdle,
		MaxActive:       c.maxActive,
		IdleTimeout:     c.idleTimeout,
		Wait:            c.wait,
		MaxConnLifetime: 0,
	}}
}

func NewCache(address string, opts ...Option) Cache {
	return NewBasic(address, opts...)
}
