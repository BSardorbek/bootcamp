package ProgrammerController

import (
	"net/http"
	"postgresql/db"
	"postgresql/models"

	"github.com/gin-gonic/gin"
)

var err interface{}

func List(c *gin.Context) {

	db := db.RUN().MustBegin()

	programmers := []models.Programmer{}

	err = db.Select(&programmers, "SELECT * FROM programmer;")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "failed",
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "all list programmer",
		"data":   programmers,
	})
}
