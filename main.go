package main

import (
	"log"

	"github.com/Alig1493/from3-accounts-modules/models"
	"github.com/Alig1493/from3-accounts-modules/repositories"
	"github.com/Alig1493/from3-accounts-modules/services"
	"github.com/google/uuid"
)

var (
	accountRepository = repositories.NewAccountRepository()
	service           = services.NewAccountService(accountRepository)
	userId            = uuid.NewString()
	organizationId    = uuid.NewString()
)

func displayResponse(responseData *models.ResponseData, responseError error) {
	if responseError != nil {
		log.Fatalf("Response error: %v", &models.ErrorData{ErrorMessage: responseError.Error()})
	} else {
		log.Printf("Successful response for data: %+v", *responseData)
	}
	log.Printf("Status code: %v", responseData.StatusCode)
}

func sampleCreate() {
	// {
	// 	'data': {
	// 	  'type': 'accounts',
	// 	  'id': 'ad27e265-9605-4b4b-a0e5-3003ea9cc4dc',
	// 	  'organisation_id': 'eb0bd6f5-c3f5-44b2-b677-acd23cdde73c',
	// 	  'attributes': {
	// 		'country': 'GB',
	// 		'base_currency': 'GBP',
	// 		'bank_id': '400300',
	// 		'bank_id_code': 'GBDSC',
	// 		'bic': 'NWBKGB22',
	// 		'name': [
	// 		  'Samantha Holder'
	// 		],
	// 		'alternative_names': [
	// 		  'Sam Holder'
	// 		],
	// 		'account_classification': 'Personal',
	// 		'joint_account': false,
	// 		'account_matching_opt_out': false,
	// 		'secondary_identification': 'A1B2C3D4'
	// 	  }
	// 	}
	// }
	country := "GB"
	account_classification := "Personal"
	jointAccount := false
	accountMatchingOptOut := false
	accountData := models.Data{
		Data: &models.AccountData{
			Type:           "accounts",
			ID:             userId,
			OrganisationID: organizationId,
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
	displayResponse(service.Create(&accountData))
}

func sampleFetch() {
	displayResponse(service.Fetch(userId))
}

func sampleDelete() {
	displayResponse(service.Delete(userId))
}

func main() {
	sampleCreate()
	sampleFetch()
	sampleDelete()
}
