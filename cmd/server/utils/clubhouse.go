package utils

import (
	"fmt"
	"net/http"
)

const API_URL = "https://www.clubhouseapi.com/api"

func extractUserIDByUsername(users []interface{}, username string) (string, error) {
	for _, user := range users {
		user_info := user.(map[string]interface{})
		if username == user_info["username"] {
			return user_info["user_id"].(string), nil
		}
	}
	return "", fmt.Errorf("No such user")
}

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

func (clubhouse Clubhouse) GetUserIDByUsername(username string) (string, error) {
	const SEARCH_ENDPOINT = "/search_users"

	body := mapToBody(map[string]interface{}{"query": username})
	req, err := http.NewRequest("POST", API_URL+SEARCH_ENDPOINT, body)
	if err != nil {
		return "", err
	}

	resp, err := clubhouse.authenticatedRequest(req)
	if err != nil {
		return "", err
	}

	users := (resp["users"].([]interface{}))
	return extractUserIDByUsername(users, username)
}

func (clubhouse Clubhouse) GetProfileByUserID(userID string) (map[string]interface{}, error) {
	const GET_PROFILE_ENDPOINT = "/get_profile"

	body := mapToBody(map[string]interface{}{"user_id": userID})
	req, err := http.NewRequest("POST", API_URL+GET_PROFILE_ENDPOINT, body)
	if err != nil {
		return nil, err
	}

	profile, err := clubhouse.authenticatedRequest(req)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (clubhouse Clubhouse) GetProfileByUsername(username string) (map[string]interface{}, error) {
	userID, err := clubhouse.GetUserIDByUsername(username)
	if err != nil {
		return nil, err
	}
	return clubhouse.GetProfileByUserID(userID)
}
