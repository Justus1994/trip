package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAllTripsHandler :
func GetAllTripsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("User requested to /trip, GET")
	if authRequest(w, r) {
		data := getAllTrips(getToken(r))
		json.NewEncoder(w).Encode(data)
	}

}

//CreateTripHandler :
func CreateTripHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("User requested a new trip with tag : %v to /trip, POST", mux.Vars(r)["tag"])
	if authRequest(w, r) {
		params := mux.Vars(r)
		if trip, err := createTrip(getToken(r), params["tag"]); err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(err.Error())
		} else {
			json.NewEncoder(w).Encode(trip)
		}
	}
}

//GetTripHandler :
func GetTripHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("User requested to /trip/%v, GET", mux.Vars(r)["id"])
	if authRequest(w, r) {
		params := mux.Vars(r)
		id, _ := strconv.ParseInt(params["id"], 10, 64)
		data := getTrip(getToken(r), id)
		json.NewEncoder(w).Encode(data)

	}
}

//DeleteTripHandler :
func DeleteTripHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("User requested to /trip/%v, DELETE", mux.Vars(r)["id"])
	if authRequest(w, r) {
		params := mux.Vars(r)
		ID, _ := strconv.ParseInt(params["id"], 10, 64)
		data := deleteTrip(getToken(r), ID)
		json.NewEncoder(w).Encode(data)
	}
}

//DeleteNodeHandler :
func DeleteNodeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("User requested to /trip/%v/%v, DELETE", mux.Vars(r)["tripid"], mux.Vars(r)["nodeid"])
	if authRequest(w, r) {
		params := mux.Vars(r)
		tripID, _ := strconv.ParseInt(params["tripid"], 10, 64)
		nodeID, _ := params["nodeid"]
		data := deleteNode(getToken(r), tripID, nodeID)
		json.NewEncoder(w).Encode(data)
	}
}

func authRequest(w http.ResponseWriter, r *http.Request) bool {
	if token, err := AuthToken(r); err != nil {
		return notFoundResponse(w)
	} else if !token {
		return unauthorizedResponse(w)
	}

	return true
}

func notFoundResponse(w http.ResponseWriter) bool {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(nil)
	return false
}

func unauthorizedResponse(w http.ResponseWriter) bool {
	log.Println("User unauthorized")
	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode("invalid token")
	return false
}
