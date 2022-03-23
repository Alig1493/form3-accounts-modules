package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/Alig1493/from3-accounts-modules/models"
	"github.com/google/uuid"
)

const UNIQUE_ID = "4c13df18-615d-4bdc-aec2-ed9a2ac36653"

func GetData(newUUID bool) models.Data {
	var id string
	if newUUID {
		id = uuid.NewString()
	} else {
		id = UNIQUE_ID
	}

	versionInt, conversionError := strconv.ParseInt(GetVersion(), 10, 64)
	if conversionError != nil {
		log.Println("Error converting string to int64")
	}
	userId := id
	organizationId := id
	country := "GB"
	account_classification := "Personal"
	jointAccount := false
	accountMatchingOptOut := false
	accountData := models.Data{
		Data: &models.AccountData{
			Type:           "accounts",
			ID:             userId,
			OrganisationID: organizationId,
			Version:        &versionInt,
			Attributes: &models.AccountAttributes{
				Country:                 &country,
				BaseCurrency:            "GBP",
				BankID:                  "400300",
				BankIDCode:              "GBDSC",
				Bic:                     "NWBKGB22",
				Name:                    []string{"Samantha", "Holder"},
				AlternativeNames:        []string{"Sam", "Holder"},
				AccountClassification:   &account_classification,
				JointAccount:            &jointAccount,
				AccountMatchingOptOut:   &accountMatchingOptOut,
				SecondaryIdentification: "A1B2C3D4",
			},
		},
	}
	return accountData
}

func GetUrl() string {
	return "http://" + os.Getenv("ACCOUNT_SERVICE_NAME") + ":" + os.Getenv("ACCOUNT_SERVICE_PORT")
}

func GetVersion() string {
	return os.Getenv("ACCOUNT_SERVICE_VERSION")
}
