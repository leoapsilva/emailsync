package model

type SyncResponse struct {
	SyncedContacts int32     `json:"syncedContacts"`
	Contacts       []Contact `json:"contacts"`
}
