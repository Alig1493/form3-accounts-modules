package services

import (
	"errors"
	"os"
	"testing"

	"github.com/google/uuid"

	"github.com/Alig1493/from3-accounts-modules/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const uniqueId = "4c13df18-615d-4bdc-aec2-ed9a2ac36653"

func getData() models.Data {
	userId := uniqueId
	organizationId := uniqueId
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
	return accountData
}

func getUrl() string {
	return "http://" + os.Getenv("ACCOUNT_SERVICE_NAME") + ":" + os.Getenv("ACCOUNT_SERVICE_PORT")
}

func getVersion() string {
	return os.Getenv("ACCOUNT_SERVICE_VERSION")
}

type MockAccountRepository struct {
	mock.Mock
}

func (mock *MockAccountRepository) Create(url string, accountData *models.Data) (*models.Data, error) {
	args := mock.Called(url, accountData)
	return args.Get(0).(*models.Data), args.Error(1)
}

func (mock *MockAccountRepository) Fetch(url string, accountId string) (*models.Data, error) {
	args := mock.Called(url, accountId)
	switch args.Get(0).(*models.Data).Data.ID {
	case uniqueId:
		accountData := getData()
		return &accountData, nil
	default:
		return nil, errors.New("Invalid accountId")
	}
}

func (mock *MockAccountRepository) Delete(url string, version string, accountId string) error {
	if version != getVersion() {
		return errors.New("Invalid version number")
	}
	if accountId != uniqueId {
		return errors.New("Invalid accountId")
	}
	args := mock.Called(url, version, accountId)
	return args.Error(0)
}

func TestCreate(t *testing.T) {
	postUrl := getUrl() + "/v1/organisation/accounts"
	accountData := getData()
	mockObject := &MockAccountRepository{}
	mockObject.On("Create", postUrl, &accountData).Return(&accountData, nil)
	testService := NewAccountService(mockObject)
	responseData, responseError := testService.Create(&accountData)
	mockObject.AssertExpectations(t)
	assert.Nil(t, responseError)
	assert.Equal(t, responseData, &accountData)
}

func TestDelete(t *testing.T) {
	version := getVersion()
	deleteUrl := getUrl() + "/v1/organisation/accounts/" + uniqueId
	mockObject := &MockAccountRepository{}
	mockObject.On("Delete", deleteUrl, version, uniqueId).Return(nil)
	testService := NewAccountService(mockObject)
	responseError := testService.Delete(uniqueId)
	mockObject.AssertExpectations(t)
	assert.Nil(t, responseError)
}

func TestFetch(t *testing.T) {
	accountId := uuid.NewString()
	accountData := getData()
	fetchUrl := getUrl() + "/v1/organisation/accounts/" + accountId
	mockObject := &MockAccountRepository{}
	mockObject.On("Fetch", fetchUrl, accountId).Return(&accountData, nil)
	testService := NewAccountService(mockObject)
	responseData, responseError := testService.Fetch(accountId)
	mockObject.AssertExpectations(t)
	assert.Nil(t, responseError)
	assert.Equal(t, responseData, &accountData)
}
