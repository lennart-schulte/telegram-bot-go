package main

import (
	"encoding/json"
)

type sendPayload struct {
	ChatId int `json:"chat_id"`
	Text string `json:"text"`
}

func (t *telegramBot) SendMessage(chatId int, text string) error {
	url := t.baseUrl + "sendMessage"

	payload := sendPayload{
		ChatId: chatId,
		Text: text,
	}
	payloadJson, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := post(url, payloadJson)
	if err != nil {
		return err
	}

	return CheckOk(resp)
}
