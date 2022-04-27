package routes

import (
	_ "net/http"

	"e-wallet/api/services/auth"
	"e-wallet/api/services/customer"
	"e-wallet/api/services/doc"
	"e-wallet/api/services/user"

	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {

	r.GET("/", auth.WelcomeAPI)
	APIs := r.Group("/api")

	userAPI := APIs.Group("/users")
	{
		userAPI.POST("/register", auth.RegisterUser)
		userAPI.POST("/login", auth.LoginUser)
		userAPI.GET("/", user.GetAllUsers)
		userAPI.GET("/:id", user.GetUserInfo)
	}

	customerAPI := APIs.Group("/customers")
	{
		customerAPI.GET("/", customer.GetAllCustomers)
		customerAPI.GET("/:id", customer.GetCustomerInfo)
		customerAPI.POST("/add", customer.AddCustomerInfo)
		customerAPI.PUT("/update", customer.EditCustomerInfo)
		customerAPI.DELETE("/:id", customer.DeleteCustomerInfo)
	}

	docsAPI := APIs.Group("/docs")
	{

		docsAPI.GET("/:userid", doc.GetDocumentsOfUser)
		docsAPI.POST("/add", doc.AddDocumentInfo)
	}
}
