package redislock

import (
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	goredislib "github.com/redis/go-redis/v9"
)

var (
	RS *redsync.Redsync
)

func Init() {
	RS = newRs()
}

func newRs() *redsync.Redsync {
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)

	rs := redsync.New(pool)
	return rs
}
