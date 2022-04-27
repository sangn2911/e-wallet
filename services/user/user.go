package user

import (
	dbhandler "e-wallet/api/db"
	"e-wallet/api/objects"
	CustomStatus "e-wallet/api/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	id := c.Param("id")

	temp, status := dbhandler.GetUserWithID(id)

	if status != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": status.Error()},
		)
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": http.StatusOK,
				"error":  "",
				"data": map[string]interface{}{
					"username": temp.Username,
					"email":    temp.Email,
				},
				// "token":    tokenStr,
			},
		)
	}
}

func GetAllUsers(c *gin.Context) {

	users, status := dbhandler.GetAllUsers()

	if status != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": status.Error()},
		)
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": http.StatusOK,
				"error":  "",
				"data":   users,
				// "token":    tokenStr,
			},
		)
	}
}

func GetCustomerInfo(c *gin.Context) {
	id := c.Param("id")

	temp, status := dbhandler.GetCustomerWithID(id)

	if status != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": status.Error()},
		)
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": http.StatusOK,
				"error":  "",
				"data":   temp,
				// "token":    tokenStr,
			},
		)
	}
}

func GetAllCustomers(c *gin.Context) {

	customers, status := dbhandler.GetAllCustomers()

	if status != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": status.Error()},
		)
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": http.StatusOK,
				"error":  "",
				"data":   customers,
			},
		)
	}
}

func AddCustomers(c *gin.Context) {
	var customers objects.Customer

	if err := c.BindJSON(&customers); err != nil {
		return
	}

	customers, status := dbhandler.InsertCustomer(customers)

	if status != nil {
		if errors.Is(status, CustomStatus.ExistCustomer) {
			c.JSON(
				http.StatusOK,
				gin.H{"status": http.StatusOK, "error": CustomStatus.UsedEmail.Error()},
			)
		} else {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"status": http.StatusBadRequest, "error": status.Error()},
			)
		}
	} else {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": http.StatusOK,
				"error":  "",
			},
		)
	}
}
