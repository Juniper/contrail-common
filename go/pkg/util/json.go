package util

import "encoding/json"

func MustJSON(data interface{}) string {
	b, _ := json.Marshal(data)
	return string(b)
}
