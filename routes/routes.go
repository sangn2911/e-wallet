package routes

import (
	_ "net/http"

	"e-wallet/api/services/affiliate"
	"e-wallet/api/services/auth"
	"e-wallet/api/services/customer"
	"e-wallet/api/services/doc"
	"e-wallet/api/services/transaction"
	"e-wallet/api/services/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func ApiRoutes(r *gin.Engine) {
	APIs := r.Group("/api")

	userAPI := APIs.Group("/user")
	userAPI.Use(cors.Default())
	{
		userAPI.POST("/register", auth.RegisterUser)
		userAPI.POST("/login", auth.LoginUser)
		userAPI.GET("/", user.GetAllUsers)
		userAPI.GET("/:id", user.GetUserInfo)
	}

	customerAPI := APIs.Group("/customer")
	customerAPI.Use(cors.Default())
	{
		customerAPI.GET("", customer.GetAllCustomers)
		customerAPI.GET("/:id", customer.GetCustomerInfo)
		customerAPI.POST("", customer.AddCustomerInfo)
		customerAPI.PUT("", customer.EditCustomerInfo)
		customerAPI.DELETE("", customer.DeleteCustomerInfo)
	}

	docsAPI := APIs.Group("/docs")
	{
		docsAPI.GET("/:userid", doc.GetDocumentsOfUser)
		docsAPI.POST("", doc.AddDocumentInfo)
	}

	transAPI := APIs.Group("/transactions")
	transAPI.Use(cors.Default())
	{
		transAPI.GET("", transaction.GetAllTransactions)
		transAPI.POST("", transaction.AddTransaction)
		transAPI.DELETE("", transaction.DeleteTransaction)
	}

	affisAPI := APIs.Group("/affiliates")
	{
		affisAPI.GET("", affiliate.GetAllAffiliates)
		affisAPI.GET("/:id", affiliate.GetAffiliateInfo)
		affisAPI.POST("", affiliate.AddAffiliate)
		affisAPI.DELETE("", affiliate.DeleteAffiliate)
	}
}
