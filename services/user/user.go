package user

import (
	dbhandler "e-wallet/api/db"
	CustomStatus "e-wallet/api/utils"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	id := c.Param("id")

	temp, status := dbhandler.GetUserWithID(id)

	if errors.Is(status, CustomStatus.ExistUser) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status":   http.StatusOK,
				"error":    "",
				"username": temp.Username,
				"email":    temp.Email,
				// "token":    tokenStr,
			},
		)
	} else {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": status.Error()},
		)
	}
}

func GetAllUsers(c *gin.Context) {

	users, status := dbhandler.GetAllUsers()

	if errors.Is(status, CustomStatus.ExistUser) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"status": http.StatusOK,
				"error":  "",
				"users":  users,
				// "token":    tokenStr,
			},
		)
	} else {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": status.Error()},
		)
	}
}
