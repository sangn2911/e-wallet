package main

import (
	dbconn "e-wallet/api/db"
	"e-wallet/api/routes"
	customStatus "e-wallet/api/utils"
	"os"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	customStatus.InitCustomStatus()
	dbconn.StartSqlConnection()
	// run api
	router := gin.Default()
	router.Use(cors.Default())
	routes.ApiRoutes(router)
	router.Run("localhost:" + os.Getenv("PORT"))
}
