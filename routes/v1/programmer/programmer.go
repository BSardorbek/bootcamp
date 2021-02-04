package programmer

import (
	programmer "postgresql/controllers/ProgrammerController"

	"github.com/gin-gonic/gin"
)

//ProgrammerRoutes ....
func ProgrammerRoutes(r *gin.RouterGroup) {
	programmers := r.Group("/programmer")
	{
		programmers.GET("/", programmer.List)
		programmers.GET("/:id", programmer.Read)
		programmers.POST("/:id", programmer.EndTask)
	}
}
