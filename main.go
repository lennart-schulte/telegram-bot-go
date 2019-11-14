package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const (
	API_KEY          = "APIKEY"
)

func main() {
	// load API key from file
	content, err := ioutil.ReadFile(API_KEY)
	if err != nil {
		fmt.Println(err)
		return
	}
	apiKey := strings.TrimSpace(string(content))

	// access basic telegram bot information
	t := NewTelegramBot(apiKey)
	getMe, err := t.GetMe();
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(getMe.Result.Username)
}
