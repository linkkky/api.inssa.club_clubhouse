package test

import (
	"inssa_club_clubhouse_backend/cmd/server/utils"
	"inssa_club_clubhouse_backend/configs"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const notExistingUsername = "notexistinguser!@#$%^&*()(*&^%$#@!"

var clubhouse = utils.SingletonClubhouse()

func setupClubhouse() {
	UUID := configs.Envs["CLUBHOUSE_ACCOUNT_UUID"]
	USER_ID, err := strconv.Atoi(configs.Envs["CLUBHOUSE_ACCOUNT_USER_ID"])
	if err != nil {
		panic(err)
	}
	AUTH_TOKEN := configs.Envs["CLUBHOUSE_ACCOUNT_AUTH_TOKEN"]
	utils.SingletonClubhouse().SetAccount(UUID, USER_ID, AUTH_TOKEN)
}

// tools for convenient test

func getUserIDByUsernameWithWrongAuthTokenTest(t *testing.T) {
	clubhouse.SetAccount("123", 123, "123")
	_, err := clubhouse.GetUserIDByUsername("yeon.gyu.kim")
	assert.NotEqual(t, nil, err)
	setupClubhouse() // restore to original clubhouse info
}

func getUserIDByUsernameOfNotExistingUserTest(t *testing.T) {
	_, err := clubhouse.GetUserIDByUsername(notExistingUsername)
	assert.NotEqual(t, nil, err)
}

func getUserIDByUsernameTest(t *testing.T) {
	_, err := clubhouse.GetUserIDByUsername("yeon.gyu.kim")
	assert.Equal(t, nil, err)
}

// GetUserIDByUsername tests

func getProfileByUserIDWithWrongAuthTokenTest(t *testing.T) {
	clubhouse.SetAccount("123", 123, "123")
	_, err := clubhouse.GetProfileByUserID("711498010")
	assert.NotEqual(t, nil, err)
	setupClubhouse()
}

func getProfileByUserIDOfNotExistingUserTest(t *testing.T) {
	_, err := clubhouse.GetProfileByUserID("-1")
	assert.NotEqual(t, nil, err)
}

func getProfileByUserIDTest(t *testing.T) {
	_, err := clubhouse.GetProfileByUserID("711498010")
	assert.Equal(t, nil, err)
}

// GetProfileByUserID tests

func getProfileByUsernameWithWrongAuthTokenTest(t *testing.T) {
	clubhouse.SetAccount("123", 123, "123")
	_, err := clubhouse.GetProfileByUserID("yeon.gyu.kim")
	assert.NotEqual(t, nil, err)
	setupClubhouse()
}

func getProfileByUsernameOfNotExistingUserTest(t *testing.T) {
	_, err := clubhouse.GetUserIDByUsername(notExistingUsername)
	assert.NotEqual(t, nil, err)
}

func getProfileByUsernameTest(t *testing.T) {
	_, err := clubhouse.GetProfileByUsername("yeon.gyu.kim")
	assert.Equal(t, nil, err)
}

// GetProfileByUserID tests

func TestClubhouse(t *testing.T) {
	t.Run("GetUserIDByUsername", func(t *testing.T) {
		t.Run("Retrieving a profile with wrong auth token", getUserIDByUsernameWithWrongAuthTokenTest)
		t.Run("Retrieving the profile of not existing user should fail", getUserIDByUsernameOfNotExistingUserTest)
		t.Run("Retrieving the profile of yeon.gyu.kim should success", getUserIDByUsernameTest)
	})
	t.Run("GetProfileByUserID", func(t *testing.T) {
		t.Run("Retrieving a profile with wrong auth token", getProfileByUserIDWithWrongAuthTokenTest)
		t.Run("Retrieving the profile of not existing user should fail", getProfileByUserIDOfNotExistingUserTest)
		t.Run("Retrieving the profile of 711498010(yeon.gyu.kim) should success", getProfileByUserIDTest)
	})
	t.Run("GetProfileByUsername", func(t *testing.T) {
		t.Run("Retrieving a profile with wrong auth token", getProfileByUsernameWithWrongAuthTokenTest)
		t.Run("Retrieving the profile of not existing user should fail", getProfileByUsernameOfNotExistingUserTest)
		t.Run("Retrieving the profile of 711498010(yeon.gyu.kim) should success", getProfileByUsernameTest)
	})
}
