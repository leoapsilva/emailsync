package api

import (
	"emailsync/core"
	"net/http"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

var srv http.Handler

// func CreateGroupV1(server *echo.Echo) *echo.Echo {
// 	v1 := server.Group("")
// 	controller.ConfigServerEndpoints(v1)
// 	return server
// }

// func LoadVercel(server *echo.Echo) *echo.Echo {
// 	server.Debug = true
// 	server.HideBanner = true
// 	log.SetReportCaller(true)
// 	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: log.StandardLogger().Out,
// 		CustomTimeFormat: "2006-01-02 15:04:05.00000",
// 		Format: `[${time_custom}] [ECHO ] [api/server.go:0] [id=${id}] [remote_ip=${remote_ip}] ` +
// 			`[status=${status}] [error=${error}] [protocol=${protocol}] [agent=${user_agent}] ` +
// 			`[latency=${latency}] [latency_human=${latency_human}] [bytes_in=${bytes_in}] ` +
// 			`[bytes_out=${bytes_out}] – ${method} ${host}${uri}` + "\n"}))
// 	server.Use(middleware.Recover())
// 	server.Use(middleware.CORS())
// 	return CreateGroupV1(server)
// }

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Info("Vercel HTTP Handler")
	server := &core.Server{Echo: echo.New()}
	// config.LoadEnvVariables()

	srv = server.LoadDefault()
	srv.ServeHTTP(w, r)
}
