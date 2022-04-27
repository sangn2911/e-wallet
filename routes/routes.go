package routes

import (
	_ "net/http"

	"e-wallet/api/services/auth"
	"e-wallet/api/services/user"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {

	APIs := r.Group("/api")

	userAuth := APIs.Group("/users")
	{
		userAuth.POST("/register", auth.RegisterUser)
		userAuth.POST("/login", auth.LoginUser)
		userAuth.GET("/", user.GetAllUsers)
		userAuth.GET("/:id", user.GetUserInfo)
	}
	customerAuth := APIs.Group("/customers")
	{
		customerAuth.GET("/", user.GetAllCustomers)
		customerAuth.GET("/:id", user.GetCustomerInfo)
		customerAuth.POST("/add", user.AddCustomers)
	}
}
