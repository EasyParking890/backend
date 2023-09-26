package routers

import (
	"github.com/easyPark/backend/src/api/handlers"
	"github.com/gin-gonic/gin"
)

func Heaalth(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()

	r.GET("/", handler.Health)
}
