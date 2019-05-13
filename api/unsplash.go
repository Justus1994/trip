package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

//Unsplash : object to help consume the Unsplash API
type Unsplash struct {
	host        string
	endpoint    string
	paramKeys   []string
	paramValues []string
}

//NewUnsplash : constructor for Unsplash
func NewUnsplash(endpoint string) *Unsplash {
	unsplash := new(Unsplash)
	unsplash.host = "http://api.unsplash.com/"
	unsplash.endpoint = endpoint
	return unsplash
}

func (unsplash *Unsplash) buildURL() string {
	return unsplash.host + unsplash.endpoint
}

func (unsplash *Unsplash) addParam(key string, value string) {
	unsplash.paramKeys = append(unsplash.paramKeys, key)
	unsplash.paramValues = append(unsplash.paramValues, value)

}

//Send : sends a new http Request to the endpoint stored in unsplash struct
func (unsplash *Unsplash) Send() []byte {
	req, _ := http.NewRequest("GET", unsplash.buildURL(), nil)

	q := req.URL.Query()
	for i := range unsplash.paramKeys {
		q.Add(unsplash.paramKeys[i], unsplash.paramValues[i])
	}
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	return data

}
