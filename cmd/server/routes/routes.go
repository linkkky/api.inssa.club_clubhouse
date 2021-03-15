package routes

import (
	"inssa_club_clubhouse_backend/cmd/server/controllers"

	"github.com/gin-gonic/gin"
)

// RouteInfo is a struct for a route
type RouteInfo struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

// GetRoutes is a function which returns the information of routes
func GetRoutes() []RouteInfo {
	c := controllers.NewController()
	routeInfo := []RouteInfo{
		{Method: "GET", Path: "/profile/:username", Handler: c.GetProfile},
	}
	return routeInfo
}
