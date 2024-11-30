package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	log.Info("Starting web server on vercel...")
	api := &API{Echo: echo.New()}
	log.Info("[Success] Web server started on vercel.")
	api.LoadDefault().StartAPI()
	api.ServeHTTP(w, r)

}
