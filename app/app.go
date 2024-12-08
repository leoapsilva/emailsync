package app

import (
	"emailsync/service/server"

	"emailsync/logger"
	log "github.com/sirupsen/logrus"
)

func Start() {
	logger.Init()

	log.Info("Starting web server...")
	api := server.New()
	log.Info("[Success] Web server started.")
	api.LoadDefault().StartLocalAPI()
}
