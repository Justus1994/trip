package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var authTestToken []byte

func setupAuthTest() {
	log.Println("Started AuthTest")
	authTestToken = newToken()
	save(authTestToken)
}
func TestAuthNewToken(t *testing.T) {
	setupAuthTest()
	req, err := http.NewRequest("GET", "/auth", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(auth)
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//asserts: a new Token!
	expected := `c34a993f-4c98-4bbd-ac13-8bba1e78a786`
	token := recorder.Body.Bytes()
	var actual string
	json.Unmarshal(token, &actual)
	if expected == actual {
		t.Errorf("handler returned unexpected body: got %v want %v", expected, actual)
	}
}

func TestAuth(t *testing.T) {
	req, err := http.NewRequest("GET", "/auth", nil)
	req.Header.Set("Authorization", string(authTestToken))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	hf := http.HandlerFunc(auth)

	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//asserts: Some Token
	expected := string(authTestToken)
	token := recorder.Body.Bytes()
	var actual string
	json.Unmarshal(token, &actual)
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", expected, actual)
	}
}

func TestAuthModifyToken(t *testing.T) {
	req, err := http.NewRequest("GET", "/trip", nil)
	req.Header.Set("Authorization", "ModifiedToken")
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	newRouter().ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}

}
