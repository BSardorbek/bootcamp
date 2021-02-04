package teamlead

import (
	teamlead "postgresql/controllers/TeamleadController"

	"github.com/gin-gonic/gin"
)

func TeamleadRoutes(r *gin.RouterGroup) {
	teamleads := r.Group("/teamlead")
	{

		teamleads.GET("/", teamlead.List)
		teamleads.GET("/:id", teamlead.Read)
		teamleads.POST("/:id", teamlead.CheckTask)
		teamleads.GET("/:id/:pid", teamlead.TaskAssignP)

	}
}
