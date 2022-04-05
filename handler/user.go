package handler

import (
	"api/e-wallet/db"
	"api/e-wallet/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	id := c.Param("id")
	tmp, _ := strconv.Atoi(id)
	res, err := db.GetUserById(tmp)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't find user")
		return
	}
	fmt.Println("User id", id)
	c.IndentedJSON(http.StatusOK, res)
	return
}
func GetUsers(c *gin.Context) {
	res, err := db.GetUsers()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't find user")
		return
	}
	fmt.Println("Get all user")
	c.IndentedJSON(http.StatusOK, res)
	return
}
func AddUser(c *gin.Context) {
	var user entities.User
	if err := c.BindJSON(&user); err != nil {
		return
	}
	res, err := db.AddUser(user.Username, user.Email, user.Password)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't add User")
		return
	}
	c.IndentedJSON(http.StatusOK, res)

	//res := c.BindJSON("lastName")
	//fmt.Println(res)
	return
}
