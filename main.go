package main

import (
	dbhandler "e-wallet/api/db"
	"e-wallet/api/routes"
	customStatus "e-wallet/api/utils"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	customStatus.InitCustomStatus()
	dbhandler.StartSqlConnection()
	// run api
	router := gin.Default()
	router.Use(cors.Default())
	routes.ApiRoutes(router)
	router.Run("localhost:" + os.Getenv("PORT"))
}
