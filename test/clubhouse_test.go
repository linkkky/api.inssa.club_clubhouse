package test

import (
	"inssa_club_clubhouse_backend/cmd/server/utils"
	"inssa_club_clubhouse_backend/configs"
	"strconv"
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
