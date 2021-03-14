package utils

const API_URL = "https://www.clubhouseapi.com/api"

type Clubhouse struct {
	uuid                  string
	userID                int
	authToken             string
	HEADERS               map[string]string
	AUTHENTICATED_HEADERS map[string]string
}
