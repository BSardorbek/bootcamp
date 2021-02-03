package v1

import (
	director "postgresql/routes/v1/director"
	programmer "postgresql/routes/v1/programmer"
	task "postgresql/routes/v1/task"
	teamlead "postgresql/routes/v1/teamlead"

	"github.com/gin-gonic/gin"
)

// API --> /api/v1/
func ApplyRoutes(r *gin.RouterGroup) {
	api := r.Group("/v1")
	{
		// API --> /api/v1/director
		director.DirectorRoutes(api)
		// API --> /api/v1/teamlead
		teamlead.TeamleadRoutes(api)
		// API --> /api/v1/programmer
		programmer.ProgrammerRoutes(api)
		// API --> /api/v1/task
		task.TaskRoutes(api)
	}
}
