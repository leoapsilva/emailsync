package mailchimp

import (
	"emailsync/config"
	"emailsync/model"
	"emailsync/service"
	"encoding/json"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	ListMembersEndpoint = "/lists/@LIST_ID/members"
)

func getListMembers() (*model.MailchimpListMembers, error) {
	log.Info("getListContacts")

	var mailchimpListMembers model.MailchimpListMembers
	APIURL := config.GetEnvVariable("MAILCHIMP_API_URL")
	endpoint := strings.Replace(ListMembersEndpoint, "@LIST_ID", config.GetEnvVariable("MAILCHIMP_LIST_ID"), -1)

	log.Infof("Getting the list members from %s ...", APIURL+endpoint)

	mailchimpAPI := service.NewWithConnection(model.Connection{URL: APIURL})
	mailchimpAPI.SetBasicAuth("user", config.GetEnvVariable("MAILCHIMP_API_KEY"))

	response, err := mailchimpAPI.Get(endpoint, json.RawMessage{})
	log.Info(string(response))

	if err != nil {
		log.Error(err)
		return nil, err
	}

	err = json.Unmarshal(response, &mailchimpListMembers)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &mailchimpListMembers, nil
}

// func addListMember() {

// }

// func deleteListMember() {

// }

func GetMapContacts() (*model.MapContacts, error) {
	log.Info("GetMapContacts")

	mailchimpListMembers, err := getListMembers()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	retMapContacts := mailchimpListMembers.ToMapContacts()

	log.Infof("Got [%d] contacts", retMapContacts.Length())
	log.Debugf("Got [%d] contacts: %v", retMapContacts.Length(), *retMapContacts)

	return retMapContacts, nil
}
