package mailchimp

import (
	"emailsync/config"
	"emailsync/model"
	"emailsync/service"
	"emailsync/utils"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

const (
	ListMembersEndpoint = "/lists/{list_id}/members"
	ListMemberEndpoint  = "/lists/{list_id}/members/{member_id}"
)

func getListMembers() (*model.MailchimpListMembers, *model.ErrorResponse) {
	log.Info("getListContacts")

	var mailchimpListMembers model.MailchimpListMembers
	APIURL := config.GetEnvVariable("MAILCHIMP_API_URL")

	log.Infof("Getting the list members from %s ...", APIURL+ListMembersEndpoint)

	mailchimpAPI := service.NewWithConnection(model.Connection{URL: APIURL})

	APIKey, err := utils.DecodeBase64(config.GetEnvVariable("MAILCHIMP_API_KEY"))
	if err != nil {
		errorResponse := model.SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error decoding Mailchimp API Key")
		return nil, errorResponse
	}

	mailchimpAPI.SetBasicAuth("user", APIKey)

	pathParams := map[string]string{
		"list_id": config.GetEnvVariable("MAILCHIMP_LIST_ID"),
	}

	queryParams := map[string]string{
		"count": "1000",
	}

	response, err := mailchimpAPI.SetQueryParams(queryParams).
		SetPathParams(pathParams).
		Get(ListMembersEndpoint)

	log.Info(string(response))

	if err != nil {
		errorResponse := model.GetErrorResponse(response)
		return nil, errorResponse
	}

	err = json.Unmarshal(response, &mailchimpListMembers)
	if err != nil {
		errorResponse := model.SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error Unmarshal Mailchimp List Members")
		return nil, errorResponse
	}

	return &mailchimpListMembers, nil
}

func addListMember(member *model.MailchimpMember) (*model.MailchimpMember, *model.ErrorResponse) {
	log.Info("addListMember")

	var addedMailchimpMember model.MailchimpMember

	addingMember, err := json.Marshal(member)
	if err != nil {
		errorResponse := model.SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error Unmarshal Mailchimp Member")
		return nil, errorResponse
	}

	APIURL := config.GetEnvVariable("MAILCHIMP_API_URL")

	mailchimpAPI := service.NewWithConnection(model.Connection{URL: APIURL})

	APIKey, err := utils.DecodeBase64(config.GetEnvVariable("MAILCHIMP_API_KEY"))
	if err != nil {
		errorResponse := model.SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error decoding Mailchimp API Key")
		return nil, errorResponse
	}

	mailchimpAPI.SetBasicAuth("user", APIKey)
	pathParams := map[string]string{
		"list_id": config.GetEnvVariable("MAILCHIMP_LIST_ID"),
	}

	response, err := mailchimpAPI.SetPathParams(pathParams).Post(ListMembersEndpoint, addingMember)

	log.Debug(string(response))

	if err != nil {
		errorResponse := model.GetErrorResponse(response)
		return nil, errorResponse
	}

	err = json.Unmarshal(response, &addedMailchimpMember)
	if err != nil {
		errorResponse := model.SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error Unmarshal Mailchimp Member")
		return nil, errorResponse
	}

	return &addedMailchimpMember, nil

}

func archiveMember(memberId string) *model.ErrorResponse {

	APIURL := config.GetEnvVariable("MAILCHIMP_API_URL")

	mailchimpAPI := service.NewWithConnection(model.Connection{URL: APIURL})

	APIKey, err := utils.DecodeBase64(config.GetEnvVariable("MAILCHIMP_API_KEY"))
	if err != nil {
		errorResponse := model.SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error decoding Mailchimp API Key")
		return errorResponse
	}

	mailchimpAPI.SetBasicAuth("user", APIKey)

	pathParams := map[string]string{
		"list_id":   config.GetEnvVariable("MAILCHIMP_LIST_ID"),
		"member_id": memberId,
	}

	response, err := mailchimpAPI.Delete(ListMemberEndpoint, pathParams, map[string]string{})

	log.Debug(string(response))

	if err != nil {
		errorResponse := model.GetErrorResponse(response)
		return errorResponse
	}

	return nil
}

func AddContact(contact *model.Contact) *model.ErrorResponse {
	log.Info("AddContact")

	errorResponse := new(model.ErrorResponse)

	member := contact.ToMailchimpMember()

	log.Infof("Adding contact [%s] added to Mailchimp List.", contact.Email)

	_, errorResponse = addListMember(member)
	if errorResponse != nil {
		return errorResponse
	}

	log.Infof("Contact [%s] added to Mailchimp List.", contact.Email)

	return nil
}

func GetMapContacts() (*model.MapContacts, *model.ErrorResponse) {
	log.Info("GetMapContacts")

	mailchimpListMembers, errorResponse := getListMembers()
	if errorResponse != nil {
		return nil, errorResponse
	}

	retMapContacts := mailchimpListMembers.ToMapContacts()

	log.Infof("Got [%d] contacts", retMapContacts.Length())

	return retMapContacts, nil
}

func ArchiveAllMembers() *model.ErrorResponse {

	mailchimpListMembers, errorResponse := getListMembers()
	if errorResponse != nil {
		return errorResponse
	}

	for _, member := range mailchimpListMembers.Members {
		errorResponse = archiveMember(member.Id)
		if errorResponse != nil {
			return errorResponse
		}
		log.Infof("Archiving member [%s]", member.EmailAddress)
	}

	return nil
}
