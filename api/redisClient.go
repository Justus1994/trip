package main

import (
	"log"
	"os"

	"github.com/go-redis/redis"
)

//NewClient creates new redis client
func NewClient() *redis.Client {
	_, env := os.LookupEnv("HUNGRYHONEYDOCKER")
	var client *redis.Client
	if !env {
		client = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	} else {
		client = redis.NewClient(&redis.Options{
			Addr:     "redis:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}

	_, err := client.Ping().Result()
	if err != nil {
		log.Println("can't establish conntection to redis", err)
	} else {
		log.Println("connected to redis")
	}

	return client
}
