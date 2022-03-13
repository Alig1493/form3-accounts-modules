package repositories

import (
	"sync"
	"testing"

	"github.com/Alig1493/from3-accounts-modules/models"
	"github.com/Alig1493/from3-accounts-modules/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	accountData           *models.Data
	testAccountRepository = NewAccountRepository()
	once                  sync.Once
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
	postUrl := utils.GetUrl() + "/v1/organisation/accounts"
	responseData, responseError := testAccountRepository.Create(postUrl, data)
	assert.Nil(t, responseError, responseError)
	assert.EqualValues(t, responseData, &models.ResponseData{Response: data, StatusCode: 201})
}

func TestFakeApiFetch(t *testing.T) {
	data := getAccountData()
	fetchUrl := utils.GetUrl() + "/v1/organisation/accounts/" + *&data.Data.ID
	responseData, responseError := testAccountRepository.Fetch(fetchUrl, *&data.Data.ID)
	assert.Nil(t, responseError)
	assert.EqualValues(t, responseData, &models.ResponseData{Response: data, StatusCode: 200})
}

func TestFakeApiDelete(t *testing.T) {
	data := getAccountData()
	deleteUrl := utils.GetUrl() + "/v1/organisation/accounts/" + *&data.Data.ID
	responseData, responseError := testAccountRepository.Delete(deleteUrl, utils.GetVersion(), *&data.Data.ID)
	assert.Nil(t, responseError)
	assert.EqualValues(t, responseData, &models.ResponseData{Response: nil, StatusCode: 204})
}

func TestFakeAPINonExistingDataFetch(t *testing.T) {
	randomUUID := uuid.NewString()
	errorMessage := "record " + randomUUID + " does not exist"
	fetchUrl := utils.GetUrl() + "/v1/organisation/accounts/" + randomUUID
	responseData, responseError := testAccountRepository.Fetch(fetchUrl, randomUUID)
	assert.EqualError(t, responseError, errorMessage)
	assert.Nil(t, responseData)
}

func TestFakeAPINonExistingDataDelete(t *testing.T) {
	randomUUID := uuid.NewString()
	errorMessage := "record " + randomUUID + " does not exist"
	fetchUrl := utils.GetUrl() + "/v1/organisation/accounts/" + randomUUID
	responseData, responseError := testAccountRepository.Fetch(fetchUrl, randomUUID)
	assert.EqualError(t, responseError, errorMessage)
	assert.Nil(t, responseData)
}

func TestFakeAPICreateWithEmptyData(t *testing.T) {
	postUrl := utils.GetUrl() + "/v1/organisation/accounts"
	responseData, responseError := testAccountRepository.Create(postUrl, &models.Data{})
	assert.Nil(t, responseData)
	assert.EqualError(t, responseError, "invalid account data")
}
