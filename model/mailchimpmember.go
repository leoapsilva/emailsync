package model

import "encoding/json"

type Address struct {
	Addr1   string `json:"addr1"`
	Addr2   string `json:"addr2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

type MergeFields struct {
	FName    string          `json:"FNAME,omitempty"`
	LName    string          `json:"LNAME,omitempty"`
	Address  json.RawMessage `json:"ADDRESS,omitempty"`
	Phone    string          `json:"PHONE,omitempty"`
	Birthday string          `json:"BIRTHDAY,omitempty"`
	Company  string          `json:"COMPANY,omitempty"`
}

type MailchimpMember struct {
	Id                          string          `json:"id,omitempty"`
	EmailAddress                string          `json:"email_address"`
	UniqueEmailId               string          `json:"unique_email_id,omitempty"`
	ContactId                   string          `json:"contact_id,omitempty"`
	FullName                    string          `json:"full_name,omitempty"`
	WebId                       int             `json:"web_id,omitempty"`
	EmailType                   string          `json:"email_type,omitempty"`
	Status                      string          `json:"status"`
	ConsentsToOneToOneMessaging bool            `json:"consents_to_one_to_one_messaging,omitempty"`
	SmsPhoneNumber              string          `json:"sms_phone_number,omitempty"`
	SmsSubscriptionStatus       string          `json:"sms_subscription_status,omitempty"`
	SmsSubscriptionLastUpdated  string          `json:"sms_subscription_last_updated,omitempty"`
	IpSignup                    string          `json:"ip_signup,omitempty"`
	TimestampSignup             string          `json:"timestamp_signup,omitempty"`
	IpOpt                       string          `json:"ip_opt,omitempty"`
	Language                    string          `json:"language,omitempty"`
	Vip                         bool            `json:"vip,omitempty"`
	EmailClient                 string          `json:"email_client,omitempty"`
	Source                      string          `json:"source,omitempty"`
	TagsCount                   int             `json:"tags_count,omitempty"`
	Tags                        []string        `json:"tags,omitempty"`
	ListId                      string          `json:"list_id,omitempty"`
	MergeFields                 MergeFields     `json:"merge_fields,omitempty"`
	Links                       json.RawMessage `json:"_links,omitempty"`
	Location                    json.RawMessage `json:"location,omitempty"`
	Stats                       json.RawMessage `json:"stats,omitempty"`
}

type MailchimpListMembers struct {
	Members    []MailchimpMember `json:"members"`
	TotalItens int               `json:"total_items"`
	ListId     string            `json:"list_id"`
	Links      json.RawMessage
}

func (l *MailchimpListMembers) ToMapContacts() *MapContacts {

	retMapContacts := make(MapContacts)

	for _, mailchimpMember := range l.Members {
		retMapContacts[mailchimpMember.EmailAddress] = Contact{
			FirstName: "",
			LastName:  "",
			Email:     mailchimpMember.EmailAddress,
		}
	}

	return &retMapContacts
}

func (m *MailchimpMember) ToContact() *Contact {
	contact := &Contact{
		FirstName: m.MergeFields.FName,
		LastName:  m.MergeFields.LName,
		Email:     m.EmailAddress,
	}

	return contact
}
