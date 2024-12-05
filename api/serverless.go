package api

import (
	"emailsync/core"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

var srv http.Handler

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Info("Vercel HTTP Handler")
	server := &core.Server{Echo: echo.New()}

	srv = server.LoadDefault()
	srv.ServeHTTP(w, r)
}
