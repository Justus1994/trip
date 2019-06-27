package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

var dbToken = "validtoken"

// Auth : return : value of Authorization header or a new token
func Auth(r *http.Request) ([]byte, error) {
	token, err := AuthToken(r)
	if err != nil {
		return nil, err
	}
	if token {
		return []byte(getToken(r)), nil
	}

	newToken := newToken()
	return save(newToken), nil
}

//AuthToken : authoring a token return : true if authorized
func AuthToken(r *http.Request) (bool, error) {
	token := getToken(r)
	exist, err := redisClient.SIsMember(dbToken, string(token)).Result()
	if err != nil {
		log.Printf("error with redis occurred, error : %v", err)
		return false, err
	}

	return exist, nil

}

//getToken : return token from http header
func getToken(r *http.Request) string {
	return r.Header.Get("Authorization")
}

func newToken() []byte {
	log.Print("generate new token")
	return []byte(uuid.New().String())
}

func save(token []byte) []byte {
	log.Print("save new Token")
	redisClient.SAdd(dbToken, token)
	return token
}
