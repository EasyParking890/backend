package handler

import (
	spacehandler "easyparking/app/handler/space"
	userhandler "easyparking/app/handler/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

var Server *gin.Engine

var userBase = &userhandler.UserBase{}
var spaceBase = &spacehandler.SpaceBase{}

func init() {
	spaceBase = spacehandler.NewSpaceDB()

	Server = gin.Default()

	stage := Server.Group("/dev")

	user := stage.Group("/user")
	user.GET("/list", Demo)
	user.GET("", Demo)

	space := stage.Group("/space")
	space.GET("", spaceBase.GetParkingSpaceMetaData)
	space.POST("", spaceBase.CreateParkingSpaceMetaData)
	space.PUT("", spaceBase.UpdateParkingSpaceMetaData)

}
func Demo(ctx *gin.Context) {
	fmt.Println("Hello")
}
