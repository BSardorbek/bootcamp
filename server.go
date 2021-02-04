package main

import (
	route "postgresql/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()

	route.API(app)

	app.Run(":5000")
}
