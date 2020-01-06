package redis

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

type Cache interface {
	Basic
}

type Basic interface {
	Set(key, value string, expiry time.Duration) error
	SetNX(key, ttl time.Duration) (bool, error)
	Get(key string) (string, error)
	Del(key string) error
}

func NewBasic(address string, opts ...Option) Basic {
	c := &config{}
	for _, o := range opts {
		o(c)
	}

	return &redigo{c: c, pool: redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", address)
			if nil != err {
				return nil, err
			}
			if c.auth {
				if _, err := conn.Do("AUTH", c.password); err != nil {
					conn.Close()
					return nil, err
				}
			}

			return conn, nil
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
