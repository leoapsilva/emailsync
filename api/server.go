package api

import (
	"emailsync/config"
	"emailsync/controller"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

type API struct {
	*echo.Echo
}

func New() *API {
	return &API{
		Echo: echo.New(),
	}
}

func (api *API) LoadDefault() *API {
	api.Debug = true
	api.HideBanner = true
	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: log.StandardLogger().Out,
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Format: `[${time_custom}] [ECHO ] [api/server.go:0] [id=${id}] [remote_ip=${remote_ip}] ` +
			`[status=${status}] [error=${error}] [protocol=${protocol}] [agent=${user_agent}] ` +
			`[latency=${latency}] [latency_human=${latency_human}] [bytes_in=${bytes_in}] ` +
			`[bytes_out=${bytes_out}] – ${method} ${host}${uri}` + "\n"}))
	api.Use(middleware.Recover())
	api.Use(middleware.CORS())
	api.Server.Addr = ":" + config.ConfigEnvVariable("PORT")
	return api.CreateGroupV1()
}

func (api *API) CreateGroupV1() *API {
	v1 := api.Group("")
	controller.ConfigEndpoints(v1)
	return api
}

func (api *API) StartLocalAPI() {
	log.Infof("Starting service on port [%s]", config.ConfigEnvVariable("PORT"))
	if err := api.Start(api.Server.Addr); err != nil {
		log.Error(err)
	}
}
func (api *API) LoadVercel() *API {
	api.Debug = true
	api.HideBanner = true
	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: log.StandardLogger().Out,
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Format: `[${time_custom}] [ECHO ] [api/server.go:0] [id=${id}] [remote_ip=${remote_ip}] ` +
			`[status=${status}] [error=${error}] [protocol=${protocol}] [agent=${user_agent}] ` +
			`[latency=${latency}] [latency_human=${latency_human}] [bytes_in=${bytes_in}] ` +
			`[bytes_out=${bytes_out}] – ${method} ${host}${uri}` + "\n"}))
	api.Use(middleware.Recover())
	api.Use(middleware.CORS())
	return api.CreateGroupV1()
}

func Handler(w http.ResponseWriter, r *http.Request) {

	log.Info("Starting web server on vercel...")
	api := &API{Echo: echo.New()}
	log.Info("[Success] Web server started on vercel.")
	api.LoadVercel()
	api.ServeHTTP(w, r)

}