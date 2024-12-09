package mailchimp

import (
	"crypto/rand"
	"emailsync/config"
	"emailsync/logger"
	"emailsync/model"
	"math/big"
	"net/http"
	"reflect"
	"testing"
)

func Test_addListMember(t *testing.T) {
	config.LoadEnvVariables()
	logger.Init()

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(999))
	if err != nil {
		t.Errorf("addListMember() error = %v", err)

	}

	emailRandomNumber := randomNumber.Text(10)

	type args struct {
		member *model.MailchimpMember
	}
	tests := []struct {
		name              string
		args              args
		want              *model.MailchimpMember
		wantErr           bool
		wantErrorResponse *model.ErrorResponse
	}{
		{
			name: "Test Successfully Add a email address",
			args: args{
				member: &model.MailchimpMember{
					EmailAddress: "Terrance" + emailRandomNumber + "@hotmail.com",
					Status:       "subscribed",
					ListId:       "5c2f756353",
					FullName:     "Cole Zieme",
					EmailType:    "html",
				},
			},
			want: &model.MailchimpMember{
				EmailAddress: "Terrance" + emailRandomNumber + "@hotmail.com",
				Status:       "subscribed",
				ListId:       "5c2f756353",
			},
			wantErr:           false,
			wantErrorResponse: nil,
		},
		{
			name: "Test Error Add a email address twice",
			args: args{
				member: &model.MailchimpMember{
					EmailAddress: "Terrance" + emailRandomNumber + "@hotmail.com",
					Status:       "subscribed",
					ListId:       "5c2f756353",
					FullName:     "Cole Zieme",
					EmailType:    "html",
				},
			},
			want:    nil,
			wantErr: true,
			wantErrorResponse: &model.ErrorResponse{
				Title:  "Member Exists",
				Status: http.StatusBadRequest,
				Detail: "Terrance" + emailRandomNumber + "@hotmail.com is already a list member. Use PUT to insert or update list members.",
			},
		},
	}

	ArchiveContact(&model.Contact{FirstName: "Cole", LastName: "Zime", Email: "Terrance" + emailRandomNumber + "@hotmail.com"})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("addListMember [%s]", tt.args.member.EmailAddress)
			_, err := addListMember(tt.args.member)
			if err != nil {
				if !reflect.DeepEqual(err, tt.wantErrorResponse) {
					t.Errorf("addListMember() err = %v, wantErrorResponse %v", err, tt.wantErrorResponse)
					return
				}
			}
		})
	}
}

func Test_archiveMember(t *testing.T) {
	config.LoadEnvVariables()
	logger.Init()

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(999))
	if err != nil {
		t.Errorf("archiveMember() error = %v", err)

	}
	emailRandomNumber := randomNumber.Text(10)

	randomNumber2, err := rand.Int(rand.Reader, big.NewInt(999))
	if err != nil {
		t.Errorf("archiveMember() error = %v", err)

	}
	emailRandomNumberNotAdded := randomNumber2.Text(10)

	addListMember(&model.MailchimpMember{
		EmailAddress: "Terrance" + emailRandomNumber + "@hotmail.com",
		Status:       "subscribed",
		ListId:       "5c2f756353",
		FullName:     "Cole Zieme",
		EmailType:    "html"})

	type args struct {
		memberId string
	}
	tests := []struct {
		name string
		args args
		want *model.ErrorResponse
	}{
		{
			name: "Test Successfully Archive a member",
			args: args{
				memberId: "Terrance" + emailRandomNumber + "@hotmail.com",
			},
			want: nil,
		},
		{
			name: "Test Error Archive a member not added",
			args: args{
				memberId: "NeverAddedMember" + emailRandomNumberNotAdded + "@hotmail.com",
			},
			want: &model.ErrorResponse{
				Title:  "Resource Not Found",
				Status: http.StatusNotFound,
				Detail: "The requested resource could not be found."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := archiveMember(tt.args.memberId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("archiveMember() = %v, want %v", got, tt.want)
			}
		})
	}
}
