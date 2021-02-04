package TeamleadController

import (
	"net/http"
	"postgresql/db"
	"postgresql/models"

	"github.com/gin-gonic/gin"
)

var err interface{}

func List(c *gin.Context) {

	db := db.RUN().MustBegin()

	teamleads := []models.Teamlead{}

	err = db.Select(&teamleads, "SELECT * FROM teamlead;")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "failed",
			"message": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "all list teamlead",
		"data":   teamleads,
	})
}


func TaskAssignP(c *gin.Context) {

}
