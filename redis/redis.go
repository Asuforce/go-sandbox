package main

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

const (
	key      = "KEY"
	val      = "VALUE"
	ipPort   = "127.0.0.1:6379"
	password = "foobared"
)

func main() {
	c := redisConnection()
	defer c.Close()

	redisSet(key, val, c)
	fmt.Println(redisGet(key, c))
}

func redisConnection() redis.Conn {
	c, err := redis.Dial("tcp", ipPort)
	if err != nil {
		log.Fatalf("redis.Dial got error: %v", err)
	}
	if _, err := c.Do("AUTH", password); err != nil {
		c.Close()
		log.Fatalf("c.DO(AUTH)  got error: %v", err)
	}
	return c
}

func redisSet(key string, value string, c redis.Conn) {
	c.Do("SET", key, value)
}

func redisGet(key string, c redis.Conn) string {
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		log.Fatalf("redis.String got error: %v", err)
	}
	return s
}
