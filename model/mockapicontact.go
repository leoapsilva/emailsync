package model

type MockAPIContact struct {
	CreatedAt string `json:"createdAt"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
	Id        string `json:"id"`
}

type MockAPIListContacts []MockAPIContact

func (l *MockAPIListContacts) ToMapContacts() *MapContacts {

	retMapContacts := make(MapContacts)

	for _, mockAPIContact := range *l {
		retMapContacts[mockAPIContact.Email] = Contact{
			FirstName: mockAPIContact.FirstName,
			LastName:  mockAPIContact.LastName,
			Email:     mockAPIContact.Email,
		}
	}

	return &retMapContacts
}

func (l *MockAPIListContacts) ToListContacts() *ListContacts {

	var retListContacts ListContacts

	for _, mockAPIContact := range *l {
		retListContacts = append(retListContacts,
			Contact{
				FirstName: mockAPIContact.FirstName,
				LastName:  mockAPIContact.LastName,
				Email:     mockAPIContact.Email,
			})
	}

	return &retListContacts
}
