package mockapi

import (
	"emailsync/config"
	"emailsync/model"
	"emailsync/service"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetErrorResponse(response json.RawMessage) *model.ErrorResponse {
	e := new(model.ErrorResponse)
	err := json.Unmarshal(response, e)
	if err != nil {
		e.Detail = err.Error()
		e.Status = http.StatusInternalServerError
		e.Title = "Error Unmarshal Error Response"
		return e
	}
	return e
}

func SetErrorResponse(detail string, status int, title string) *model.ErrorResponse {
	e := new(model.ErrorResponse)
	e.Detail = detail
	e.Status = status
	e.Title = title
	return e
}

func GetMapContacts() (*model.MapContacts, *model.ErrorResponse) {
	log.Info("GetMapContacts")

	mockAPIListContacts, errorResponse := GetListContacts()
	if errorResponse != nil {
		return nil, errorResponse
	}

	retMapContacts := mockAPIListContacts.ToMapContacts()

	log.Infof("Got [%d] contacts", retMapContacts.Length())
	log.Debugf("Got [%d] contacts: %v", retMapContacts.Length(), *retMapContacts)

	return retMapContacts, nil
}

func GetListContacts() (*model.MockAPIListContacts, *model.ErrorResponse) {
	log.Info("getListContacts")

	var mockAPIListContacts model.MockAPIListContacts
	mockAPIURL := config.GetEnvVariable("MOCK_API_URL")
	contactsEndpoint := config.GetEnvVariable("MOCK_API_CONTACTS_ENDPOINT")

	log.Infof("Getting the contacts from %s ...", mockAPIURL+contactsEndpoint)

	mockAPI := service.NewWithConnection(model.Connection{URL: mockAPIURL})

	response, err := mockAPI.Get(contactsEndpoint)
	if err != nil {
		errorResponse := GetErrorResponse(response)
		return nil, errorResponse
	}

	err = json.Unmarshal(response, &mockAPIListContacts)
	if err != nil {
		errorResponse := SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error Unmarshal MockAPI List Contacts")
		return nil, errorResponse
	}

	return &mockAPIListContacts, nil
}
