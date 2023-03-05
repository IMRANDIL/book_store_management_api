package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)



func ParseBody(r *http.Request, x interface{}) error {
	if r.Method != "POST" && r.Method != "PUT" {
		return fmt.Errorf("Invalid request method: %s", r.Method)
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(x); err != nil {
		return err
	}

	return nil
}
