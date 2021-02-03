package teamlead

import (
	teamlead "postgresql/controllers/TeamleadController"

	"github.com/gin-gonic/gin"
)

func TeamleadRoutes(r *gin.RouterGroup) {
	teamleads := r.Group("/teamlead")
	{
		// teamleads.POST("/", create)
		teamleads.GET("/", teamlead.List)
		// teamleads.GET("/:id", read)
		// teamleads.DELETE("/:id", remove)
		// teamleads.PATCH("/:id", update)
	}
}
