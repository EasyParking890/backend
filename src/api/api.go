package api

import (
	"fmt"
	"log"

	"github.com/easyPark/backend/src/api/routers"
	"github.com/easyPark/backend/src/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitServer(cfg *config.Config) {
	r := gin.New()
	r.Use(cors.Default())
	RegisterRoutes(r)
	err := r.Run(fmt.Sprintf(":%s", cfg.Server.ExternalPort))
	if err != nil {
		log.Println("server up and runnig")
	} else {
		fmt.Print(err)
	}
}

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")

	v1 := api.Group("/v1")

	{
		health := v1.Group("/health")
		routers.Heaalth(health)
	}
}
