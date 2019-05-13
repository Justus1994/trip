package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var redisClient = newClient()

func newRouter() *mux.Router {

	/***** All Endpoints are available under localhost:8080/api if deployed via docker compose *****/
	r := mux.NewRouter()

	/***** Handle Trip ******/
	r.HandleFunc("/trip", GetAllTripsHandler).Methods("GET")
	r.HandleFunc("/trip/{tag}", CreateTripHandler).Methods("POST")
	r.HandleFunc("/trip/{id}", GetTripHandler).Methods("GET")
	r.HandleFunc("/trip/{id}/nodes", CreateNodesHandler).Methods("POST")
	r.HandleFunc("/trip/{tripid}/nodes/{nodeid}", DeleteNodeHandler).Methods("DELETE")

	/***** Handle authentication *****/
	r.HandleFunc("/auth", auth).Methods("GET")

	return r
}

func auth(w http.ResponseWriter, r *http.Request) {
	w.Write(Auth(r))
}

func main() {
	router := newRouter()
	log.Println("server listening on port 8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", router))

}
