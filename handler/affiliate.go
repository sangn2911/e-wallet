package handler

import (
	"api/e-wallet/db"
	"api/e-wallet/entities"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAffiliates(c *gin.Context) {
	res, err := db.GetAffiliates()
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't find Any Affliates")
		return
	}
	fmt.Println("Get all Affiliates infor")
	c.IndentedJSON(http.StatusOK, res)
	return
}
func AddAffiliate(c *gin.Context) {
	var affiliate entities.Affiliate
	if err := c.BindJSON(&affiliate); err != nil {
		return
	}
	res, err := db.AddAffiliate(affiliate.AffiliateName, affiliate.District, affiliate.Address, affiliate.PhoneNumber, affiliate.Fax, affiliate.Email)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't add affiliates")
		return
	}
	c.IndentedJSON(http.StatusOK, res)

	//res := c.BindJSON("lastName")
	//fmt.Println(res)
	return
}
func DeleteAffiliate(c *gin.Context) {
	var id int
	if err := c.BindJSON(&id); err != nil {
		return
	}
	err := db.DeleteAffiliate(strconv.Itoa(id))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't find ID")
		return
	}
	c.IndentedJSON(http.StatusOK, id)
	fmt.Println(id)
	return
}
func GetAffiliate(c *gin.Context) {
	id := c.Param("id")
	tmp, _ := strconv.Atoi(id)
	res, err := db.GetAffiliateByID(tmp)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "Can't find user")
		return
	}
	fmt.Println("User id", id)
	c.IndentedJSON(http.StatusOK, res)
	return
}
