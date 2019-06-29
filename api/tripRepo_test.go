package main

import (
	"log"
	"math/rand"
	"strings"
	"testing"
)

var tripRepoTestToken []byte

func setupRepoTest() {
	log.Println("Started TripRepoTest")
	tripRepoTestToken = newToken()
	save(tripRepoTestToken)
}
func TestCreateTrip(t *testing.T) {
	setupRepoTest()
	tripTest, _ := createTrip(string(tripRepoTestToken), "indonesia")

	expectedNot := ""
	count := 0

	for i := range tripTest.Nodes {
		corrupted := false
		title := tripTest.Nodes[i].Location.Title
		country := tripTest.Nodes[i].Location.Country
		city := tripTest.Nodes[i].Location.City

		if strings.Contains(title, country) {
			if country != "" {
				t.Errorf("createTrip returned unexpected trip: title has a country value in Index %v", i)
			}
		}
		if title == expectedNot && country == expectedNot && city == expectedNot {
			t.Errorf("createTrip returned unexpected trip: got corrupted node in Index %v", i)
			corrupted = true
		}

		if corrupted == true {
			count = count + 1
		}

	}
	log.Printf("got %v corrupted objects from %v objects", count, len(tripTest.Nodes))
}

func TestCreateTripInvalidTags(t *testing.T) {
	_, err := createTrip(string(tripRepoTestToken), "definitelyNoValidTag")
	if err == nil {
		t.Error("createTrip should return a error")
	}

}

func TestGetTrip(t *testing.T) {
	expected, _ := createTrip(string(tripRepoTestToken), "indonesia")
	actual := getTrip(string(tripRepoTestToken), 0)
	for i := range expected.Nodes {
		if expected.Nodes[i] != actual.Nodes[i] {
			t.Errorf("TestGetTrip return unexpected trip got invalid node %v want %v", expected.Nodes[i], actual.Nodes[i])
		}
	}
}

func TestGetAllTrips(t *testing.T) {
	expected := 2
	allTrips := getAllTrips(string(tripRepoTestToken))
	actual := len(allTrips)
	if expected != actual {
		t.Errorf("TestGetAllTrips return unexpected len of trips got %v want %v", expected, actual)
	}

}

func TestDeleteTrip(t *testing.T) {
	deleteTrip(string(tripRepoTestToken), 0)
	allTrips := getAllTrips(string(tripRepoTestToken))
	expected := 1
	actual := len(allTrips)

	if expected != actual {
		t.Errorf("TestDeleteTrips return unexpected len of trips got %v want %v", expected, actual)
	}
}

func TestDeleteNode(t *testing.T) {
	allTrips := getAllTrips(string(tripRepoTestToken))
	expectedNot := allTrips[0].Nodes[rand.Intn(len(allTrips[0].Nodes))]
	actualTrip := deleteNode(string(tripRepoTestToken), 0, expectedNot.ID)

	for i := range actualTrip.Nodes {
		if expectedNot == actualTrip.Nodes[i] {
			t.Errorf("TestDeleteNode return unexpected Trips identical Node in index %v", i)
		}
	}
}

func benchmarkCleanNodes(tag string, b *testing.B) {
	newtrip := NewTrip(getNodes(tag))
	var cleantrip Trip
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		cleanNodes(newtrip.Nodes, &cleantrip)
	}
}

func BenchmarkCleanNodesStuttgart(b *testing.B) { benchmarkCleanNodes("stuttgart", b) }
func BenchmarkCleanNodesBodensee(b *testing.B)  { benchmarkCleanNodes("bodensee", b) }
func BenchmarkCleanNodesIndonesia(b *testing.B) { benchmarkCleanNodes("indonesia", b) }
func BenchmarkCleanNodesMalaysia(b *testing.B)  { benchmarkCleanNodes("malaysia", b) }
func BenchmarkCleanNodesLondon(b *testing.B)    { benchmarkCleanNodes("london", b) }
func BenchmarkCleanNodesParis(b *testing.B)     { benchmarkCleanNodes("Paris", b) }
