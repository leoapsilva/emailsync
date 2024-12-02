package model

type MapContacts map[string]Contact

func (thisMap *MapContacts) SetDifference(otherMap *MapContacts) *MapContacts {

	setDifference := make(MapContacts)

	for _, contact := range *thisMap {
		_, exists := (*otherMap)[contact.Email]
		if !exists {
			setDifference[contact.Email] = contact
		}
	}

	return &setDifference
}

func (m *MapContacts) Length() int {
	return len(map[string]Contact(*m))
}
