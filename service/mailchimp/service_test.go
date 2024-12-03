package mailchimp

import (
	"emailsync/config"
	"emailsync/model"
	"encoding/json"
	"reflect"
	"testing"
)

func Test_addListMember(t *testing.T) {
	config.LoadEnvVariables()

	type args struct {
		member json.RawMessage
	}
	tests := []struct {
		name    string
		args    args
		want    *model.MailchimpMember
		wantErr bool
	}{
		{
			name: "Add a email address twice",
			args: args{
				member: []byte(`{"email_address":"Terrance342@hotmail.com","full_name":"Cole Zieme","email_type":"html","status":"subscribed","list_id":"5c2f756353"}`),
			},
			want: &model.MailchimpMember{
				EmailAddress: "Terrance342@hotmail.com",
				Status:       "subscribed",
				ListId:       "5c2f756353",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := addListMember(tt.args.member)
			if (err != nil) != tt.wantErr {
				t.Errorf("addListMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addListMember() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArchiveAllMembers(t *testing.T) {
	config.LoadEnvVariables()

	tests := []struct {
		name string
		want *model.ErrorResponse
	}{
		{
			name: "Archiving all the contacts",
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ArchiveAllMembers()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArchiveAllMembers() = %v, want %v", got, tt.want)
			}
		})
	}
}
