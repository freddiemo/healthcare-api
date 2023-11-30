package utilstests

import (
	"encoding/json"
	"fmt"
)

func Deserialize(d string) (map[string]interface{}, error) {
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(d), &m); err != nil {
		return nil, fmt.Errorf("failed to deserialize. %w", err)
	}

	return m, nil
}

func DeserializeList(d string) ([]map[string]interface{}, error) {
	var l []map[string]interface{}
	if err := json.Unmarshal([]byte(d), &l); err != nil {
		return nil, fmt.Errorf("failed to deserialize a list. %w", err)
	}

	return l, nil
}
