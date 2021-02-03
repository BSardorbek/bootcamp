package direcor

import (
	director "postgresql/controllers/DirectorController"

	"github.com/gin-gonic/gin"
)

//director...
func DirectorRoutes(r *gin.RouterGroup) {
	directors := r.Group("/director")
	{
		// directors.POST("/", create)
		directors.GET("/", director.List)
		// directors.GET("/:id", read)
		// directors.DELETE("/:id", remove)
		// directors.PATCH("/:id", update)
	}
}
