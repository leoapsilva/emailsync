package service

import (
	"crypto/tls"
	"emailsync/model"
	"net/http"
	"time"

	resty "github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

type ServiceAPI struct {
	api *resty.Client
	con model.Connection
}

func (p *ServiceAPI) EnableDebug() {
	p.api.SetDebug(true)
	p.api.SetLogger(logrus.StandardLogger())
}

func (p *ServiceAPI) SetDebug(debug bool) {
	p.api.SetDebug(debug)
}

func (p *ServiceAPI) SetTransport(transport *http.Transport) *ServiceAPI {
	p.api.SetTransport(transport)
	return p
}

func (p *ServiceAPI) SetConnection(con model.Connection) *ServiceAPI {
	p.con = con
	return p
}
func (p *ServiceAPI) SetAPI(transport *http.Transport) *ServiceAPI {
	p.api = resty.New()
	p.api = p.api.SetTransport(transport)
	p.api = p.api.SetHeader("Content-Type", "application/json;charset=UTF-8")
	p.EnableDebug()
	return p
}

func NewWithConnection(con model.Connection) *ServiceAPI {
	return newApi().SetConnection(con)
}

func newApi() *ServiceAPI {
	var s ServiceAPI
	s.SetAPI(&http.Transport{
		MaxIdleConnsPerHost: 20,
		IdleConnTimeout:     10 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: false},
	})
	return &s
}
