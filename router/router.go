package router

import (
	"api/e-wallet/handler"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func Router() {
	router := gin.Default()
	port := os.Getenv("PORT")
	address := fmt.Sprintf("%s:%s", "0:0:0:0", port)
	fmt.Println(address)
	router.Use(cors.Default())
	router.GET("/api/user", handler.GetUsers)
	router.GET("/api/user/:id", handler.GetUser)
	router.GET("/api/customer", handler.GetCustomers)
	router.POST("api/customer", handler.AddCustomer)
	router.GET("/api/customer/:id", handler.GetCustomer)
	router.PUT("api/customer/:id", handler.EditCustomer)
	router.DELETE("api/customer", handler.DeleteCustomer)
	router.GET("api/transactions", handler.GetTransactions)
	router.POST("api/transactions", handler.AddTransaction)
	router.DELETE("api/transactions", handler.DeleteTransaction)
	router.GET("api/affiliates", handler.GetAffiliates)
	router.GET("api/affiliates/info/:id", handler.GetAffiliate)
	router.POST("api/affiliates", handler.AddAffiliate)
	router.DELETE("api/affiliates", handler.DeleteAffiliate)
	router.POST("api/user/register", handler.AddUser)
	router.GET("/api/document/:id", handler.GetDocument)
	router.POST("/api/document/:id", handler.AddDocument)
	fmt.Println("Server is running...")
	router.Run(":" + port)
}
