package handler

import (
	"api/e-wallet/db"
	"api/e-wallet/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetTransactions(c *gin.Context) {
	res, err := db.GetTransactions()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't find Any transaction")
		return
	}
	fmt.Println("Get all transaction infor")
	c.IndentedJSON(http.StatusOK, res)
	return
}
func AddTransaction(c *gin.Context) {
	var tran entities.Transaction
	if err := c.BindJSON(&tran); err != nil {
		return
	}
	res, err := db.AddTransaction(tran.SenderName, tran.ReceiverName, tran.Date, tran.Money, tran.Message)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't add Transaction")
		return
	}
	c.IndentedJSON(http.StatusOK, res)

	//res := c.BindJSON("lastName")
	//fmt.Println(res)
	return
}
func DeleteTransaction(c *gin.Context) {
	var id int
	if err := c.BindJSON(&id); err != nil {
		return
	}
	err := db.DeleteTransaction(strconv.Itoa(id))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't find ID")
		return
	}
	c.IndentedJSON(http.StatusOK, id)
	fmt.Println(id)
	return
}
