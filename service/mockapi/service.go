package mockapi

import (
	"emailsync/config"
	"emailsync/model"
	"emailsync/service"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetMapContacts() (*model.MapContacts, *model.ErrorResponse) {
	log.Info("GetMapContacts")

	listContacts, errorResponse := GetListContacts()
	if errorResponse != nil {
		return nil, errorResponse
	}

	retMapContacts := listContacts.ToMapContacts()

	log.Infof("Got [%d] contacts", retMapContacts.Length())

	return retMapContacts, nil
}

func GetListContacts() (*model.ListContacts, *model.ErrorResponse) {
	log.Info("getListContacts")

	var mockAPIListContacts model.MockAPIListContacts
	mockAPIURL := config.GetEnvVariable("MOCK_API_URL")
	contactsEndpoint := config.GetEnvVariable("MOCK_API_CONTACTS_ENDPOINT")

	log.Infof("Getting the contacts from %s ...", mockAPIURL+contactsEndpoint)

	mockAPI := service.NewWithConnection(model.Connection{URL: mockAPIURL})

	response, err := mockAPI.Get(contactsEndpoint)
	if err != nil {
		errorResponse := model.GetErrorResponse(response)
		return nil, errorResponse
	}

	err = json.Unmarshal(response, &mockAPIListContacts)
	if err != nil {
		errorResponse := model.SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error Unmarshal MockAPI List Contacts")
		return nil, errorResponse
	}

	return mockAPIListContacts.ToListContacts(), nil
}
