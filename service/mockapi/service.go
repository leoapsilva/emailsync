package mockapi

import (
	"emailsync/config"
	"emailsync/model"
	"emailsync/service"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func GetListContacts() (*model.MapContacts, error) {
	log.Info("GetListContacts")

	var mockAPIListContacts model.MockAPIListContacts
	mockAPIURL := config.GetEnvVariable("MOCK_API_URL")
	contactsEndpoint := config.GetEnvVariable("MOCK_API_CONTACTS_ENDPOINT")

	log.Infof("Getting the contacts from %s ...", mockAPIURL+contactsEndpoint)

	mockAPI := service.NewWithConnection(model.Connection{URL: mockAPIURL})

	response, err := mockAPI.Get(contactsEndpoint, json.RawMessage{})
	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(response, &mockAPIListContacts)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	retMapContacts := mockAPIListContacts.ToMapContacts()

	log.Infof("Got [%d] contacts", retMapContacts.Length())
	log.Debugf("Got [%d] contacts: %v", retMapContacts.Length(), *retMapContacts)

	return retMapContacts, nil
}
