package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var tripHandlerTestToken []byte

func setupHandlerTest() {
	log.Println("Started TripHandlerTest")
	tripHandlerTestToken = newToken()
	save(tripHandlerTestToken)
}

func TestCreateTripHandler(t *testing.T) {
	setupHandlerTest()
	req, err := http.NewRequest("POST", "/trip/bodensee", nil)
	req.Header.Set("Authorization", string(tripHandlerTestToken))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	newRouter().ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

func TestGetAllTripsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/trip", nil)
	req.Header.Set("Authorization", string(tripHandlerTestToken))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	newRouter().ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetTripHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/trip/0", nil)
	req.Header.Set("Authorization", string(tripHandlerTestToken))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	newRouter().ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteNodeHandler(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/trip/0/nodes/0", nil)
	req.Header.Set("Authorization", string(tripHandlerTestToken))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	newRouter().ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestDeleteTripHandler(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/trip/0", nil)
	req.Header.Set("Authorization", string(tripHandlerTestToken))
	if err != nil {
		t.Fatal(err)
	}
	recorder := httptest.NewRecorder()

	newRouter().ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
