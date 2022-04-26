package auth

import (
	dbhandler "e-wallet/api/db"
	"e-wallet/api/objects"
	CustomStatus "e-wallet/api/utils"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/golang-jwt/jwt"
)

// func GenerateJWTAuthentication(username string, password string) (string, error) {
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	claims := token.Claims.(jwt.MapClaims)

// 	claims["authorized"] = true
// 	claims["user"] = username
// 	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()

// 	tokenStr, err := token.SignedString([]byte("noice"))

// 	if err != nil {
// 		return "", fmt.Errorf("token sign error: %v", err)
// 	}

// 	return tokenStr, nil
// }

func RegisterUser(c *gin.Context) {
	var user objects.User

	if err := c.BindJSON(&user); err != nil {
		return
	}

	user, status := dbhandler.RegisterUser(user)

	if status != nil {
		if errors.Is(status, CustomStatus.ExistUser) {
			c.JSON(
				http.StatusOK,
				gin.H{"status": http.StatusOK, "error": status.Error()},
			)
		} else {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"status": http.StatusBadRequest, "error": status.Error()},
			)
		}
	} else {
		// tokenStr, err := GenerateJWTAuthentication(user.Username, user.Passwd)
		// user.AuthToken = tokenStr
		if status != nil {
			fmt.Println("Error:", status)
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"status":   http.StatusOK,
				"error":    "",
				"username": user.Username,
				"email":    user.Email,
				// "token":    tokenStr,
			},
		)
	}
}

func LoginUser(c *gin.Context) {
	var user objects.User

	if err := c.BindJSON(&user); err != nil {
		return
	}

	user, status := dbhandler.LoginUser(user)

	if status != nil {
		if errors.Is(status, CustomStatus.UserNotFound) {
			c.JSON(
				http.StatusOK,
				gin.H{"status": http.StatusOK, "error": status},
			)
		} else {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"status": http.StatusBadRequest, "error": status},
			)
		}
	} else {
		// tokenStr, err := GenerateJWTAuthentication(user.Username, user.Passwd)
		// user.AuthToken = tokenStr
		if status != nil {
			fmt.Println("Error:", status)
		}

		c.JSON(
			http.StatusOK,
			gin.H{
				"status":   http.StatusOK,
				"error":    "",
				"username": user.Username,
				"email":    user.Email,
				// "token":    tokenStr,
			},
		)
	}
}

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

func IsAuthorized() {

}
