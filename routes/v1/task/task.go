package task

import (
	task "postgresql/controllers/TaskController"

	"github.com/gin-gonic/gin"
)

//TaskRoutes ...
func TaskRoutes(r *gin.RouterGroup) {
	tasks := r.Group("/task")
	{
		tasks.GET("/", task.List)
		tasks.GET("/:id", task.Read)

	}
}
