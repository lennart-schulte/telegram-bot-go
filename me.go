package main

import (
	"encoding/json"
)

// {"ok":true,"result":{"id":123,"is_bot":true,"first_name":"name_bot","username":"name_bot"}}
type MeInfo struct {
	Ok bool `json:"ok"`
	Result struct {
		Id int `json:"id"`
		IsBot bool `json:"is_bot"`
		FirstName string `json:"first_name"`
		Username string `json:"username"`
	} `json:"result"`
}

func (t *telegramBot) GetMe() (*MeInfo, error) {
	url := t.baseUrl + "getMe"
	getMe, err := get(url)
	if err != nil {
		return nil, err
	}

	if err = CheckOk(getMe); err != nil {
		return nil, err
	}

	var meParsed MeInfo
	if err := json.Unmarshal(getMe, &meParsed); err != nil {
		return nil, err
	}

	return &meParsed, nil
}
