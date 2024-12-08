package mockapi

import (
	"emailsync/config"
	"emailsync/model"
	"testing"
)

func TestGetMapContacts(t *testing.T) {

	config.LoadEnvVariables()

	retMapContacts := make(model.MapContacts, 24)

	tests := []struct {
		name    string
		want    *model.MapContacts
		wantErr bool
	}{
		{
			name:    "Test Successful GetMapContacts",
			want:    &retMapContacts,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMapContacts()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListContacts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got.Length() == 0 {
				t.Errorf("GetListContacts()\n got %d\nwant %d", got.Length(), retMapContacts.Length())
			}
		})
	}
}

func TestGetListContacts(t *testing.T) {

	config.LoadEnvVariables()

	retListContacts := make(model.ListContacts, 24)

	tests := []struct {
		name    string
		want    *model.ListContacts
		wantErr bool
	}{
		{
			name:    "Test Successful GetListContacts",
			want:    &retListContacts,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetListContacts()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetListContacts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(*got) == 0 {
				t.Errorf("GetListContacts()\n got %d\nwant %d", len(*got), len(retListContacts))
			}
		})
	}
}
