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

func profileCleaner(profile map[string]interface{}, fields []string) {
	for _, field := range fields {
		delete(profile, field)
	}
}

// NewClubhouseProfile is a constructor for ClubhouseProfile
func NewClubhouseProfile(profile map[string]interface{}) *ClubhouseProfile {
	clubhouseProfile := ClubhouseProfile{}
	clubhouseProfile.UserID = uint64(profile["user_id"].(float64))
	profileCleaner(profile, []string{"user_id", "follows_me", "invited_by_club", "is_blocked_by_network", "mutual_follows", "mutual_follows_count", "notification_type"})
	clubhouseProfile.Profile = profile
	return &clubhouseProfile
}
