package mailchimp

import (
	"emailsync/config"
	"emailsync/model"
	"reflect"
	"testing"
)

func Test_addListMember(t *testing.T) {
	config.LoadEnvVariables()

	type args struct {
		member *model.MailchimpMember
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
				member: &model.MailchimpMember{
					EmailAddress: "Terrance342@hotmail.com",
					Status:       "subscribed",
					ListId:       "5c2f756353",
					FullName:     "Cole Zieme",
					EmailType:    "html",
				},
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
