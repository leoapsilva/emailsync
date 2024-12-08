package model

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Title  string `json:"title"`
	Status int    `json:"status"`
	Detail string `json:"detail"`
}

func GetErrorResponse(response json.RawMessage) *ErrorResponse {
	e := new(ErrorResponse)
	err := json.Unmarshal(response, e)
	if err != nil {
		e.Detail = err.Error()
		e.Status = http.StatusInternalServerError
		e.Title = "Error Unmarshal Error Response"
		return e
	}
	return e
}

func SetErrorResponse(detail string, status int, title string) *ErrorResponse {
	e := new(ErrorResponse)
	e.Detail = detail
	e.Status = status
	e.Title = title
	return e
}

func (e *ErrorResponse) ToJsonRawMessage() *json.RawMessage {
	response, err := json.Marshal(e)
	if err != nil {
		log.Error(err)
		return nil
	}

	return (*json.RawMessage)(&response)
}
