package customer

import (
	dbcustomer "e-wallet/api/db/dbcustomer"
	"e-wallet/api/objects"
	CustomStatus "e-wallet/api/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCustomerInfo(c *gin.Context) {
	var id string = c.Param("id")
	// if err := c.BindJSON(&id); err != nil {
	// 	c.JSON(
	// 		http.StatusBadRequest,
	// 		gin.H{"status": http.StatusBadRequest, "error": err.Error()},
	// 	)
	// }

	temp, status := dbcustomer.GetCustomerWithID(id)
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
	customers, status := dbcustomer.GetAllCustomers()
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

func AddCustomerInfo(c *gin.Context) {
	var customer objects.Customer
	if err := c.BindJSON(&customer); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": err.Error()},
		)
	}

	customer, status := dbcustomer.InsertCustomer(customer)
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
				"data":   customer,
			},
		)
	}
}

func EditCustomerInfo(c *gin.Context) {
	var customers objects.Customer
	if err := c.BindJSON(&customers); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": err.Error()},
		)
	}

	customer, status := dbcustomer.EditCustomerInfo(customers)
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
				"data":   customer,
			},
		)
	}
}

func DeleteCustomerInfo(c *gin.Context) {
	var id int
	if err := c.BindJSON(&id); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": err.Error()},
		)
	}

	status := dbcustomer.DeleteCustomer(id)
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
			},
		)
	}
}
