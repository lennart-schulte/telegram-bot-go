package main

import (
	"encoding/json"
	"errors"
)


// {"ok":false,"error_code":401,"description":"Unauthorized"}
/*type getFailed struct {
	Ok bool `json:"ok"`
	ErrorCode int `json:"error_code"`
	Description string `json:"description"`
}*/

func CheckOk(getResponse []byte) error {
	// checks the "ok" field in the response
	var getInterface interface{}
	err := json.Unmarshal(getResponse, &getInterface)
	if err != nil {
		return err
	}
	if !getInterface.(map[string]interface{})["ok"].(bool) {
		return errors.New("Request was NOT okay " + string(getResponse))
	}
	return nil
}
