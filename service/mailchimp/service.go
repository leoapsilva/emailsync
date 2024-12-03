package mailchimp

import (
	"emailsync/config"
	"emailsync/model"
	"emailsync/service"
	"encoding/json"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"
)

const (
	ListMembersEndpoint                       = "/lists/@LIST_ID/members"
	ListMembersPathParametersCount1000        = "count=1000"
	ListMembersPathParametersStatusArchieved  = "status=archived"
	ListMembersPathParametersStatusSubscribed = "status=archived"
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

func getListMembers(pathParameters string) (*model.MailchimpListMembers, *model.ErrorResponse) {
	log.Info("getListContacts")

	var mailchimpListMembers model.MailchimpListMembers
	APIURL := config.GetEnvVariable("MAILCHIMP_API_URL")
	endpoint := strings.Replace(ListMembersEndpoint, "@LIST_ID", config.GetEnvVariable("MAILCHIMP_LIST_ID"), -1)

	log.Infof("Getting the list members from %s ...", APIURL+endpoint+"?"+pathParameters)

	mailchimpAPI := service.NewWithConnection(model.Connection{URL: APIURL})

	APIKey, err := config.DecodeBase64(config.GetEnvVariable("MAILCHIMP_API_KEY"))
	if err != nil {
		errorResponse := SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error decoding Mailchimp API Key")
		return nil, errorResponse
	}

	mailchimpAPI.SetBasicAuth("user", APIKey)

	response, err := mailchimpAPI.Get(endpoint, json.RawMessage{})
	log.Info(string(response))

	if err != nil {
		errorResponse := GetErrorResponse(response)
		return nil, errorResponse
	}

	err = json.Unmarshal(response, &mailchimpListMembers)
	if err != nil {
		errorResponse := SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error Unmarshal Mailchimp List Members")
		return nil, errorResponse
	}

	return &mailchimpListMembers, nil
}

func addListMember(member json.RawMessage) (*model.MailchimpMember, *model.ErrorResponse) {
	log.Info("addListMember")

	var mailchimpMember model.MailchimpMember

	APIURL := config.GetEnvVariable("MAILCHIMP_API_URL")
	endpoint := strings.Replace(ListMembersEndpoint, "@LIST_ID", config.GetEnvVariable("MAILCHIMP_LIST_ID"), -1)

	log.Infof("Adding list member from %s ...", APIURL+endpoint)

	mailchimpAPI := service.NewWithConnection(model.Connection{URL: APIURL})

	APIKey, err := config.DecodeBase64(config.GetEnvVariable("MAILCHIMP_API_KEY"))
	if err != nil {
		errorResponse := SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error decoding Mailchimp API Key")
		return nil, errorResponse
	}

	mailchimpAPI.SetBasicAuth("user", APIKey)

	response, err := mailchimpAPI.Post(endpoint, member)

	if err != nil {
		errorResponse := GetErrorResponse(response)
		return nil, errorResponse
	}

	err = json.Unmarshal(response, &mailchimpMember)
	if err != nil {
		errorResponse := SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error Unmarshal Mailchimp Member")
		return nil, errorResponse
	}

	return &mailchimpMember, nil
}

func archiveMember(memberId string) *model.ErrorResponse {

	APIURL := config.GetEnvVariable("MAILCHIMP_API_URL")
	endpoint := strings.Replace(ListMembersEndpoint, "@LIST_ID", config.GetEnvVariable("MAILCHIMP_LIST_ID"), -1)

	log.Infof("Adding list member from %s ...", APIURL+endpoint+"/"+memberId)

	mailchimpAPI := service.NewWithConnection(model.Connection{URL: APIURL})

	APIKey, err := config.DecodeBase64(config.GetEnvVariable("MAILCHIMP_API_KEY"))
	if err != nil {
		errorResponse := SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error decoding Mailchimp API Key")
		return errorResponse
	}

	mailchimpAPI.SetBasicAuth("user", APIKey)

	response, err := mailchimpAPI.Delete(endpoint, json.RawMessage{})
	log.Info(string(response))

	if err != nil {
		errorResponse := GetErrorResponse(response)
		return errorResponse
	}

	return nil
}

func AddContact(contact *model.Contact) *model.ErrorResponse {
	log.Info("AddContact")

	errorResponse := new(model.ErrorResponse)

	member := &model.MailchimpMember{
		ListId:       config.GetEnvVariable("MAILCHIMP_LIST_ID"),
		EmailAddress: contact.Email,
		FullName:     contact.FirstName + " " + contact.LastName,
		Status:       "subscribed",
		EmailType:    "html",
		MergeFields: model.MergeFields{
			FName: contact.FirstName,
			LName: contact.LastName,
		},
	}

	payload, err := json.Marshal(member)
	if err != nil {
		errorResponse := SetErrorResponse(err.Error(), http.StatusInternalServerError, "Error Unmarshal Mailchimp Member")
		return errorResponse
	}

	log.Infof("Adding contact [%s] added to Mailchimp List.", contact.Email)
	log.Infof("%s", payload)

	_, errorResponse = addListMember(payload)
	if errorResponse != nil {
		return errorResponse
	}

	log.Infof("Contact [%s] added to Mailchimp List.", contact.Email)

	return nil
}

func GetMapContacts() (*model.MapContacts, *model.ErrorResponse) {
	log.Info("GetMapContacts")

	mailchimpListMembers, errorResponse := getListMembers(ListMembersPathParametersCount1000)
	if errorResponse != nil {
		return nil, errorResponse
	}

	retMapContacts := mailchimpListMembers.ToMapContacts()

	log.Infof("Got [%d] contacts", retMapContacts.Length())

	return retMapContacts, nil
}

func ArchiveAllMembers() *model.ErrorResponse {

	mailchimpListMembers, errorResponse := getListMembers(ListMembersPathParametersCount1000)
	if errorResponse != nil {
		return errorResponse
	}

	for _, member := range mailchimpListMembers.Members {
		errorResponse = archiveMember(member.EmailAddress)
		if errorResponse != nil {
			return errorResponse
		}
		log.Infof("Archiving member [%s]", member.EmailAddress)
	}

	return nil
}
