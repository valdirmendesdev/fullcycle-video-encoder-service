package utils

import json2 "encoding/json"

func IsJson(s string) error {
	var json struct{}

	if err := json2.Unmarshal([]byte(s), &json); err != nil {
		return err
	}
	return nil
}