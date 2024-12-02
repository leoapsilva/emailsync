package mockapi

import (
	"emailsync/config"
	"emailsync/model"
	"encoding/json"
	"testing"
)

func TestGetListContacts(t *testing.T) {

	config.LoadEnvVariables()

	var mockAPIListContacts model.MockAPIListContacts
	mockAPIListContactsJSON := json.RawMessage(`[{"createdAt":"2022-02-18T16:32:23.057Z","firstName":"Michelle","lastName":"Gaylord","email":"Kirk.Fritsch104@icloud.com","avatar":"https://cdn.fakercloud.com/avatars/dshster_128.jpg","id":"115"},{"createdAt":"2022-02-18T18:09:28.068Z","firstName":"Deborah","lastName":"Schinner","email":"Corbin.Abshire812@gmail.com","avatar":"https://cdn.fakercloud.com/avatars/spacewood__128.jpg","id":"116"},{"createdAt":"2022-02-18T16:41:12.035Z","firstName":"Jessika","lastName":"Auer","email":"Dillon400@hotmail.com","avatar":"https://cdn.fakercloud.com/avatars/xalionmalik_128.jpg","id":"117"},{"createdAt":"2022-02-18T15:13:59.635Z","firstName":"Geo","lastName":"Schmitt","email":"Cierra_Walsh492@outlook.com","avatar":"https://cdn.fakercloud.com/avatars/sircalebgrove_128.jpg","id":"118"},{"createdAt":"2022-02-18T13:40:05.955Z","firstName":"Floyd","lastName":"Gerlach","email":"Adalberto191@gmail.com","avatar":"https://cdn.fakercloud.com/avatars/dreizle_128.jpg","id":"119"},{"createdAt":"2022-02-18T08:17:35.041Z","firstName":"Hoyt","lastName":"Grady","email":"Elvie.Hagenes257@hotmail.com","avatar":"https://cdn.fakercloud.com/avatars/anaami_128.jpg","id":"120"},{"createdAt":"2022-02-18T01:36:33.505Z","firstName":"Royce","lastName":"Kunze","email":"Dallin_Powlowski305@outlook.com","avatar":"https://cdn.fakercloud.com/avatars/ruehldesign_128.jpg","id":"121"},{"createdAt":"2022-02-18T14:41:41.858Z","firstName":"Eileen","lastName":"Schowalter","email":"Jameson966@hotmail.com","avatar":"https://cdn.fakercloud.com/avatars/brenmurrell_128.jpg","id":"122"},{"createdAt":"2022-02-18T15:02:25.539Z","firstName":"Leonie","lastName":"Strosin","email":"Raheem.DAmore353@hotmail.com","avatar":"https://cdn.fakercloud.com/avatars/mhaligowski_128.jpg","id":"123"},{"createdAt":"2022-02-18T01:22:48.341Z","firstName":"Noemie","lastName":"Gleichner","email":"Cristina352@hotmail.com","avatar":"https://cdn.fakercloud.com/avatars/moscoz_128.jpg","id":"124"},{"createdAt":"2022-02-18T15:56:23.326Z","firstName":"Kaleb","lastName":"Robel","email":"Leonel.Sipes532@hotmail.com","avatar":"https://cdn.fakercloud.com/avatars/lepinski_128.jpg","id":"125"},{"createdAt":"2022-02-18T04:53:58.106Z","firstName":"Harold","lastName":"Mann","email":"Alexys_Aufderhar812@icloud.com","avatar":"https://cdn.fakercloud.com/avatars/AlbertoCococi_128.jpg","id":"126"},{"createdAt":"2022-02-18T16:26:20.492Z","firstName":"Danyka","lastName":"Witting","email":"Jamil864@gmail.com","avatar":"https://cdn.fakercloud.com/avatars/rude_128.jpg","id":"127"},{"createdAt":"2022-02-18T04:13:01.747Z","firstName":"Alvena","lastName":"Marks","email":"Trace.Johnston533@gmail.com","avatar":"https://cdn.fakercloud.com/avatars/itskawsar_128.jpg","id":"128"},{"createdAt":"2022-02-18T04:09:47.743Z","firstName":"Christopher","lastName":"Marquardt","email":"Felipe845@icloud.com","avatar":"https://cdn.fakercloud.com/avatars/kurtinc_128.jpg","id":"129"},{"createdAt":"2022-02-18T20:05:15.394Z","firstName":"Tod","lastName":"Stehr","email":"Reginald_Bechtelar37@outlook.com","avatar":"https://cdn.fakercloud.com/avatars/elliotnolten_128.jpg","id":"130"},{"createdAt":"2022-02-18T17:00:53.930Z","firstName":"Delphia","lastName":"Huels","email":"Myrtis884@outlook.com","avatar":"https://cdn.fakercloud.com/avatars/jayphen_128.jpg","id":"131"},{"createdAt":"2022-02-18T00:34:59.582Z","firstName":"Nels","lastName":"Brakus","email":"Rebeka_Thompson462@gmail.com","avatar":"https://cdn.fakercloud.com/avatars/davidmerrique_128.jpg","id":"132"},{"createdAt":"2022-02-18T03:47:26.577Z","firstName":"Bethel","lastName":"Rau","email":"Benedict_Kunze242@hotmail.com","avatar":"https://cdn.fakercloud.com/avatars/janpalounek_128.jpg","id":"133"},{"createdAt":"2022-02-18T16:56:51.886Z","firstName":"Charley","lastName":"Hermann","email":"Jennie.Kessler158@gmail.com","avatar":"https://cdn.fakercloud.com/avatars/0therplanet_128.jpg","id":"134"},{"createdAt":"2022-02-18T05:17:55.961Z","firstName":"Cole","lastName":"Zieme","email":"Terrance257@icloud.com","avatar":"https://cdn.fakercloud.com/avatars/madshensel_128.jpg","id":"135"},{"createdAt":"2022-02-18T09:21:05.918Z","firstName":"Lee","lastName":"Runolfsson","email":"Jackie941@hotmail.com","avatar":"https://cdn.fakercloud.com/avatars/deviljho__128.jpg","id":"136"},{"createdAt":"2022-02-18T18:50:12.509Z","firstName":"Zander","lastName":"Greenholt","email":"Austin903@yahoo.com","avatar":"https://cdn.fakercloud.com/avatars/jm_denis_128.jpg","id":"137"},{"createdAt":"2022-02-18T11:09:33.659Z","firstName":"Rene","lastName":"Nienow","email":"Michael_Dach990@yahoo.com","avatar":"https://cdn.fakercloud.com/avatars/jacobbennett_128.jpg","id":"138"}]`)

	err := json.Unmarshal(mockAPIListContactsJSON, &mockAPIListContacts)
	if err != nil {
		t.Errorf("GetListContacts() error = %v", err)
	}

	retMapContacts := mockAPIListContacts.ToMapContacts()

	tests := []struct {
		name    string
		want    *model.MapContacts
		wantErr bool
	}{
		{
			name:    "GetListContactsMockAPI",
			want:    retMapContacts,
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
