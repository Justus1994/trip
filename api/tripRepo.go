package main

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
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

func createTrip(token string, tag string) (Trip, error) {
	newtrip := NewTrip(getNodes(tag))
	var cleantrip Trip

	cleanNodes(newtrip.Nodes, &cleantrip)

	json, _ := json.Marshal(cleantrip)

	if len(cleantrip.Nodes) == 0 {
		return cleantrip, errors.New("can't find any pictures for tag " + tag)
	}
	redisClient.LPush((token + ":trips"), json)

	return cleantrip, nil
}

func getTrip(token string, id int64) Trip {
	var trip Trip
	data := redisClient.LRange(token+":trips", id, id).Val()
	json.Unmarshal([]byte(data[0]), &trip)
	return trip
}

func deleteNode(token string, tripID int64, nodeID string) Trip {
	var trip Trip
	data := redisClient.LRange(token+":trips", tripID, tripID).Val()
	json.Unmarshal([]byte(data[0]), &trip)
	newNodeID := findNode(trip.Nodes, nodeID)
	trip.Nodes = append(trip.Nodes[:newNodeID], trip.Nodes[newNodeID+1:]...)
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
	unsplash := NewUnsplash("photos/random/")
	unsplash.addParam("query", tag)
	unsplash.addParam("count", "30")
	unsplash.addParam("featured", "true")
	unsplash.addParam("client_id", os.Getenv("UNSPLASH_API_KEY"))

	return unsplash.Send()
}

func cleanNodes(nodes []node, trip *Trip) {
	var cleanedNodes []node
	for _, element := range nodes {
		if element.Location.Country == "" && element.Location.City == "" && element.Location.Title == "" {
		} else {
			cleanedNodes = append(cleanedNodes, element)
		}
	}

	for i, element := range cleanedNodes {
		cleanedNodes[i].Location.Title = strings.Replace(element.Location.Title, element.Location.Country, "", -1)
	}
	trip.Nodes = cleanedNodes

}

func findNode(nodes []node, id string) int {
	for i, node := range nodes {
		if node.ID == id {
			return i
		}
	}
	return 0
}
