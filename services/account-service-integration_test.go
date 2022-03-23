package services

import (
	"net/http"
	"sync"
	"testing"

	"github.com/Alig1493/from3-accounts-modules/models"
	"github.com/Alig1493/from3-accounts-modules/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	accountData        *models.Data
	testAccountService = NewAccountService()
	once               sync.Once
)

func getAccountData() *models.Data {
	once.Do(func() {
		data := utils.GetData(true)
		accountData = &data
	})
	return accountData
}

func TestFakeApiCreate(t *testing.T) {
	data := getAccountData()
	responseData, responseError := testAccountService.Create(data)
	assert.Nil(t, responseError, responseError)
	assert.EqualValues(t, responseData, &models.ResponseData{Response: data, StatusCode: 201})
}

func TestFakeApiFetch(t *testing.T) {
	data := getAccountData()
	responseData, responseError := testAccountService.Fetch(*&data.Data.ID)
	assert.Nil(t, responseError)
	assert.EqualValues(t, responseData, &models.ResponseData{Response: data, StatusCode: 200})
}

func TestFakeApiDelete(t *testing.T) {
	data := getAccountData()
	responseData, responseError := testAccountService.Delete(*&data.Data.ID)
	assert.Nil(t, responseError)
	assert.EqualValues(t, responseData, &models.ResponseData{Response: nil, StatusCode: 204})
}

func TestFakeAPINonExistingDataFetch(t *testing.T) {
	randomUUID := uuid.NewString()
	errorMessage := "record " + randomUUID + " does not exist"
	responseData, responseError := testAccountService.Fetch(randomUUID)
	assert.ObjectsAreEqual(responseError, &models.ErrorData{ErrorMessage: errorMessage, StatusCode: http.StatusBadRequest})
	assert.Nil(t, responseData)
}

func TestFakeAPINonExistingDataDelete(t *testing.T) {
	randomUUID := uuid.NewString()
	errorMessage := "record " + randomUUID + " does not exist"
	responseData, responseError := testAccountService.Fetch(randomUUID)
	assert.ObjectsAreEqual(responseError, &models.ErrorData{ErrorMessage: errorMessage, StatusCode: http.StatusBadRequest})
	assert.Nil(t, responseData)
}

func TestFakeAPICreateWithEmptyData(t *testing.T) {
	responseData, responseError := testAccountService.Create(&models.Data{})
	assert.Nil(t, responseData)
	errorMessage := "invalid account data"
	assert.ObjectsAreEqual(responseError, &models.ErrorData{ErrorMessage: errorMessage, StatusCode: http.StatusBadRequest})
}
