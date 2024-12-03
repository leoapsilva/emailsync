package model

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMailchimpListMembers_ToMapContacts(t *testing.T) {

	memberJSON := json.RawMessage(`{"id":"733ab0f42721fccf3aa9144b38639872","email_address":"leoapsilva@gmail.com","unique_email_id":"ed2d63b846","contact_id":"328bf3b2a26ec04019521b9e4b6a40fe","full_name":"Leonardo","web_id":608421101,"email_type":"html","status":"subscribed","consents_to_one_to_one_messaging":true,"sms_phone_number":"","sms_subscription_status":"","sms_subscription_last_updated":"","merge_fields":{"FNAME":"Leonardo","LNAME":"Silva","ADDRESS":{"addr1":"Trio\\nRua Itanhaem\\nSao Jose Do Rio Preto, SP 15050-457\\nBrazil","addr2":"","city":"","state":"","zip":"","country":"US"},"PHONE":"","BIRTHDAY":"","COMPANY":""},"stats":{"avg_open_rate":0,"avg_click_rate":0},"ip_signup":"","timestamp_signup":"","ip_opt":"177.21.136.50","timestamp_opt":"2024-12-01T02:20:10+00:00","member_rating":2,"last_changed":"2024-12-01T02:20:10+00:00","language":"","vip":false,"email_client":"","location":{"latitude":0,"longitude":0,"gmtoff":0,"dstoff":0,"country_code":"","timezone":"","region":""},"source":"Admin Add","tags_count":0,"tags":[],"list_id":"5c2f756353","_links":[{"rel":"self","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/733ab0f42721fccf3aa9144b38639872","method":"GET","targetSchema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/Response.json"},{"rel":"parent","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members","method":"GET","targetSchema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/CollectionResponse.json","schema":"https://us9.api.mailchimp.com/schema/3.0/Paths/Lists/Members/Collection.json"},{"rel":"update","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/733ab0f42721fccf3aa9144b38639872","method":"PATCH","targetSchema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/Response.json","schema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/PATCH.json"},{"rel":"upsert","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/733ab0f42721fccf3aa9144b38639872","method":"PUT","targetSchema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/Response.json","schema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/PUT.json"},{"rel":"delete","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/733ab0f42721fccf3aa9144b38639872","method":"DELETE"},{"rel":"activity","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/733ab0f42721fccf3aa9144b38639872/activity","method":"GET","targetSchema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/Activity/Response.json"},{"rel":"goals","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/733ab0f42721fccf3aa9144b38639872/goals","method":"GET","targetSchema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/Goals/Response.json"},{"rel":"notes","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/733ab0f42721fccf3aa9144b38639872/notes","method":"GET","targetSchema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/Notes/CollectionResponse.json"},{"rel":"events","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/733ab0f42721fccf3aa9144b38639872/events","method":"POST","targetSchema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/Events/POST.json"},{"rel":"delete_permanent","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/733ab0f42721fccf3aa9144b38639872/actions/delete-permanent","method":"POST"}]}`)
	member := new(MailchimpMember)

	err := json.Unmarshal(memberJSON, member)
	if err != nil {
		t.Errorf("MailchimpMembers Unmarshal = %v", err)
	}

	memberJSON2 := json.RawMessage(`{"id":"d8a530da21eadd6a6c60cc6891afb50a","email_address":"Leonel.Sipes800@icloud.com","unique_email_id":"87ab98e26d","contact_id":"05bba16dce00e8f9713b3e0d4739b02e","full_name":"","web_id":608421133,"email_type":"html","status":"pending","consents_to_one_to_one_messaging":true,"sms_phone_number":"","sms_subscription_status":"","sms_subscription_last_updated":"","merge_fields":{"FNAME":"","LNAME":"","ADDRESS":"","PHONE":"","BIRTHDAY":"","COMPANY":""},"stats":{"avg_open_rate":0,"avg_click_rate":0},"ip_signup":"177.21.136.50","timestamp_signup":"2024-12-02T14:06:57+00:00","ip_opt":"","timestamp_opt":"","member_rating":2,"last_changed":"2024-12-02T14:06:57+00:00","language":"","vip":false,"email_client":"","location":{"latitude":0,"longitude":0,"gmtoff":0,"dstoff":0,"country_code":"","timezone":"","region":""},"source":"API - Generic","tags_count":0,"tags":[],"list_id":"5c2f756353","_links":[{"rel":"self","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/d8a530da21eadd6a6c60cc6891afb50a","method":"GET","targetSchema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/Response.json"},{"rel":"parent","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members","method":"GET","targetSchema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/CollectionResponse.json","schema":"https://us9.api.mailchimp.com/schema/3.0/Paths/Lists/Members/Collection.json"},{"rel":"update","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/d8a530da21eadd6a6c60cc6891afb50a","method":"PATCH","targetSchema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/Response.json","schema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/PATCH.json"},{"rel":"upsert","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/d8a530da21eadd6a6c60cc6891afb50a","method":"PUT","targetSchema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/Response.json","schema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/PUT.json"},{"rel":"delete","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/d8a530da21eadd6a6c60cc6891afb50a","method":"DELETE"},{"rel":"activity","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/d8a530da21eadd6a6c60cc6891afb50a/activity","method":"GET","targetSchema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/Activity/Response.json"},{"rel":"goals","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/d8a530da21eadd6a6c60cc6891afb50a/goals","method":"GET","targetSchema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/Goals/Response.json"},{"rel":"notes","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/d8a530da21eadd6a6c60cc6891afb50a/notes","method":"GET","targetSchema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/Notes/CollectionResponse.json"},{"rel":"events","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/d8a530da21eadd6a6c60cc6891afb50a/events","method":"POST","targetSchema":"https://us9.api.mailchimp.com/schema/3.0/Definitions/Lists/Members/Events/POST.json"},{"rel":"delete_permanent","href":"https://us9.api.mailchimp.com/3.0/lists/5c2f756353/members/d8a530da21eadd6a6c60cc6891afb50a/actions/delete-permanent","method":"POST"}]}`)

	err = json.Unmarshal(memberJSON2, member)
	if err != nil {
		t.Errorf("MailchimpMembers Unmarshal = %v", err)
	}

	tests := []struct {
		name string
		l    *MailchimpListMembers
		want *MapContacts
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.ToMapContacts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MailchimpListMembers.ToMapContacts() = %v, want %v", got, tt.want)
			}
		})
	}
}
