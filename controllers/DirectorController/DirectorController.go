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

func Read(c *gin.Context) {

	db := db.RUN().MustBegin()

	var director models.Director

	id := c.Param("id")

	err := db.Get(&director, "SELECT * FROM director WHERE id = $1", id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
			"data":   err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "get director by id",
			"data":   director,
		})
	}

}

//TaskAssignT ...
func TaskAssignT(c *gin.Context) {
	// fmt.Println(c.Param("id"), c.Param("tid"))

	taskid := c.Param("id")
	teamleadid := c.Param("tid")

	db := db.RUN().MustBegin()

	data, err := db.Query("UPDATE task SET tid = $1 WHERE id = $2", teamleadid, taskid)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  "failed",
			"message": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "assign task teamlead",
			"data":   data,
		})

	}

}

func FinishTask(c *gin.Context) {

}
