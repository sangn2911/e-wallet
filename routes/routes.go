package routes

import (
	_ "net/http"

	auth "e-wallet/api/services"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {

	APIs := r.Group("/api")

	userAuth := APIs.Group("/user")
	{
		userAuth.POST("/register", auth.RegisterUser)
		userAuth.POST("/login", auth.LoginUser)
		userAuth.GET("/:id", auth.GetUserInfo)
	}
}
