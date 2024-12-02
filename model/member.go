package model

import "encoding/json"

type MailchimpMember struct {
	Id                          string          `json:"id"`
	EmailAddress                string          `json:"email_address"`
	UniqueEmailId               string          `json:"unique_email_id"`
	ContactId                   string          `json:"contact_id"`
	FullName                    string          `json:"full_name"`
	WebId                       int             `json:"web_id"`
	EmailType                   string          `json:"email_type"`
	Status                      string          `json:"status"`
	ConsentsToOneToOneMessaging bool            `json:"consents_to_one_to_one_messaging"`
	SmsPhoneNumber              string          `json:"sms_phone_number"`
	SmsSubscriptionStatus       string          `json:"sms_subscription_status"`
	SmsSubscriptionLastUpdated  string          `json:"sms_subscription_last_updated"`
	IpSignup                    string          `json:"ip_signup"`
	TimestampSignup             string          `json:"timestamp_signup"`
	IpOpt                       string          `json:"ip_opt"`
	Language                    string          `json:"language"`
	Vip                         bool            `json:"vip"`
	EmailClient                 string          `json:"email_client"`
	Source                      string          `json:"source"`
	TagsCount                   int             `json:"tags_count"`
	Tags                        []string        `json:"tags"`
	ListId                      string          `json:"list_id"`
	Links                       json.RawMessage `json:"_links"`
	MergeFields                 json.RawMessage `json:"merge_fields"`
	Location                    json.RawMessage `json:"location"`
	Stats                       json.RawMessage `json:"stats"`
}

type MailchimpListMembers struct {
	Members    []MailchimpMember `json:"members"`
	TotalItens int               `json:"total_items"`
	ListId     string            `json:"list_id"`
	Links      json.RawMessage   `json:"_links`
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
