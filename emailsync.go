package main

import (
	"emailsync/core"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Starting web server...")
	api := &core.API{Echo: echo.New()}
	log.Info("[Success] Web server started.")
	api.LoadDefault().StartAPI()
}
