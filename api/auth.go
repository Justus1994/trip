package main

import (
	"log"
	"net/http"
	"os/exec"
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
	out, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal("Can't create new Token", err)
	}
	//remove last char from string because exec.command returns with newLine
	return out[:len(out)-1]
}

func save(token []byte) []byte {
	log.Print("save new Token")
	redisClient.SAdd(dbToken, token)
	return token
}
