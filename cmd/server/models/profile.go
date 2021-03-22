package models

import (
	"time"

	"github.com/kamva/mgm/v3"
)

// ClubhouseProfile is a profile of clubhouse
type ClubhouseProfile struct {
	mgm.DefaultModel `bson:",inline"`
	UserID           uint64                 `json:"user_id" bson:"user_id"`
	Username         string                 `json:"username" bson:"username"`
	Profile          map[string]interface{} `json:"profile" bson:"profile"`
}

// ClubhouseProfileResponse is a struct for documentation
type ClubhouseProfileResponse struct {
	ID        string    `json:"id" example:"604f7b3253883ed1649ed13a"`
	CreatedAt time.Time `json:"created_at" example:"2021-03-15T15:20:18.066663Z"`
	UpdatedAt time.Time `json:"updated_at" example:"2021-03-15T15:20:18.066664Z"`
	UserID    int       `json:"user_id" example:"711498010"`
	Username  string    `json:"username" example:"yeon.gyu.kim"`
	Profile   struct {
		Name         string `json:"name" example:"YeonGyu Kim"`
		Username     string `json:"username" example:"yeon.gyu.kim"`
		Instagram    string `json:"instagram" example:"yeon.gyu.kim"`
		Twitter      string `json:"twitter"`
		NumFollowers int    `json:"num_followers" example:"84"`
		NumFollowing int    `json:"num_following" example:"163"`
		Bio          string `json:"bio" example:"BackEnd Programmer. 선린인터넷고 정보보호과 3학년 재학중입니다.\n꽈뚜룹, 이근 성대모사 잘합니다.\n\nPython, Node, Go 관심 많습니다."`
		PhotoURL     string `json:"photo_url" example:"https://clubhouseprod.s3.amazonaws.com/711498010_dcbb5803-9985-4ec5-b36a-0ac86284ff92"`
		Clubs        []struct {
			ClubID              int    `json:"club_id" example:"12345"`
			Description         string `json:"description" example:"this is a strip club"`
			EnablePrivate       bool   `json:"enable_private" example:"true"`
			IsCommunity         bool   `json:"is_community" example:"false"`
			IsFollowAllowed     bool   `json:"is_follow_allowed" example:"true"`
			IsMembershipPrivate bool   `json:"is_membership_private" example:"false"`
			Name                string `json:"name" example:"strip club"`
			NumFollowers        int    `json:"num_followers" example:"100"`
			NumMembers          int    `json:"num_members" example:"20"`
			NumOnline           int    `json:"num_online" example:"3"`
			PhotoURL            string `json:"photo_url" example:"https://clubhouseprod.s3.amazonaws.com:443/club_<club_id>_<guid>_thumbnail_250x250"`
			Rules               []struct {
				Desc  string `json:"desc" example:"description"`
				Title string `json:"title" example:"title"`
			} `json:"rules"`
			URL string `json:"url" example:"https://www.joinclubhouse.com/club/<clubname>"`
		} `json:"clubs"`
		Displayname          string `json:"displayname"`
		InvitedByUserProfile struct {
			Name     string `json:"name" example:"Go Taegeon"`
			PhotoURL string `json:"photo_url" example:"https://clubhouseprod.s3.amazonaws.com/5652354_2c1840a1-02c0-40ee-b98c-f45598712bcd_thumbnail_250x250"`
			UserID   int    `json:"user_id" example:"5652354"`
			Username string `json:"username" example:"gtg7784"`
		} `json:"invited_by_user_profile"`
		TimeCreated time.Time `json:"time_created" example:"2021-02-08T15:03:08.132077+00:00"`
		URL         string    `json:"url" example:"https://www.joinclubhouse.com/@yeon.gyu.kim"`
	} `json:"profile"`
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
	clubhouseProfile.Username = profile["username"].(string)
	profileCleaner(profile, []string{"user_id", "follows_me", "invited_by_club", "is_blocked_by_network", "mutual_follows", "mutual_follows_count", "notification_type"})
	clubhouseProfile.Profile = profile
	return &clubhouseProfile
}
