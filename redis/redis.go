package main

import (
	"fmt"
	"log"

	"github.com/garyburd/redigo/redis"
)

const (
	key      = "KEY"
	val      = "VALUE"
	ip_port  = "127.0.0.1:6379"
	password = "foobared"
)

func main() {
	c := redisConnection()
	defer c.Close()

	redisSet(key, val, c)
	fmt.Println(redisGet(key, c))
}

func redisConnection() redis.Conn {

	c, err := redis.Dial("tcp", ip_port)
	if err != nil {
		log.Fatalf("redis.Dial got error: %T", err)
	}
	if _, err := c.Do("AUTH", password); err != nil {
		c.Close()
		log.Fatalf("c.DO(AUTH)  got error: %T", err)
	}
	return c
}

func redisSet(key string, value string, c redis.Conn) {
	c.Do("SET", key, value)
}

func redisGet(key string, c redis.Conn) string {
	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		log.Fatalf("redis.String got error: %f", err)
	}
	return s
}
