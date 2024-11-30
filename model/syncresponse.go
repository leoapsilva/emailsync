package model

type Contacts struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type SyncResponse struct {
	SyncedContacts int32      `json:"syncedContacts"`
	Contacts       []Contacts `json:"contacts"`
}
