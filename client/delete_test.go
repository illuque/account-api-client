package client

import (
	model2 "github.com/illuque/account-api-client/client/model"
	"reflect"
	"testing"
	"time"
)

func TestAccountHttpClient_Delete(t *testing.T) {
	accountHttpClient := NewAccountApiClient("http://localhost:8080/v1/organisation/accounts", 2*time.Second)

	type args struct {
		id      model2.DeleteId
		version int64
	}

	account := buildNewAccount()

	tests := []struct {
		name          string
		args          args
		wantDeleted   bool
		wantErrorData *model2.ErrorData
	}{
		{
			name: "removed correctly when existing for id and versions",
			args: args{
				id: model2.DeleteId{Id: account.ID, Version: *account.Version},
			},
			wantDeleted:   false,
			wantErrorData: nil,
		},
		{
			name: "not found when non existing id",
			args: args{
				id: model2.DeleteId{Id: account.ID, Version: *account.Version},
			},
			wantDeleted:   false,
			wantErrorData: model2.NewNotFound("Specified resource does not exist"),
		},
		{
			name: "conflict for existing id but non existing version",
			args: args{
				id: model2.DeleteId{Id: account.ID, Version: 10},
			},
			wantDeleted:   false,
			wantErrorData: model2.NewConflict("Specified version incorrect"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			switch tt.name {
			case "removed correctly when existing for id and versions", "conflict for existing id but non existing version":
				// TODO:I meter en un seed en vez de ejecutar a mano!
				_, err := accountHttpClient.Create(account)
				if err != nil {
					t.FailNow()
				}
			}

			gotDeleted, gotErrorData := accountHttpClient.Delete(tt.args.id)

			if gotDeleted != tt.wantDeleted {
				t.Errorf("Delete() gotDeleted = %v, want %v", gotDeleted, tt.wantDeleted)
			}

			if !reflect.DeepEqual(gotErrorData, tt.wantErrorData) {
				t.Errorf("Delete() gotErrorData = %v, want %v", gotErrorData, tt.wantErrorData)
				return
			}
		})
	}
}
