package services

import (
	"errors"
	"testing"

	"github.com/google/uuid"

	"github.com/Alig1493/from3-accounts-modules/models"
	"github.com/Alig1493/from3-accounts-modules/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
}

func (mock *MockAccountRepository) Create(url string, accountData *models.Data) (*models.ResponseData, error) {
	args := mock.Called(url, accountData)
	return &models.ResponseData{Response: args.Get(0).(*models.Data), StatusCode: 201}, args.Error(1)
}

func (mock *MockAccountRepository) Fetch(url string, accountId string) (*models.ResponseData, error) {
	args := mock.Called(url, accountId)
	switch args.Get(0).(*models.Data).Data.ID {
	case utils.UNIQUE_ID:
		return &models.ResponseData{Response: args.Get(0).(*models.Data), StatusCode: 200}, nil
	default:
		return nil, errors.New("Invalid accountId")
	}
}

func (mock *MockAccountRepository) Delete(url string, version string, accountId string) (*models.ResponseData, error) {
	if version != utils.GetVersion() {
		return &models.ResponseData{Response: nil, StatusCode: 400}, errors.New("Invalid version number")
	}
	if accountId != utils.UNIQUE_ID {
		return &models.ResponseData{Response: nil, StatusCode: 400}, errors.New("Invalid accountId")
	}
	args := mock.Called(url, version, accountId)
	return &models.ResponseData{Response: nil, StatusCode: 204}, args.Error(0)
}

func TestCreate(t *testing.T) {
	postUrl := utils.GetUrl() + "/v1/organisation/accounts"
	accountData := utils.GetData(false)
	mockObject := &MockAccountRepository{}
	mockObject.On("Create", postUrl, &accountData).Return(&accountData, nil)
	testService := NewAccountService(mockObject)
	responseData, responseError := testService.Create(&accountData)
	mockObject.AssertExpectations(t)
	assert.Nil(t, responseError)
	assert.Equal(t, responseData, &models.ResponseData{Response: &accountData, StatusCode: 201})
}

func TestDelete(t *testing.T) {
	version := utils.GetVersion()
	deleteUrl := utils.GetUrl() + "/v1/organisation/accounts/" + utils.UNIQUE_ID
	mockObject := &MockAccountRepository{}
	mockObject.On("Delete", deleteUrl, version, utils.UNIQUE_ID).Return(nil)
	testService := NewAccountService(mockObject)
	responseData, responseError := testService.Delete(utils.UNIQUE_ID)
	mockObject.AssertExpectations(t)
	assert.Nil(t, responseError)
	assert.Equal(t, responseData, &models.ResponseData{Response: nil, StatusCode: 204})
}

func TestFetch(t *testing.T) {
	accountId := uuid.NewString()
	accountData := utils.GetData(false)
	fetchUrl := utils.GetUrl() + "/v1/organisation/accounts/" + accountId
	mockObject := &MockAccountRepository{}
	mockObject.On("Fetch", fetchUrl, accountId).Return(&accountData, nil)
	testService := NewAccountService(mockObject)
	responseData, responseError := testService.Fetch(accountId)
	mockObject.AssertExpectations(t)
	assert.Nil(t, responseError)
	assert.Equal(t, responseData, &models.ResponseData{Response: &accountData, StatusCode: 200})
}
