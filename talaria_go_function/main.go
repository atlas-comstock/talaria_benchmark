package main

import "encoding/json"

var keysDefined = []string{
	"time", "data", "id", "uuid", "ios", "android", "apple",
}

func isKeyInDefinded(key string) bool {
	for _, k := range keysDefined {
		if key == k {
			return true
		}
	}
	return false
}

//save it in data
func ComputeRow(rowArg map[string]interface{}) interface{} {
	data := make(map[string]interface{})
	for key, val := range rowArg {
		if !isKeyInDefinded(key) {
			data[key] = val
			delete(rowArg, key)
		}
	}
	res, _ := json.Marshal(data)
	return string(res)
}
