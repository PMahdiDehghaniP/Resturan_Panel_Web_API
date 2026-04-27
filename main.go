package main

import (
	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/api"
	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/config"
	"github.com/PMahdiDehghaniP/Resturan_Panel_Web_API/data/db"
)

func main() {
	appConfig := config.GetConfig()

	db.InitPostgresDB(appConfig)
	defer db.ClosePostgresDB()

	api.InitApiServer()
}
