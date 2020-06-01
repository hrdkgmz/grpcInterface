package util

import (
	"encoding/json"
	"strings"
)

//DecodeJSON Define
func DecodeJSON(jStr string, target interface{}) error {
	d := json.NewDecoder(strings.NewReader(jStr))
	d.UseNumber()
	err := d.Decode(target)
	if err != nil {
		return err
	}
	return nil
}

//EncodeJSON Define
func EncodeJSON(val interface{}) (string, error) {
	jsonBytes, err := json.Marshal(val)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
