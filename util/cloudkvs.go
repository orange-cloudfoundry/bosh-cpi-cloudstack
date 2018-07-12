package util

import (
	"github.com/cppforlife/bosh-cpi-go/apiv1"
	"encoding/json"
	"fmt"
)

func ConvertVMMeta(meta apiv1.VMMeta) map[string]interface{} {
	var data map[string]interface{}
	b, _ := meta.MarshalJSON()
	json.Unmarshal(b, &data)
	return data
}

func ConvertMapToTags(meta apiv1.VMMeta) map[string]string {
	data := ConvertVMMeta(meta)

	tags := make(map[string]string)

	for k, v := range data {
		tags[k] = fmt.Sprint(v)
	}
	return tags
}
