package util

import (
	"encoding/json"
	"fmt"
)

type MetaMarshal interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

func ConvertVMMeta(meta MetaMarshal) map[string]interface{} {
	var data map[string]interface{}
	b, _ := meta.MarshalJSON()
	err := json.Unmarshal(b, &data)
	if err != nil {
		fmt.Printf("error while unmarshalling meta: %v", err)
		return nil
	}
	return data
}

func ConvertMapToTags(meta MetaMarshal) map[string]string {
	data := ConvertVMMeta(meta)

	tags := make(map[string]string)

	for k, v := range data {
		tags[k] = fmt.Sprint(v)
	}
	return tags
}
