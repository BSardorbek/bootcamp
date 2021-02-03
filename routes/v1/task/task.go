package task

import (
	task "postgresql/controllers/TaskController"

	"github.com/gin-gonic/gin"
)

//programmer...
func TaskRoutes(r *gin.RouterGroup) {
	tasks := r.Group("/task")
	{
		// teamleads.POST("/", create)
		tasks.GET("/", task.List)
		// teamleads.GET("/:id", read)
		// teamleads.DELETE("/:id", remove)
		// teamleads.PATCH("/:id", update)
	}
}
