package main

import (
	"inssa_club_clubhouse_backend/cmd/server/docs"
	"inssa_club_clubhouse_backend/cmd/server/middlewares"
	"inssa_club_clubhouse_backend/cmd/server/routes"
	"inssa_club_clubhouse_backend/cmd/server/utils"
	"inssa_club_clubhouse_backend/configs"
	"strconv"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRoutes(engine *gin.Engine) {
	IS_ENABLE_SWAGGER := configs.Envs["IS_ENABLE_SWAGGER"]
	if IS_ENABLE_SWAGGER == "true" {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	for _, controller := range routes.GetRoutes() {
		engine.Handle(controller.Method, controller.Path, controller.Handler)
	} // setup routes
}

func setupClubhouseAccount() {
	UUID := configs.Envs["CLUBHOUSE_ACCOUNT_UUID"]
	USER_ID, err := strconv.Atoi(configs.Envs["CLUBHOUSE_ACCOUNT_USER_ID"])
	if err != nil {
		panic(err)
	}
	AUTH_TOKEN := configs.Envs["CLUBHOUSE_ACCOUNT_AUTH_TOKEN"]
	logrus.WithFields(logrus.Fields{"UUID": UUID, "USER_ID": USER_ID, "AUTH_TOKEN": AUTH_TOKEN}).Info("Clubhouse Account")
	utils.SingletonClubhouse().SetAccount(UUID, USER_ID, AUTH_TOKEN)
}

func setupDocuments() {
	docs.SwaggerInfo.Title = "clubhouse-api.inssa.club"
	docs.SwaggerInfo.Description = "The REST API for clubhouse service of api.inssa.club"
	docs.SwaggerInfo.Host = "api.inssa.club"
	docs.SwaggerInfo.Version = "0.1"
	docs.SwaggerInfo.BasePath = "/clubhouse"
	docs.SwaggerInfo.Schemes = []string{"https"}
}

func runServer(engine *gin.Engine) {
	MODE := configs.Envs["GIN_MODE"]
	PORT := ":" + configs.Envs["SERVER_PORT"]

	if MODE == "release" {
		gateway.ListenAndServe(PORT, engine)
	} else {
		engine.Run(PORT)
	}
}

func main() {
	configs.InitDB()
	engine := gin.Default()
	setupClubhouseAccount()
	middlewares.Setup(engine)
	setupDocuments()
	setupRoutes(engine)
	runServer(engine)
}
