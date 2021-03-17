package utils

import (
	"inssa_club_clubhouse_backend/configs"

	mgm "github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() {
	err := mgm.SetDefaultConfig(nil, "clubhouse", options.Client().ApplyURI(configs.Envs["MONGO_URI"]))
	if err != nil {
		panic(err)
	}
}
