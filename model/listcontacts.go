package model

type ListContacts []Contact

func (l *ListContacts) ToMapContacts() *MapContacts {

	retMapContacts := make(MapContacts)

	for _, contact := range *l {
		retMapContacts[contact.Email] = contact
	}

	return &retMapContacts
}
