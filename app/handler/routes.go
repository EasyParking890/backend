package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var Server *gin.Engine

func init() {
	Server = gin.Default()

	dev := Server.Group("/dev")

	dev.GET("/demo", Demo)
}
func Demo(ctx *gin.Context) {
	fmt.Println("Hello")
}
