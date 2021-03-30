package controllers

import (
	"errors"
	"inssa_club_clubhouse_backend/cmd/server/models"
	"inssa_club_clubhouse_backend/cmd/server/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

func getProfileMapFromServer(username string) (map[string]interface{}, error) {
	// if the cached data is not usable, query the data and cache it
	clubhouse := utils.SingletonClubhouse()
	profile, err := clubhouse.GetProfileByUsername(username)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func getProfileFromDB(username string) (*models.ClubhouseProfile, error) {
	result := []models.ClubhouseProfile{}
	err := mgm.Coll(&models.ClubhouseProfile{}).SimpleFind(&result, bson.M{"username": username})
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, errors.New("no such user")
	}
	return &result[0], nil
}

func syncProfile(profileDocument *models.ClubhouseProfile) error {
	recentProfile, err := getProfileMapFromServer(profileDocument.Username)
	if err != nil {
		return err
	}
	profileDocument.Profile = recentProfile
	err = mgm.Coll(profileDocument).Update(profileDocument)
	return err
}

func getProfile(username string) (*models.ClubhouseProfile, error) {
	// get data from db
	profileDocument, err := getProfileFromDB(username)

	if err == nil { // if found the data from db
		updatedAt := profileDocument.UpdatedAt.Add(time.Hour * 2)
		currentTime := time.Now()
		if currentTime.Before(updatedAt) { // if the document updated within 2 hours
			return profileDocument, nil
		} else { // if the document is old
			err := syncProfile(profileDocument)
			return profileDocument, err
		}
	} else { // if the cached data is not usable, query the data and cache it
		profileMap, err := getProfileMapFromServer(username)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		profileDocument = models.NewClubhouseProfile(profileMap)
		err = mgm.Coll(profileDocument).Create(profileDocument)
		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		return profileDocument, nil
		// validate updated_at
	}
}

// GetProfile godoc
// @Summary Retrieve a profile from clubhouse by given username
// @Description Retrieve a profile from clubhouse by given username
// @Accept  json
// @Produce  json
// @Param username path string true "Username"
// @Success 200 {object} models.ClubhouseProfileResponse
// @Failure 404
// @Router /profile/{username} [get]
func (ctrler *Controller) GetProfile(c *gin.Context) {
	USERNAME := c.Param("username")
	clubhouseProfile, err := getProfile(USERNAME)
	if err != nil {
		c.Data(http.StatusNotFound, gin.MIMEHTML, nil)
		return
	}
	c.JSON(http.StatusOK, clubhouseProfile)
}
