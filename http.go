package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
)


// performs HTTP GET and returns the body as byte array
func get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// performs HTTP POST, payload is marshalled json, returns response as byte array
func post(url string, payload []byte) ([]byte, error) {
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
