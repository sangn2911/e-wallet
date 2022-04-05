package handler

import (
	"api/e-wallet/db"
	"api/e-wallet/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCustomers(c *gin.Context) {
	res, err := db.GetCustomers()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't find customers")
		return
	}
	fmt.Println("Get all customers information")
	c.IndentedJSON(http.StatusOK, res)
	return
}
func GetCustomer(c *gin.Context) {
	id := c.Param("id")
	tmp, _ := strconv.Atoi(id)
	res, err := db.GetCustomerByID(tmp)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't find user")
		return
	}
	fmt.Println("Customer id", id)
	c.IndentedJSON(http.StatusOK, res)
	return
}
func AddCustomer(c *gin.Context) {
	var customer entities.Customer
	if err := c.BindJSON(&customer); err != nil {
		return
	}
	res, err := db.AddCustomer(customer.FirstName, customer.LastName, customer.DateOfBirth, customer.Email, customer.Nationality, customer.Address)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't add customer")
		return
	}
	c.IndentedJSON(http.StatusOK, res)
	fmt.Println("data", customer.Email, customer.LastName)
	//res := c.BindJSON("lastName")
	//fmt.Println(res)
	return
}
func DeleteCustomer(c *gin.Context) {
	var id int
	if err := c.BindJSON(&id); err != nil {
		return
	}
	err := db.DeleteCustomer(strconv.Itoa(id))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't find ID")
		return
	}
	c.IndentedJSON(http.StatusOK, id)
	fmt.Println(id)
	return
}
func EditCustomer(c *gin.Context) {
	var customer entities.Customer
	if err := c.BindJSON(&customer); err != nil {
		return
	}
	id := c.Param("id")
	tmp, _ := strconv.Atoi(id)
	res, err := db.EditCustomer(tmp, customer.FirstName, customer.LastName, customer.DateOfBirth, customer.Email, customer.Nationality, customer.Address)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't add customer")
		return
	}
	c.IndentedJSON(http.StatusOK, res)
	fmt.Println("data", customer.Email, customer.LastName)
	//res := c.BindJSON("lastName")
	//fmt.Println(res)
	return
}
