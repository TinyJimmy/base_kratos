package utils

import (
	"encoding/json"
)

func MapToStruct(m map[string]interface{}, result interface{}) error {
	am, _ := json.Marshal(m)
	return json.Unmarshal(am, &result)
}
