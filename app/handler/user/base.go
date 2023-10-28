package userhandler

import dbstore "easyparking/aws/dynamodb"

type UserBase struct {
	DBConn *dbstore.DBConn
}

func NewSpaceDB() (UserBase, error) {
	userBase := UserBase{
		DBConn: dbstore.NewDB(),
	}
	return userBase, nil
}
