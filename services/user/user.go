package user

import (
	dbuser "e-wallet/api/db/dbuser"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	var id int
	if err := c.BindJSON(&id); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": err.Error()},
		)
	}

	temp, status := dbuser.GetUserWithID(id)
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
	users, status := dbuser.GetAllUsers()

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
