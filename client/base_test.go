package client

import (
	"github.com/google/uuid"
	"github.com/illuque/account-api-client/client/model"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func buildClient() AccountClient {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	baseApiUrl := os.Getenv("ACCOUNT_API_BASE_URL")

	return NewAccountApiClient(baseApiUrl+"/v1/organisation/accounts", 2*time.Second)
}

func buildNewAccount() model.AccountData {
	var countryUK = "UK"
	var version = int64(0)

	return model.AccountData{
		Attributes: &model.AccountAttributes{
			BankID:     "400305",
			BankIDCode: "GBDSC",
			Bic:        "LHVBEE22",
			Country:    &countryUK,
			Name: []string{
				"James Bond",
			},
		},
		ID:             uuid.New().String(),
		OrganisationID: "15a63614-6ae1-4f5b-8f43-7d4dfcb37e76",
		Type:           "accounts",
		Version:        &version,
	}
}
