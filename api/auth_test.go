package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var token string

func TestAuthNewToken(t *testing.T) {
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
	token = recorder.Body.String()
	if token == expected {
		t.Errorf("handler returned unexpected body: got %v want %v", token, expected)
	}
}

func TestAuth(t *testing.T) {
	req, err := http.NewRequest("GET", "/auth", nil)
	req.Header.Set("Authorization", token)
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
	expected := token
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
