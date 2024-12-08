package model

import (
	"emailsync/config"
	"emailsync/utils"
	"net/http"
)

type Contact struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (contact *Contact) ToMailchimpMember() *MailchimpMember {

	if ok, _ := contact.Validate(); !ok {
		return nil
	}

	member := &MailchimpMember{
		ListId:       config.GetEnvVariable("MAILCHIMP_LIST_ID"),
		EmailAddress: contact.Email,
		FullName:     contact.FirstName + " " + contact.LastName,
		Status:       "subscribed",
		EmailType:    "html",
		MergeFields: MergeFields{
			FName: contact.FirstName,
			LName: contact.LastName,
		},
	}

	return member
}

func (contact *Contact) Validate() (bool, *ErrorResponse) {

	if !utils.EmailValidation(contact.Email) {
		errorResponse := SetErrorResponse("Malformed email address", http.StatusUnprocessableEntity, "Invalid email address")
		return false, errorResponse
	}

	if contact.FirstName == "" {
		errorResponse := SetErrorResponse("First name cannot be empty", http.StatusUnprocessableEntity, "First name empty")
		return false, errorResponse
	}

	if contact.LastName == "" {
		errorResponse := SetErrorResponse("Last name cannot be empty", http.StatusUnprocessableEntity, "Last name empty")
		return false, errorResponse
	}

	return true, nil
}
