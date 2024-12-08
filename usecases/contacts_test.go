package usecases

import (
	"crypto/rand"
	"emailsync/config"
	"emailsync/logger"
	"emailsync/model"
	"emailsync/service/mailchimp"
	"math/big"
	"testing"
)

func TestAddContacts(t *testing.T) {
	config.LoadEnvVariables()
	logger.Init()

	randomNumber, err := rand.Int(rand.Reader, big.NewInt(999))
	if err != nil {
		t.Errorf("archiveMember() error = %v", err)

	}
	emailRandomNumber := randomNumber.Text(10)

	type args struct {
		listContacts *model.ListContacts
	}

	inListContacts := &model.ListContacts{
		{
			FirstName: "Taylor",
			LastName:  "Swift",
			Email:     "taylor.swift" + emailRandomNumber + "@gmail.com",
		},
		{
			FirstName: "Tom",
			LastName:  "Jobim",
			Email:     "tom.jobim" + emailRandomNumber + "@yahoo.com",
		},
		{
			FirstName: "Joseph",
			LastName:  "Climber",
			Email:     "joseph.climber" + emailRandomNumber + "@aol.com",
		},
		{
			FirstName: "Jeffrey",
			LastName:  "Combs",
			Email:     "jeffrey.combs" + emailRandomNumber + "@hotmail.com",
		},
	}

	outListContacts := inListContacts

	tests := []struct {
		name string
		args args
		want *model.SyncResponse
	}{
		{
			name: "Test Successful AddContacts Once",
			args: args{
				listContacts: inListContacts,
			},
			want: &model.SyncResponse{
				SyncedContacts: 4,
				Contacts:       *outListContacts,
			},
		},
		{
			name: "Test Unsuccessfully AddContacts Twice",
			args: args{
				listContacts: inListContacts,
			},
			want: &model.SyncResponse{
				SyncedContacts: 0,
				Contacts:       make(model.ListContacts, 0),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AddContacts(tt.args.listContacts)
			if got.SyncedContacts != tt.want.SyncedContacts {
				t.Errorf("AddContacts() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, contact := range *inListContacts {
		mailchimp.ArchiveContact(&model.Contact{
			FirstName: contact.FirstName,
			LastName:  contact.LastName,
			Email:     contact.Email})
	}
}
