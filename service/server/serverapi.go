package server

import (
	"emailsync/config"
	"emailsync/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

type ServerAPI struct {
	*echo.Echo
}

func New() *ServerAPI {
	return &ServerAPI{
		Echo: echo.New(),
	}
}

func (server *ServerAPI) LoadDefault() *ServerAPI {
	server.Debug = true
	server.HideBanner = true
	log.SetReportCaller(true)
	server.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: log.StandardLogger().Out,
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Format: `[${time_custom}] [ECHO ] [api/server.go:0] [id=${id}] [remote_ip=${remote_ip}] ` +
			`[status=${status}] [error=${error}] [protocol=${protocol}] [agent=${user_agent}] ` +
			`[latency=${latency}] [latency_human=${latency_human}] [bytes_in=${bytes_in}] ` +
			`[bytes_out=${bytes_out}] – ${method} ${host}${uri}` + "\n"}))
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())
	return server.CreateGroupV1()
}

func (server *ServerAPI) CreateGroupV1() *ServerAPI {
	v1 := server.Group("")
	controller.ConfigServerEndpoints(v1)
	return server
}

func (server *ServerAPI) StartLocalAPI() {
	config.LoadEnvVariables()
	log.Infof("Starting service on port [%s]", config.GetEnvVariable("SERVER_PORT"))

	server.Server.Addr = ":" + config.GetEnvVariable("SERVER_PORT")
	if err := server.Start(server.Server.Addr); err != nil {
		log.Error(err)
	}
}
