package model

import (
	"net/http"
	"reflect"
	"testing"
)

func TestContact_Validate(t *testing.T) {
	tests := []struct {
		name    string
		contact *Contact
		want    bool
		want1   *ErrorResponse
	}{
		{
			name: "Test Successfull Validate",
			contact: &Contact{
				FirstName: "Adrian",
				LastName:  "Monk",
				Email:     "adrian.monk@yahoo.com",
			},
			want:  true,
			want1: nil,
		},
		{
			name: "Test Error First Name Validate",
			contact: &Contact{
				FirstName: "",
				LastName:  "Stottlemeyer",
				Email:     "leland.stottlemeyer@sanfranciscopolice.org",
			},
			want: false,
			want1: &ErrorResponse{
				Title:  "First name empty",
				Status: http.StatusUnprocessableEntity,
				Detail: "First name cannot be empty",
			},
		},
		{
			name: "Test Error Last Name Validate",
			contact: &Contact{
				FirstName: "Randal",
				LastName:  "",
				Email:     "randal.disher@sanfranciscopolice.org",
			},
			want: false,
			want1: &ErrorResponse{
				Title:  "Last name empty",
				Status: http.StatusUnprocessableEntity,
				Detail: "Last name cannot be empty",
			},
		},
		{
			name: "Test Error Email Validate",
			contact: &Contact{
				FirstName: "Sharona",
				LastName:  "Flamming",
				Email:     "sharona",
			},
			want: false,
			want1: &ErrorResponse{
				Title:  "Invalid email address",
				Status: http.StatusUnprocessableEntity,
				Detail: "Malformed email address",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.contact.Validate()
			if got != tt.want {
				t.Errorf("Contact.Validate() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Contact.Validate() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
