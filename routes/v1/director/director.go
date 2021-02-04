package direcor

import (
	director "postgresql/controllers/DirectorController"

	"github.com/gin-gonic/gin"
)

//DirectorRoutes ...
func DirectorRoutes(r *gin.RouterGroup) {
	directors := r.Group("/director")
	{
		directors.GET("/", director.List)
		directors.GET("/:id", director.Read)
	}
}
