package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAllTripsHandler :
func GetAllTripsHandler(w http.ResponseWriter, r *http.Request) {
	if authRequest(w, r) {
		data := getAllTrips(getToken(r))
		json.NewEncoder(w).Encode(data)
	}

}

//CreateTripHandler :
func CreateTripHandler(w http.ResponseWriter, r *http.Request) {
	if authRequest(w, r) {
		params := mux.Vars(r)
		trip := createTrip(getToken(r), params["tag"])
		json.NewEncoder(w).Encode(trip)
	}
}

//GetTripHandler :
func GetTripHandler(w http.ResponseWriter, r *http.Request) {
	if authRequest(w, r) {
		params := mux.Vars(r)
		id, _ := strconv.ParseInt(params["id"], 10, 64)
		data := getTrip(getToken(r), id)
		json.NewEncoder(w).Encode(data)

	}
}

//CreateNodesHandler :
func CreateNodesHandler(w http.ResponseWriter, r *http.Request) {
	authRequest(w, r)
}

//DeleteTripHandler :
func DeleteTripHandler(w http.ResponseWriter, r *http.Request) {
	if authRequest(w, r) {
		params := mux.Vars(r)
		ID, _ := strconv.ParseInt(params["id"], 10, 64)
		data := deleteTrip(getToken(r), ID)
		json.NewEncoder(w).Encode(data)
	}
}

//DeleteNodeHandler :
func DeleteNodeHandler(w http.ResponseWriter, r *http.Request) {
	if authRequest(w, r) {
		params := mux.Vars(r)
		tripID, _ := strconv.ParseInt(params["tripid"], 10, 64)
		nodeID, _ := params["nodeid"]
		data := deleteNode(getToken(r), tripID, nodeID)
		json.NewEncoder(w).Encode(data)
	}
}

func authRequest(w http.ResponseWriter, r *http.Request) bool {
	if !AuthToken(r) {
		w.Write([]byte("invalid token"))
		return false
	}
	return true
}
