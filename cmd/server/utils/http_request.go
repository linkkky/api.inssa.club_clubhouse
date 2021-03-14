package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
)

func responseBodyToString(body io.ReadCloser) string {
	defer body.Close()
	bytes, _ := ioutil.ReadAll(body)
	strBody := string(bytes)
	return strBody
}

func responseBodyToMap(body io.ReadCloser) (map[string]interface{}, error) {
	resp := map[string]interface{}{}
	err := json.Unmarshal([]byte(responseBodyToString(body)), &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func mapToBody(mapBody map[string]interface{}) *bytes.Buffer {
	marshaledBody, _ := json.Marshal(mapBody)
	return bytes.NewBuffer(marshaledBody)
}
