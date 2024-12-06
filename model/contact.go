package model

import "emailsync/config"

type Contact struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (contact *Contact) ToMailchimpMember() *MailchimpMember {

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
