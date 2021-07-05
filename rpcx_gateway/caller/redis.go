package main

import (
	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
)

func NewRedisClient() *goredislib.Client {
	return goredislib.NewClient(&goredislib.Options{
		Addr: "localhost:6379",
	})
}

func NewRedisMutex(client *goredislib.Client) *redsync.Mutex {
	pool := goredis.NewPool(client)

	rs := redsync.New(pool)

	mutexname := "my-global-rmutex"
	return rs.NewMutex(mutexname)
}