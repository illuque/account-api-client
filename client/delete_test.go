package client

import (
	"github.com/illuque/account-api-client/model/client_error"
	"testing"
	"time"
)

func TestAccountHttpClient_Delete(t *testing.T) {
	accountHttpClient := NewAccountApiClient("http://localhost:8080/v1/organisation/accounts", 2*time.Second)

	type args struct {
		id      string
		version int64
	}

	account := buildNewAccount()

	tests := []struct {
		name          string
		args          args
		wantDeleted   bool
		wantErrorData *client_error.ErrorData
	}{
		{
			name: "removed correctly when existing for id and versions",
			args: args{
				id:      account.ID,
				version: *account.Version,
			},
			wantDeleted:   true,
			wantErrorData: nil,
		},
		{
			name: "not found when non existing id",
			args: args{
				id:      account.ID,
				version: *account.Version,
			},
			wantDeleted:   false,
			wantErrorData: client_error.NewNotFound("Specified resource does not exist"),
		},
		{
			name: "conflict for existing id but non existing version",
			args: args{
				id:      account.ID,
				version: 10,
			},
			wantDeleted:   false,
			wantErrorData: client_error.NewNotFound("Specified version incorrect"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			switch tt.name {
			case "removed correctly when existing for id and versions":
				// TODO:I meter en un seed en vez de ejecutar a mano!
				_, err := accountHttpClient.Create(account)
				if err != nil {
					t.FailNow()
				}
			}

			gotDeleted, gotErrorData := accountHttpClient.Delete(tt.args.id, tt.args.version)
			if gotDeleted != tt.wantDeleted {
				t.Errorf("Delete() gotDeleted = %v, want %v", gotDeleted, tt.wantDeleted)
			}
			if (gotErrorData != nil) && (gotErrorData.Code != tt.wantErrorData.Code || gotErrorData.Retryable != tt.wantErrorData.Retryable) {
				t.Errorf("Delete() gotErrorData = %v, want %v", gotErrorData, tt.wantErrorData)
				return
			}
		})
	}
}
