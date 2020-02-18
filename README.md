# gocache
wrapper for all (to-be) kinds of imdb memory cache

## Supported providers
- Redis

## Install
```bash
$ go get github.com/febytanzil/gocache
```

## Usage
```go
package main

import (
    "github.com/febytanzil/gocache/redis"
    "time"
)

func main() {
    basicRedis := redis.NewBasic("127.0.0.1:6379", redis.WithConnection(30, 10, 60*time.Second))
    
    err := basicRedis.Set("key", "value", 60*time.Second)
    if nil != err {
    	// handle the error
    }
}
```