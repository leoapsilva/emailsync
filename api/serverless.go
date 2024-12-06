package api

import (
	"emailsync/service/server"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var srv http.Handler

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Info("Vercel HTTP Handler")
	server := server.New()

	srv = server.LoadDefault()
	srv.ServeHTTP(w, r)
}
