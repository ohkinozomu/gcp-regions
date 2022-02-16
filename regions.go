package regions

import (
	"encoding/json"
)

var list map[string]Region

func init() {
	list, _ = List()
}

func List() (map[string]Region, error) {
	data := []byte(REGION_DATA)

	var m map[string]Region
	if err := json.Unmarshal(data, &m); err != nil {
		return m, err
	}

	return m, nil
}

func IsValid(name string) bool {
	if _, ok := list[name]; ok {
		return true
	}
	return false
}
