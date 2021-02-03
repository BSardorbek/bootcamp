package programmer

import (
	programmer "postgresql/controllers/ProgrammerController"

	"github.com/gin-gonic/gin"
)

//programmer...
func ProgrammerRoutes(r *gin.RouterGroup) {
	programmers := r.Group("/programmer")
	{
		// teamleads.POST("/", create)
		programmers.GET("/", programmer.List)
		// teamleads.GET("/:id", read)
		// teamleads.DELETE("/:id", remove)
		// teamleads.PATCH("/:id", update)
	}
}
