package model

type SyncResponse struct {
	SyncedContacts int       `json:"syncedContacts"`
	Contacts       []Contact `json:"contacts"`
}
