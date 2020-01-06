package redis

import "time"

type Option func(c *config)

type config struct {
	idleTimeout time.Duration
	maxActive   int
	maxIdle     int
	wait        bool
	password    string
	auth        bool
}

func WithConnection(maxActive, maxIdle int, idleTimeout time.Duration) Option {
	return func(c *config) {
		c.idleTimeout = idleTimeout
		c.maxActive = maxActive
		c.maxIdle = maxIdle
	}
}

func WaitPool(wait bool) Option {
	return func(c *config) {
		c.wait = wait
	}
}

func WithAuth(password string) Option {
	return func(c *config) {
		c.password = password
		c.auth = true
	}
}
