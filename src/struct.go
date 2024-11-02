package main

import "github.com/garyburd/redigo/redis"

type Event struct {
	Type string
	Title string
	Description string
	Addtime int64
	Tiptime int64
	Owner string // used as the redis key prefix
}

type RedisConfigs struct {
	address string
	network string
	options redis.DialOption
}

type ServerConfigs struct {
	addr string
	static string
}
