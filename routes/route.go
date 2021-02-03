package route

import (
	v1 "postgresql/routes/v1"

	"github.com/gin-gonic/gin"
)

// API --> /api/
func API(r *gin.Engine) {
	api := r.Group("/api")
	{
		v1.ApplyRoutes(api)
	}
}
