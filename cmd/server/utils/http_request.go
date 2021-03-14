package utils

import (
	"io"
	"io/ioutil"
)

func responseBodyToString(body io.ReadCloser) string {
	defer body.Close()
	bytes, _ := ioutil.ReadAll(body)
	strBody := string(bytes)
	return strBody
}
