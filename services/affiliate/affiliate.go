package affiliate

import (
	"e-wallet/api/db/dbaffiliate"
	"e-wallet/api/objects"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAffiliateInfo(c *gin.Context) {
	id := c.Param("id")

	temp, status := dbaffiliate.GetAffiliateByID(id)

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

func GetAllAffiliates(c *gin.Context) {
	affis, status := dbaffiliate.GetAffiliates()

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
				"data":   affis,
			},
		)
	}
}

func AddAffiliate(c *gin.Context) {
	var affi objects.Affiliate

	if err := c.BindJSON(&affi); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": err.Error()},
		)
	}

	affi, status := dbaffiliate.AddAffiliate(affi)

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
				"data":   affi,
			},
		)
	}
}

func DeleteAffiliate(c *gin.Context) {
	id := c.Param("id")
	status := dbaffiliate.DeleteAffiliate(id)

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
