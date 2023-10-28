package spacehandler

import dbstore "easyparking/aws/dynamodb"

type SpaceBase struct {
	DBConn *dbstore.DBConn
}

func NewSpaceDB() *SpaceBase {
	spaceBase := &SpaceBase{
		DBConn: dbstore.NewDB(),
	}
	return spaceBase
}
