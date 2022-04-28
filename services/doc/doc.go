package doc

import (
	dbdocument "e-wallet/api/db/dbdocument"
	"e-wallet/api/objects"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDocumentsOfUser(c *gin.Context) {
	var id int
	if err := c.BindJSON(&id); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"status": http.StatusBadRequest, "error": err.Error()},
		)
	}

	docs, status := dbdocument.GetDocumentsOfUser(id)
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
				"data":   docs,
				// "token":    tokenStr,
			},
		)
	}
}

func AddDocumentInfo(c *gin.Context) {
	var doc objects.Document
	if err := c.BindJSON(&doc); err != nil {
		return
	}

	doc, status := dbdocument.AddDocument(doc)
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
				"data":   doc,
			},
		)
	}
}
