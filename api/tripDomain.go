package main

import "encoding/json"

//Trip : Trip holds a Array of nodes
type Trip struct {
	Nodes []node
}

//NewTrip : constructor
func NewTrip(nodes []byte) *Trip {
	trip := new(Trip)
	trip.Nodes = cleanNode(nodes)
	return trip
}

func cleanNode(data []byte) []node {
	var nodes []node
	json.Unmarshal(data, &nodes)
	return nodes
}

//Node : represent a point on the world
type node struct {
	ID            string        `json:"id,omitempty"`
	PictureSource pictureSource `json:"urls,omitempty"`
	Location      location      `json:"location,omitempty"`
}

//PictureSource : represent the picture link of a node
type pictureSource struct {
	Raw     string `json:"raw,omitempty"`
	Full    string `json:"full,omitempty"`
	Regular string `json:"regular,omitempty"`
}

//Location : represent the location of a node
type location struct {
	Title    string   `json:"title,omitempty"`
	City     string   `json:"city,omitempty"`
	Country  string   `json:"country,omitempty"`
	Position position `json:"position,omitempty"`
}

//Position : represent a point on the world
type position struct {
	Lat  float64 `json:"latitude,omitempty"`
	Long float64 `json:"longitude,omitempty"`
}
