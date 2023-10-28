package spacehandler

import (
	"easyparking/internal/common/utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *SpaceBase) GetParkingSpaceMetaData(c *gin.Context) {
	utils.Log.Info().Msg("GetParkingSpaceMetaData")
	var getParkingSpaceMetaDataInput GetParkingSpaceMetaDataInput
	if errShouldBindQuery := c.ShouldBindQuery(&getParkingSpaceMetaDataInput); errShouldBindQuery != nil {
		c.JSON(400, gin.H{"msg": errShouldBindQuery.Error()})
		return
	}
	fmt.Println("op:", getParkingSpaceMetaDataInput)
	c.JSON(200, gin.H{"msg": getParkingSpaceMetaDataInput.Name})
	return

}
func (h *SpaceBase) CreateParkingSpaceMetaData(c *gin.Context) {
	utils.Log.Info().Msg("CreateParkingSpaceMetaData")
	var createParkingSpaceMetaDataInput CreateParkingSpaceMetaDataInput
	if errShouldBindQuery := c.ShouldBindJSON(&createParkingSpaceMetaDataInput); errShouldBindQuery != nil {
		c.JSON(400, gin.H{"msg": errShouldBindQuery.Error()})
		return
	}
	fmt.Println("op:", createParkingSpaceMetaDataInput)

}
func (h *SpaceBase) UpdateParkingSpaceMetaData(c *gin.Context) {
	utils.Log.Info().Msg("UpdateParkingSpaceMetaData")
	var updateParkingSpaceMetaDataInput UpdateParkingSpaceMetaDataInput
	if errShouldBindQuery := c.ShouldBindJSON(&updateParkingSpaceMetaDataInput); errShouldBindQuery != nil {
		c.JSON(400, gin.H{"msg": errShouldBindQuery.Error()})
		return
	}
	fmt.Println("op:", updateParkingSpaceMetaDataInput)
}
