package cli

import (
	"emailsync/core"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func Start() {
	log.Info("Starting web server...")
	api := &core.Server{Echo: echo.New()}
	log.Info("[Success] Web server started.")
	api.LoadDefault().StartLocalAPI()

	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.00000",
		DisableSorting:  true,
	})

	log.SetReportCaller(true)
	log.SetLevel(log.InfoLevel)
}