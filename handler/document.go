package handler

import (
	"api/e-wallet/db"
	"api/e-wallet/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetDocument(c *gin.Context) {
	id := c.Param("id")
	tmp, _ := strconv.Atoi(id)
	res, err := db.GetDocumentByID(tmp)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't find user")
		return
	}
	var a = []entities.Document{res}
	fmt.Println("Customer id", id)
	c.IndentedJSON(http.StatusOK, a)
	return
}
func AddDocument(c *gin.Context) {
	var doc entities.Document
	if err := c.BindJSON(&doc); err != nil {
		return
	}
	res, err := db.AddDocument(doc.DocType, doc.DocNumber, doc.IssuingAuthority, doc.ExpiryDate, doc.Img)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't add User")
		return
	}
	c.IndentedJSON(http.StatusOK, res)

	//res := c.BindJSON("lastName")
	//fmt.Println(res)
	return
}
