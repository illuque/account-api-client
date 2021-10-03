package client

import (
	"github.com/illuque/account-api-client/client/model"
	"reflect"
	"testing"
)

func TestAccountHttpClient_Delete(t *testing.T) {
	accountHttpClient := buildClient()

	type args struct {
		id      model.DeleteId
		version int64
	}

	account := buildNewAccount()

	tests := []struct {
		name        string
		args        args
		wantDeleted bool
		wantErr     error
	}{
		{
			name: "removed correctly when existing for id and versions",
			args: args{
				id: model.DeleteId{Id: account.ID, Version: *account.Version},
			},
			wantDeleted: true,
			wantErr:     nil,
		},
		{
			name: "not found when non existing id",
			args: args{
				id: model.DeleteId{Id: account.ID, Version: *account.Version},
			},
			wantDeleted: false,
			wantErr:     model.NewNotFound("Specified resource does not exist"),
		},
		{
			name: "conflict for existing id but non existing version",
			args: args{
				id: model.DeleteId{Id: account.ID, Version: 10},
			},
			wantDeleted: false,
			wantErr:     model.NewConflict("Specified version incorrect"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			switch tt.name {
			case "removed correctly when existing for id and versions", "conflict for existing id but non existing version":
				// create one to ensure existing
				_, err := accountHttpClient.Create(account)
				if err != nil {
					t.FailNow()
				}
			}

			gotDeleted, gotErr := accountHttpClient.Delete(tt.args.id)

			if gotDeleted != tt.wantDeleted {
				t.Errorf("Delete() gotDeleted = %v, want %v", gotDeleted, tt.wantDeleted)
			}

			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("Delete() gotErr = %v, want %v", gotErr, tt.wantErr)
				return
			}
		})
	}
}
