package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
	"strings"
)

const (
	API_KEY          = "APIKEY"
	TELEGRAM_API_URL = "https://api.telegram.org/"
)

type telegramBot struct {
	baseUrl string
}

func NewTelegramBot(apiKey string) *telegramBot {
	baseUrl := TELEGRAM_API_URL + "bot" + apiKey
	t := telegramBot{
		baseUrl: baseUrl,
	}
	return &t
}

func (t *telegramBot) getMe() ([]byte, error) {
	url := t.baseUrl + "/getMe"
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
	getMe, err := t.getMe()
	if err != nil {
		fmt.Println(err)
		return
	}

	// check if the request was ok
	var getMeInterface interface{}
	err = json.Unmarshal(getMe, &getMeInterface)
	if err != nil {
		fmt.Println(err)
		return
	}
	// this is a bit hacky, maybe at some point I define the interface
	if getMeInterface.(map[string]interface{})["ok"].(bool) {
		fmt.Println("Request was ok", string(getMe))
	} else {
		fmt.Println("Request was NOT okay", string(getMe))
	}

}
