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

func setupClubhouseProperly() {
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
	setupClubhouseProperly()
}
