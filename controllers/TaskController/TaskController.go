package TaskController

import (
	"fmt"
	"net/http"
	"postgresql/db"
	"postgresql/models"

	"github.com/gin-gonic/gin"
)

var err interface{}

//tasklist
func List(c *gin.Context) {

	db := db.RUN().MustBegin()

	task := []models.Task{}

	err = db.Select(&task, `SELECT id,title FROM task WHERE tid = 0;`)
	// err = db.Select(&task, `SELECT id,title FROM tasks;`)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "failed",
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "all list task",
			"data":   task,
		})
	}
}

func Read(c *gin.Context) {
	fmt.Println(c.Param("id"))
}


