package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

var dbToken = "validtoken"

// Auth : return : value of Authorization header or a new token
func Auth(r *http.Request) []byte {

	if AuthToken(r) {
		return []byte(getToken(r))
	}

	newToken := newToken()
	return save(newToken)
}

//AuthToken : authoring a token return : true if authorized
func AuthToken(r *http.Request) bool {
	token := getToken(r)
	exist := redisClient.SIsMember(dbToken, string(token)).Val()

	if exist {
		return true
	}

	return false
}

//getToken : return token from http header
func getToken(r *http.Request) string {
	return r.Header.Get("Authorization")
}

func newToken() []byte {
	token := uuid.New()
	return []byte(token.String())
}

func save(token []byte) []byte {
	log.Print("save new Token")
	redisClient.SAdd(dbToken, token)
	return token
}
