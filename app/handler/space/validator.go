package spacehandler

type GetParkingSpaceMetaDataInput struct {
	Id   string `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
}
type CreateParkingSpaceMetaDataInput struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}
type UpdateParkingSpaceMetaDataInput struct {
	Id      string `form:"id" json:"id"`
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}
