package models

import (
	"github.com/kamva/mgm/v3"
)

// ClubhouseProfile is a profile of clubhouse
type ClubhouseProfile struct {
	mgm.DefaultModel `bson:",inline"`
	UserID           uint64                 `json:"user_id" bson:"user_id"`
	Profile          map[string]interface{} `json:"profile" bson:"profile"`
}
