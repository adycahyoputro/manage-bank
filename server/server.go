package server

import (
	"log"

	"github.com/adycahyoputro/manage-bank/app"
	"github.com/adycahyoputro/manage-bank/config"
	"github.com/adycahyoputro/manage-bank/utils"
)

func Run() error {
	db, err := config.InitDB()
	if err != nil {
		return err
	}
	defer config.CloseDB(db)

	router := app.SetupRouter(db)

	serverAddress := utils.GetEnv("SERVER_ADDRESS")
	log.Printf("server is running on address : %s\n", serverAddress)
	if err := router.Run(serverAddress); err != nil {
		return err
	}
	return nil
}
