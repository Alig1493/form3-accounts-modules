package services

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Alig1493/from3-accounts-modules/models"
	"github.com/Alig1493/from3-accounts-modules/utils"
)

type AccountService interface {
	// implement the Create, Fetch, and Delete operations on the accounts resource.
	Create(accountData *models.Data) (*models.ResponseData, *models.ErrorData)
	Fetch(accountId string) (*models.ResponseData, *models.ErrorData)
	Delete(accountId string) (*models.ResponseData, *models.ErrorData)
}

type account struct{}

func getNewReuqest(methodtype string, url string, buffer *bytes.Buffer) (*http.Request, error) {
	request, error := http.NewRequest(methodtype, url, buffer)
	request.Header.Set("Content-Type", "application/vnd.api+json")
	return request, error
}

func getClientRequest(request *http.Request) (*http.Response, error) {
	client := &http.Client{}
	response, repsonse_error := client.Do(request)
	return response, repsonse_error
}

func NewAccountService() AccountService {
	return &account{}
}

func (*account) Create(accountData *models.Data) (*models.ResponseData, *models.ErrorData) {
	var (
		buffer    bytes.Buffer
		errorData models.ErrorData
	)
	postUrl := utils.GetUrl() + "/v1/organisation/accounts"

	encodingError := json.NewEncoder(&buffer).Encode(accountData)
	if encodingError != nil {
		log.Printf("Failed adding a new account: %v", encodingError)
		return nil, &models.ErrorData{ErrorMessage: encodingError.Error(), StatusCode: http.StatusBadRequest}
	}

	request, requestError := http.NewRequest("POST", postUrl, &buffer)
	request.Header.Set("Content-Type", "application/vnd.api+json")
	if requestError != nil {
		log.Printf("Failed to get a new request: %v", requestError)
		return nil, &models.ErrorData{ErrorMessage: requestError.Error(), StatusCode: http.StatusBadRequest}
	}

	client := &http.Client{}
	response, repsonseError := client.Do(request)
	if repsonseError != nil {
		log.Printf("Failed to get a response: %v", repsonseError)
		return nil, &models.ErrorData{ErrorMessage: repsonseError.Error(), StatusCode: http.StatusBadRequest}
	}
	defer response.Body.Close()

	log.Println("response Status:", response.StatusCode)
	log.Println("response Headers:", response.Header)
	bodyData, bodyReadError := ioutil.ReadAll(response.Body)
	if bodyReadError != nil {
		log.Printf("Failed to read the response body: %v", bodyReadError)
		return nil, &models.ErrorData{ErrorMessage: repsonseError.Error(), StatusCode: http.StatusBadRequest}
	}
	log.Println("response Body:", string(bodyData))

	if response.StatusCode != 201 {
		json.Unmarshal(bodyData, &errorData)
		return nil, &errorData
	}
	return &models.ResponseData{Response: accountData, StatusCode: response.StatusCode}, nil
}

func (*account) Fetch(accountId string) (*models.ResponseData, *models.ErrorData) {
	fetchUrl := utils.GetUrl() + "/v1/organisation/accounts/" + accountId
	var (
		errorData   models.ErrorData
		accountData models.Data
	)

	request, requestError := http.NewRequest("GET", fetchUrl, nil)
	request.Header.Add("Accept", "application/vnd.api+json")
	if requestError != nil {
		log.Printf("Failed to get a new request: %v", requestError)
		return nil, &models.ErrorData{ErrorMessage: requestError.Error(), StatusCode: http.StatusBadRequest}
	}
	client := &http.Client{}
	response, repsonseError := client.Do(request)
	if repsonseError != nil {
		log.Printf("Failed to get a response: %v", repsonseError)
		return nil, &models.ErrorData{ErrorMessage: repsonseError.Error(), StatusCode: http.StatusBadRequest}
	}
	defer response.Body.Close()

	log.Println("response Status:", response.StatusCode)
	log.Println("response Headers:", response.Header)

	bodyData, bodyError := ioutil.ReadAll(response.Body)
	if bodyError != nil {
		log.Printf("Failed to read the response body: %v", bodyData)
		return nil, &models.ErrorData{ErrorMessage: bodyError.Error(), StatusCode: http.StatusBadRequest}
	}
	log.Println("response Body:", string(bodyData))
	if response.StatusCode != 200 {
		json.Unmarshal(bodyData, &errorData)
		return nil, &errorData
	}

	bodyDataUnmarshalError := json.Unmarshal(bodyData, &accountData)
	if bodyDataUnmarshalError != nil {
		log.Printf("Failed to unmarshal body data: %v", bodyDataUnmarshalError)
		return nil, &models.ErrorData{ErrorMessage: bodyDataUnmarshalError.Error(), StatusCode: http.StatusBadRequest}
	}
	return &models.ResponseData{Response: &accountData, StatusCode: response.StatusCode}, nil
}

func (*account) Delete(accountId string) (*models.ResponseData, *models.ErrorData) {
	deleteUrl := utils.GetUrl() + "/v1/organisation/accounts/" + accountId
	var errorData models.ErrorData
	request, requestError := http.NewRequest("DELETE", deleteUrl, nil)
	request.Header.Add("Accept", "application/vnd.api+json")
	query := request.URL.Query()
	query.Add("version", utils.GetVersion())
	request.URL.RawQuery = query.Encode()
	log.Println(request.URL.String())
	if requestError != nil {
		log.Printf("Failed to get a new request: %v", requestError)
		return nil, &models.ErrorData{ErrorMessage: requestError.Error(), StatusCode: http.StatusBadRequest}
	}
	client := &http.Client{}
	response, repsonseError := client.Do(request)
	if repsonseError != nil {
		log.Printf("Failed to get a response: %v", repsonseError)
		return nil, &models.ErrorData{ErrorMessage: repsonseError.Error(), StatusCode: http.StatusBadRequest}
	}
	defer response.Body.Close()

	log.Println("response Status:", response.StatusCode)
	log.Println("response Headers:", response.Header)

	bodyData, bodyError := ioutil.ReadAll(response.Body)
	if bodyError != nil {
		log.Printf("Failed to read the response body: %v", bodyError)
		return nil, &models.ErrorData{ErrorMessage: bodyError.Error(), StatusCode: http.StatusBadRequest}
	}
	log.Println("response Body:", string(bodyData))

	if response.StatusCode != 204 {
		json.Unmarshal(bodyData, &errorData)
		return nil, &errorData
	}
	return &models.ResponseData{Response: nil, StatusCode: response.StatusCode}, nil
}
