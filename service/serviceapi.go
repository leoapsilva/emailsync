package service

import (
	"context"
	"crypto/tls"
	"emailsync/model"
	"encoding/json"
	"errors"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"

	resty "github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

const (
	Timeout20ms  time.Duration = 20 * time.Millisecond
	Timeout50ms                = 50 * time.Millisecond
	Timeout200ms               = 200 * time.Millisecond
	Timeout500ms               = 500 * time.Millisecond
	Timeout600ms               = 600 * time.Millisecond
	Timeout1s                  = 1 * time.Second
	Timeout2s                  = 2 * time.Second
	Timeout5s                  = 5 * time.Second
	Timeout10s                 = 10 * time.Second
	Timeout20s                 = 20 * time.Second
	Timeout30s                 = 30 * time.Second
	Timeout60s                 = 60 * time.Second
	Timeout2M                  = 2 * time.Minute
	Timeout3M                  = 3 * time.Minute
)

type ServiceAPI struct {
	api *resty.Client
	con model.Connection
}

func (p *ServiceAPI) EnableDebug() {
	p.api.SetDebug(true)
	p.api.SetLogger(log.StandardLogger())
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

func DefaultRetryCondition(r *resty.Response, err error) bool {
	return false
}

func (p *ServiceAPI) SetAPI(transport *http.Transport) *ServiceAPI {
	p.api = resty.New()
	p.api = p.api.SetTransport(transport)
	p.api = p.api.SetHeader("Content-Type", "application/json;charset=UTF-8")
	p.EnableDebug()
	return p
}

func (p *ServiceAPI) SetQueryParam(key string, value string) *ServiceAPI {
	p.api.SetQueryParam(key, value)
	return p
}

func (p *ServiceAPI) SetQueryParams(params map[string]string) *ServiceAPI {
	p.api.SetQueryParams(params)
	return p
}

func (p *ServiceAPI) SetPathParam(key string, value string) *ServiceAPI {
	p.api.SetPathParam(key, value)
	return p
}

func (p *ServiceAPI) SetPathParams(params map[string]string) *ServiceAPI {
	p.api.SetPathParams(params)
	return p
}

func (p *ServiceAPI) SetBasicAuth(user string, password string) *ServiceAPI {
	p.api.SetBasicAuth(user, password)
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

func (p *ServiceAPI) Post(endpoint string, payload json.RawMessage) (result json.RawMessage, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), Timeout2M)
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered. Error:\n", r)
			debug.PrintStack()
		}
		cancel()
	}()

	p.api.RetryConditions = nil
	p.api.SetRetryCount(3).SetRetryWaitTime(Timeout2s).SetRetryMaxWaitTime(Timeout10s).AddRetryCondition(DefaultRetryCondition)

	resp, err := p.api.R().SetHeader("Content-Type", "application/json;charset=UTF-8").
		SetBody(payload).
		SetError(&err).
		SetResult(&result).
		SetContext(ctx).
		Post(p.con.FormatURL(endpoint))

	return handleResponse(resp, err)
}

func (p *ServiceAPI) Get(endpoint string) (retorno json.RawMessage, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), Timeout2M)
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered. Error:\n", r)
			debug.PrintStack()
		}
		cancel()
	}()

	p.api.RetryConditions = nil
	p.api.SetRetryCount(3).SetRetryWaitTime(Timeout2s).SetRetryMaxWaitTime(Timeout10s).AddRetryCondition(DefaultRetryCondition)

	resp, err := p.api.R().SetHeader("Content-Type", "application/json;charset=UTF-8").
		SetError(&err).
		SetResult(&retorno).
		SetContext(ctx).
		Get(p.con.FormatURL(endpoint))

	return handleResponse(resp, err)
}

func (p *ServiceAPI) Put(endpoint string, payload json.RawMessage, pathParams map[string]string, queryParams map[string]string) (result json.RawMessage, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), Timeout2M)
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered. Error:\n", r)
			debug.PrintStack()
		}
		cancel()
	}()

	p.api.RetryConditions = nil
	p.api.SetRetryCount(3).SetRetryWaitTime(Timeout2s).SetRetryMaxWaitTime(Timeout10s).AddRetryCondition(DefaultRetryCondition)

	resp, err := p.api.R().SetHeader("Content-Type", "application/json;charset=UTF-8").
		SetQueryParams(queryParams).
		SetPathParams(pathParams).
		SetBody(payload).
		SetError(&err).
		SetResult(&result).
		SetContext(ctx).
		Put(p.con.FormatURL(endpoint))

	return handleResponse(resp, err)
}

func (p *ServiceAPI) Delete(endpoint string, pathParams map[string]string, queryParams map[string]string) (result json.RawMessage, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), Timeout2M)
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered. Error:\n", r)
			debug.PrintStack()
		}
		cancel()
	}()

	p.api.RetryConditions = nil
	p.api.SetRetryCount(3).SetRetryWaitTime(Timeout2s).SetRetryMaxWaitTime(Timeout10s).AddRetryCondition(DefaultRetryCondition)

	resp, err := p.api.R().SetHeader("Content-Type", "application/json;charset=UTF-8").
		SetQueryParams(queryParams).
		SetPathParams(pathParams).
		SetError(&err).
		SetResult(&result).
		SetContext(ctx).
		Delete(p.con.FormatURL(endpoint))

	return handleResponse(resp, err)
}

func handleResponse(resp *resty.Response, err error) (json.RawMessage, error) {
	if err != nil {
		if resp != nil {
			errorResponse := model.GetErrorResponse(resp.Body())
			if errorResponse == nil {
				log.Error("Error on " + resp.Request.Method + " [" + err.Error() + "]")
				errorResponse = &model.ErrorResponse{
					Title:  err.Error(),
					Status: http.StatusBadRequest,
					Detail: resp.String(),
				}
			}
			return *errorResponse.ToJsonRawMessage(), errors.New(err.Error())
		} else {
			log.Error("Error [" + err.Error() + "]")
			errorResponse := &model.ErrorResponse{
				Title:  err.Error(),
				Status: http.StatusBadRequest,
				Detail: err.Error(),
			}
			return *errorResponse.ToJsonRawMessage(), errors.New(err.Error())
		}
	} else {
		statusCode := resp.StatusCode()
		if statusCode >= http.StatusMultipleChoices {
			errorMessage := "Status Code: " + strconv.Itoa(resp.StatusCode()) + ", Status: " + resp.Status()
			errorResponse := model.GetErrorResponse(resp.Body())
			if errorResponse == nil {
				errorResponse = &model.ErrorResponse{
					Title:  resp.Status(),
					Status: statusCode,
					Detail: resp.String(),
				}
				return *errorResponse.ToJsonRawMessage(), errors.New(errorMessage)
			}
			return resp.Body(), errors.New(errorMessage)
		}
	}

	return resp.Body(), nil
}
