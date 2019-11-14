package main

import (
	"fmt"
	"encoding/json"
)

/*{"ok":true,"result":[{"update_id":1234,"message":{"message_id":1,"from":{"id":123,"is_bot":false,"first_name":"User","language_code":"en"},"chat":{"id":12,"first_name":"User","type":"private"},"date":1573736991,"text":"test"}}]}*/
type Update struct {
	UpdateId int `json:"update_id"`
	Message struct {
		MessageId int `json:"message_id"`
		From struct {
			Id int `json:"id"`
			IsBot bool `json:"is_bot"`
			FirstName string `json:"first_name"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			Id int `json:"id"`
			FirstName string `json:"first_name"`
			Type string `json:"type"`
		} `json:"chat"`
		Date int `json:"date"`
		Text string `json:"text"`
	}
}

type Updates struct {
	Ok bool `json:"ok"`
	Result []Update `json:"result"`
}


func (t *telegramBot) GetUpdates() (*Updates, error) {
	url := t.baseUrl + "getUpdates"
	updates, err := get(url)
	if err != nil {
		return nil, err
	}

	if err = CheckOk(updates); err != nil {
		return nil, err
	}

	var updatesParsed Updates
	if err := json.Unmarshal(updates, &updatesParsed); err != nil {
		return nil, err
	}

	return &updatesParsed, nil
}

func (t *telegramBot) HandleUpdates() error {
	updates, err := t.GetUpdates()
	if err != nil {
		return err
	}

	for _, update := range updates.Result {
		user := update.Message.From.FirstName
		text := update.Message.Text
		chatId := update.Message.Chat.Id

		var response string
		switch text {
		case "/ping":
			response = "Pong!"
		default:
			fmt.Println(user + ": '" + text + "' , no response") 
			continue
		}

		fmt.Println(user + ": '" + text + "' , response: '" + response + "'")
		if err := t.SendMessage(chatId, response); err != nil {
			fmt.Println(err)
		}
	}

	return nil
}

