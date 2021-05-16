package utils

import (
	"fmt"
	"net/http"
	"sync"
)

const API_URL = "https://www.clubhouseapi.com/api"

var clubhouseInstance *Clubhouse
var once sync.Once

func extractUserIDByUsername(users []interface{}, username string) (string, error) {
	for _, user := range users {
		user_info := user.(map[string]interface{})
		if username == user_info["username"] {
			return user_info["user_id"].(string), nil
		}
	}
	return "", fmt.Errorf("no such user")
}

func extractUserListFromSearchResult(result map[string]interface{}) ([]interface{}, error) {
	for key, value := range result {
		if key == "users" {
			return value.([]interface{}), nil
		}
	}
	return nil, fmt.Errorf(fmt.Sprintf("Error while retrieving the profile. %+v", result))
}

// Clubhouse is a api client for the clubhouse API
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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", resp.Status)
	}
	return responseBodyToMap(resp.Body)
}

func (clubhouse Clubhouse) authenticatedRequest(req *http.Request) (map[string]interface{}, error) {
	for key, value := range clubhouse.AUTHENTICATED_HEADERS {
		req.Header.Add(key, value)
	}
	// add Headers
	return clubhouse.request(req)
}

// SingletonClubhouse returns clubhouseInstance, using thread-safe singleton pattern
func SingletonClubhouse() *Clubhouse {
	once.Do(func() {
		clubhouseInstance = &Clubhouse{}
	})
	return clubhouseInstance
}

// SetAccount sets basic information of clubhouse
func (clubhouse *Clubhouse) SetAccount(uuid string, userID int, authToken string) {
	clubhouse.uuid = uuid
	clubhouse.userID = userID
	clubhouse.authToken = authToken
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
}

// GetUserIDByUsername retrieves userID from clubhouse API by given username
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

	users, err := extractUserListFromSearchResult(resp)
	if err != nil {
		return "", nil
	}

	return extractUserIDByUsername(users, username)
}

// GetProfileByUserID retrieves profile from clubhouse API by given userID
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
	return profile["user_profile"].(map[string]interface{}), nil
}

// GetProfileByUsername retrieves profile from clubhouse API by given username
func (clubhouse Clubhouse) GetProfileByUsername(username string) (map[string]interface{}, error) {
	userID, err := clubhouse.GetUserIDByUsername(username)
	if err != nil {
		return nil, err
	}
	return clubhouse.GetProfileByUserID(userID)
}
