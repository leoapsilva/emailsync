package tests

import (
	"emailsync/app"
	"emailsync/config"
	"emailsync/model"
	"emailsync/service"
	"emailsync/service/mailchimp"
	"encoding/json"
	"time"

	log "github.com/sirupsen/logrus"
)

var responseSyncEndpoint json.RawMessage
var serverPort string
var localAPI *service.ServiceAPI

func SetupClient() {
	config.LoadEnvVariables()
	serverPort = config.GetEnvVariable("SERVER_PORT")
	localAPI = service.NewWithConnection(model.Connection{URL: "http://localhost:" + serverPort})
}

func StartServer() {
	go app.Start()
	time.Sleep(service.Timeout2s)
}

func GetSyncEndpoint() (json.RawMessage, error) {
	return localAPI.Get("/contacts/sync")
}

func ArchiveContacts() {
	var syncResponse model.SyncResponse

	err := json.Unmarshal(responseSyncEndpoint, &syncResponse)
	if err != nil {
		log.Error(err)
		return
	}

	for _, contact := range syncResponse.Contacts {
		errorResponse := mailchimp.ArchiveContact(&contact)
		if errorResponse != nil {
			log.Errorf("[%d] - %s. %s", errorResponse.Status, errorResponse.Title, errorResponse.Detail)
			return
		}
	}
}
