package main

import (
	"github.com/easyPark/backend/src/api"
	"github.com/easyPark/backend/src/config"
)

func main() {
	cfg := config.GetConfig()
	api.InitServer(cfg)
}
