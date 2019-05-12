package main

import (
	"log"
	"net/http"
	"os/exec"
)

var dbToken = "validtoken"

// Auth : return : value of Authorization header or a new token
func Auth(r *http.Request) []byte {
	token := []byte(r.Header.Get("Authorization"))

	if AuthToken(token) {
		return token
	}

	newToken := newToken()
	return save(newToken)
}

//AuthToken : authoring a token return : true if authorized
func AuthToken(token []byte) bool {
	exist := redisClient.SIsMember(dbToken, string(token)).Val()

	if exist {
		return true
	}

	return false
}

func newToken() []byte {
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal("Can't create new Token", err)
	}
	//remove last char from string because exec.command returns with newLine
	return out[:len(out)-1]
}

func save(token []byte) []byte {
	log.P
	redisClient.SAdd(dbToken, token)
	redisClient.Set(string(token), 0, 0)
	return token
}
