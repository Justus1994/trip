package main

import (
	"log"

	"github.com/go-redis/redis"
)

//NewClient creates new redis client
func newClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Panic("can't establish connection to redis")

	}
	log.Print("connected to redis ")
	return client
}
