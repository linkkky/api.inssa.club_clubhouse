package utils

import (
	"fmt"
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

func NewClubhouse(uuid string, userID int, authToken string) *Clubhouse {
	clubhouse := Clubhouse{uuid: uuid, userID: userID, authToken: authToken}
	clubhouse.HEADERS = map[string]string{
		"User-Agent":    "clubhouse/269 (iPhone; iOS 14.1; Scale/3.00)",
		"CH-Languages":  "en-US",
		"CH-Locale":     "en_US",
		"CH-AppVersion": "0.2.15",
		"CH-AppBuild":   "269",
		"CH-DeviceId":   uuid,
		"Content-Type":  "application/json",
	}
	clubhouse.AUTHENTICATED_HEADERS = map[string]string{
		"CH-UserID":     fmt.Sprintf("%d", userID),
		"Authorization": fmt.Sprintf("Token %s", authToken),
	}
	return &clubhouse
}
