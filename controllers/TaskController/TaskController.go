package TaskController

import (
	"net/http"
	"postgresql/db"
	"postgresql/models"

	"github.com/gin-gonic/gin"
)

var err interface{}

//tasklist
func List(c *gin.Context) {

	db := db.RUN().MustBegin()

	tasks := []models.Tasks{}

	err = db.Select(&tasks, `SELECT id,title FROM tasks WHERE tid = 0;`)
	// err = db.Select(&tasks, `SELECT id,title FROM tasks;`)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "failed",
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "all list task",
			"data":   tasks,
		})
	}
}
