package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var tripTestToken []byte

func setup() {
	tripTestToken = newToken()
	save(tripTestToken)
}

func TestTripsData(t *testing.T) {
	setup()
	req, err := http.NewRequest("POST", "/trip/indonesia", nil)
	req.Header.Set("Authorization", string(tripTestToken))
	if err != nil {
		t.Fatal(err)
	}
	log.Print(req.URL)
	recorder := httptest.NewRecorder()

	newRouter().ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	data := recorder.Body.String()

	var tripTest Trip
	json.Unmarshal([]byte(data), &tripTest)

	expectedNot := ""
	count := 0

	for i := range tripTest.Nodes {
		corrupted := false
		title := tripTest.Nodes[i].Location.Title
		country := tripTest.Nodes[i].Location.Country
		city := tripTest.Nodes[i].Location.City

		if strings.Contains(title, country) {
			t.Errorf("handler returned unexpected body: title has a country value in Index %v", i)
		}
		if title == expectedNot && country == expectedNot && city == expectedNot {
			t.Errorf("handler returned unexpected body: got corrupted object in Index %v", i)
			corrupted = true
		}

		if corrupted == true {
			count = count + 1
		}

	}
	log.Printf("got %v corrupted objects from %v objects", count, len(tripTest.Nodes))
}
