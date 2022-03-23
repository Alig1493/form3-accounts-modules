package services

import (
	"net/http"
	"testing"

	"github.com/google/uuid"

	"github.com/Alig1493/from3-accounts-modules/models"
	"github.com/Alig1493/from3-accounts-modules/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAccountService struct {
	mock.Mock
}

func (mock *MockAccountService) Create(accountData *models.Data) (*models.ResponseData, *models.ErrorData) {
	args := mock.Called(accountData)
	if args.Error(1) == nil {
		return &models.ResponseData{Response: args.Get(0).(*models.Data), StatusCode: 201}, nil
	}
	return nil, &models.ErrorData{ErrorMessage: args.Error(1).Error(), StatusCode: http.StatusBadRequest}
}

func (mock *MockAccountService) Fetch(accountId string) (*models.ResponseData, *models.ErrorData) {
	args := mock.Called(accountId)
	switch args.Get(0).(*models.Data).Data.ID {
	case utils.UNIQUE_ID:
		return &models.ResponseData{Response: args.Get(0).(*models.Data), StatusCode: 200}, nil
	default:
		return nil, &models.ErrorData{ErrorMessage: "Invalid accountId", StatusCode: http.StatusBadRequest}
	}
}

func (mock *MockAccountService) Delete(accountId string) (*models.ResponseData, *models.ErrorData) {
	if accountId != utils.UNIQUE_ID {
		return nil, &models.ErrorData{ErrorMessage: "Invalid accountId", StatusCode: http.StatusBadRequest}
	}
	args := mock.Called(accountId)
	if args.Error(0) == nil {
		return &models.ResponseData{Response: nil, StatusCode: 204}, nil
	}
	return nil, &models.ErrorData{ErrorMessage: args.Error(0).Error(), StatusCode: http.StatusBadRequest}
}

func TestCreate(t *testing.T) {
	accountData := utils.GetData(false)
	mockObject := &MockAccountService{}
	mockObject.On("Create", &accountData).Return(&accountData, nil)
	responseData, responseError := mockObject.Create(&accountData)
	mockObject.AssertExpectations(t)
	assert.Nil(t, responseError)
	assert.Equal(t, responseData, &models.ResponseData{Response: &accountData, StatusCode: 201})
}

func TestDelete(t *testing.T) {
	mockObject := &MockAccountService{}
	mockObject.On("Delete", utils.UNIQUE_ID).Return(nil)
	responseData, responseError := mockObject.Delete(utils.UNIQUE_ID)
	mockObject.AssertExpectations(t)
	assert.Nil(t, responseError)
	assert.Equal(t, responseData, &models.ResponseData{Response: nil, StatusCode: 204})
}

func TestFetch(t *testing.T) {
	accountId := uuid.NewString()
	accountData := utils.GetData(false)
	mockObject := &MockAccountService{}
	mockObject.On("Fetch", accountId).Return(&accountData, nil)
	responseData, responseError := mockObject.Fetch(accountId)
	mockObject.AssertExpectations(t)
	assert.Nil(t, responseError)
	assert.Equal(t, responseData, &models.ResponseData{Response: &accountData, StatusCode: 200})
}
