package main

import (
	"encoding/json"
	"log"
)

func getAllTrips(token string) []Trip {
	var trips []Trip
	length := redisClient.LLen(token + ":trips").Val()
	data := redisClient.LRange(token+":trips", 0, length).Val()

	for i := range data {
		var trip Trip
		json.Unmarshal([]byte(data[i]), &trip)
		trips = append(trips, trip)
	}

	return trips
}

func createTrip(token string, tag string) Trip {
	trip := NewTrip(getNodes(tag))
	json, err := json.Marshal(trip)

	if err != nil {
		log.Print(err)
	}

	redisClient.LPush((token + ":trips"), json)

	return *trip
}

func getTrip(token string, id int64) Trip {
	var trip Trip
	data := redisClient.LRange(token+":trips", id, id).Val()
	json.Unmarshal([]byte(data[0]), &trip)
	return trip
}

func deleteNode(token string, tripID int64, nodeID int64) Trip {
	var trip Trip
	data := redisClient.LRange(token+":trips", tripID, tripID).Val()
	json.Unmarshal([]byte(data[0]), &trip)
	trip.Nodes = append(trip.Nodes[:nodeID], trip.Nodes[nodeID+1:]...)
	json, _ := json.Marshal(trip)
	redisClient.LSet(token+":trips", tripID, json)
	return trip
}

func deleteTrip(token string, ID int64) []Trip {
	trips := getAllTrips(token)
	json, _ := json.Marshal(trips[ID])
	redisClient.LRem(token+":trips", 1, json)
	return getAllTrips(token)
}

func getNodes(tag string) []byte {
	uAPI := NewUnsplashAPI("photos/random/")
	return uAPI.getRandPics("1", tag)
}
