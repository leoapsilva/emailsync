package main

import (
	"emailsync/api"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Starting web server...")
	api := &api.API{Echo: echo.New()}
	log.Info("[Success] Web server started.")
	api.LoadDefault().StartLocalAPI()
}
