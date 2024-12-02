package model

import (
	"reflect"
	"testing"
)

func TestMapContacts_SetDifference(t *testing.T) {
	type args struct {
		otherMap *MapContacts
	}

	thisMapTest := make(MapContacts)
	otherMapTestEmpty := make(MapContacts)
	otherMapTestPartiallySynced := make(MapContacts)
	setDifferencePartiallySynced := make(MapContacts)

	thisMapTest["Kirk.Fritsch104@icloud.com"] = Contact{
		Email:     "Kirk.Fritsch104@icloud.com",
		FirstName: "Kirk",
		LastName:  "Fritsch",
	}

	thisMapTest["Dillon400@hotmail.com"] = Contact{
		Email:     "Dillon400@hotmail.com",
		FirstName: "Jessika",
		LastName:  "Auer",
	}

	thisMapTest["smith.geo@outlook.com"] = Contact{
		Email:     "smith.geo@outlook.com",
		FirstName: "Geo",
		LastName:  "Smith",
	}

	otherMapTestPartiallySynced["Kirk.Fritsch104@icloud.com"] = Contact{
		Email:     "Kirk.Fritsch104@icloud.com",
		FirstName: "Kirk",
		LastName:  "Fritsch",
	}

	setDifferencePartiallySynced["Dillon400@hotmail.com"] = Contact{
		Email:     "Dillon400@hotmail.com",
		FirstName: "Jessika",
		LastName:  "Auer",
	}

	setDifferencePartiallySynced["smith.geo@outlook.com"] = Contact{
		Email:     "smith.geo@outlook.com",
		FirstName: "Geo",
		LastName:  "Smith",
	}

	tests := []struct {
		name    string
		thisMap *MapContacts
		args    args
		want    *MapContacts
	}{
		{
			name:    "SetDifference==thisMap",
			thisMap: &thisMapTest,
			args: args{
				otherMap: &otherMapTestEmpty,
			},
			want: &thisMapTest,
		},
		{
			name:    "SetDifference=={}",
			thisMap: &thisMapTest,
			args: args{
				otherMap: &thisMapTest,
			},
			want: &otherMapTestEmpty,
		},
		{
			name:    "SetDifference==2_elements",
			thisMap: &thisMapTest,
			args: args{
				otherMap: &otherMapTestPartiallySynced,
			},
			want: &setDifferencePartiallySynced,
		},
		{
			name:    "thisMapEmptySetDifference=={}",
			thisMap: &otherMapTestEmpty,
			args: args{
				otherMap: &setDifferencePartiallySynced,
			},
			want: &otherMapTestEmpty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.thisMap.SetDifference(tt.args.otherMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapContacts.SetDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
