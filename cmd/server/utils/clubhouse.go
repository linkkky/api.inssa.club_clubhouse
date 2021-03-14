package utils

import (
	"net/http"
)

const API_URL = "https://www.clubhouseapi.com/api"

type Clubhouse struct {
	uuid                  string
	userID                int
	authToken             string
	HEADERS               map[string]string
	AUTHENTICATED_HEADERS map[string]string
}

func (clubhouse Clubhouse) request(req *http.Request) (map[string]interface{}, error) {
	client := &http.Client{}
	for key, value := range clubhouse.HEADERS {
		req.Header.Add(key, value)
	}
	// add Headers
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return responseBodyToMap(resp.Body)
}

func (clubhouse Clubhouse) authenticatedRequest(req *http.Request) (map[string]interface{}, error) {
	client := &http.Client{}
	for key, value := range clubhouse.HEADERS {
		req.Header.Add(key, value)
	}
	for key, value := range clubhouse.AUTHENTICATED_HEADERS {
		req.Header.Add(key, value)
	}
	// add Headers
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return responseBodyToMap(resp.Body)
}
