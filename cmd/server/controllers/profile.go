package controllers

import (
	"fmt"
	"inssa_club_clubhouse_backend/cmd/server/models"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func getProfileFromDB(username string) (models.ClubhouseProfile, error) {
	result := []models.ClubhouseProfile{}
	err := mgm.Coll(&models.ClubhouseProfile{}).SimpleFind(&result, bson.M{"username": username})
	if err != nil {
		return models.ClubhouseProfile{}, err
	}
	if len(result) == 0 {
		return models.ClubhouseProfile{}, fmt.Errorf("No such user")
	}
	return result[0], nil
}
