package DirectorController

import (
	"net/http"
	"postgresql/db"
	"postgresql/models"

	"github.com/gin-gonic/gin"
)

var err interface{}

func List(c *gin.Context) {

	db := db.RUN().MustBegin()

	directors := []models.Director{}

	err = db.Select(&directors, "SELECT * FROM director;")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "failed",
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "all list director",
		"data":   directors,
	})
}
