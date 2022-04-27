package transaction

import (
	dbtran "e-wallet/api/db/dbtransaction"
	"e-wallet/api/objects"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context) {
	trans, status := dbtran.GetAllTransactions()

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
				"data":   trans,
			},
		)
	}
}

func AddTransaction(c *gin.Context) {
	var tran objects.Transaction

	if err := c.BindJSON(&tran); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": err.Error()},
		)
	}

	tran, status := dbtran.AddTransaction(tran)

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
				"data":   tran,
			},
		)
	}
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	t, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": err.Error()},
		)
	}
	status := dbtran.DeleteTransaction(t)

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
