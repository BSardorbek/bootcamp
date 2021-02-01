package main

import (
	"postgresql/db"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	db.RUN().MustExec()

	route.Run(":3000")
}
