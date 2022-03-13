package repositories

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Alig1493/from3-accounts-modules/models"
)

type repository struct{}

type AccountRepository interface {
	Create(url string, accountData *models.Data) (*models.ResponseData, error)
	Fetch(url string, accountId string) (*models.ResponseData, error)
	Delete(url string, version string, accountId string) (*models.ResponseData, error)
}

func NewAccountRepository() AccountRepository {
	return &repository{}
}

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

func (*repository) Create(url string, accountData *models.Data) (*models.ResponseData, error) {
	var (
		buffer    bytes.Buffer
		errorData models.ErrorData
	)
	encodingError := json.NewEncoder(&buffer).Encode(accountData)
	if encodingError != nil {
		log.Fatalf("Failed adding a new account: %v", encodingError)
		return nil, encodingError
	}

	request, requestError := http.NewRequest("POST", url, &buffer)
	request.Header.Set("Content-Type", "application/vnd.api+json")
	if requestError != nil {
		log.Fatalf("Failed to get a new request: %v", requestError)
		return nil, requestError
	}

	client := &http.Client{}
	response, repsonseError := client.Do(request)
	if repsonseError != nil {
		log.Fatalf("Failed to get a response: %v", repsonseError)
		return nil, repsonseError
	}
	defer response.Body.Close()

	log.Println("response Status:", response.StatusCode)
	log.Println("response Headers:", response.Header)
	bodyData, bodyReadError := ioutil.ReadAll(response.Body)
	if bodyReadError != nil {
		log.Fatalf("Failed to read the response body: %v", bodyReadError)
		return nil, bodyReadError
	}
	log.Println("response Body:", string(bodyData))

	if response.StatusCode != 201 {
		json.Unmarshal(bodyData, &errorData)
		return nil, errors.New(errorData.ErrorMessage)
	}
	return &models.ResponseData{Response: accountData, StatusCode: response.StatusCode}, nil
}

func (*repository) Fetch(url string, accountId string) (*models.ResponseData, error) {
	var (
		errorData   models.ErrorData
		accountData models.Data
	)

	request, requestError := http.NewRequest("GET", url, nil)
	request.Header.Add("Accept", "application/vnd.api+json")
	if requestError != nil {
		log.Fatalf("Failed to get a new request: %v", requestError)
		return nil, requestError
	}
	client := &http.Client{}
	response, repsonseError := client.Do(request)
	if repsonseError != nil {
		log.Fatalf("Failed to get a response: %v", repsonseError)
		return nil, repsonseError
	}
	defer response.Body.Close()

	log.Println("response Status:", response.StatusCode)
	log.Println("response Headers:", response.Header)

	bodyData, bodyError := ioutil.ReadAll(response.Body)
	if bodyError != nil {
		log.Fatalf("Failed to read the response body: %v", bodyData)
		return nil, bodyError
	}
	log.Println("response Body:", string(bodyData))
	if response.StatusCode != 200 {
		json.Unmarshal(bodyData, &errorData)
		return nil, errors.New(errorData.ErrorMessage)
	}

	bodyDataUnmarshalError := json.Unmarshal(bodyData, &accountData)
	if bodyDataUnmarshalError != nil {
		log.Fatalf("Failed to unmarshal body data: %v", bodyDataUnmarshalError)
		return nil, bodyDataUnmarshalError
	}
	return &models.ResponseData{Response: &accountData, StatusCode: response.StatusCode}, nil
}

func (*repository) Delete(url string, version string, accountId string) (*models.ResponseData, error) {
	var errorData models.ErrorData
	request, requestError := http.NewRequest("DELETE", url, nil)
	request.Header.Add("Accept", "application/vnd.api+json")
	query := request.URL.Query()
	query.Add("version", version)
	request.URL.RawQuery = query.Encode()
	log.Println(request.URL.String())
	if requestError != nil {
		log.Fatalf("Failed to get a new request: %v", requestError)
		return nil, requestError
	}
	client := &http.Client{}
	response, repsonseError := client.Do(request)
	if repsonseError != nil {
		log.Fatalf("Failed to get a response: %v", repsonseError)
		return nil, repsonseError
	}
	defer response.Body.Close()

	log.Println("response Status:", response.StatusCode)
	log.Println("response Headers:", response.Header)

	bodyData, bodyError := ioutil.ReadAll(response.Body)
	if bodyError != nil {
		log.Fatalf("Failed to read the response body: %v", bodyError)
		return nil, bodyError
	}
	log.Println("response Body:", string(bodyData))

	if response.StatusCode != 204 {
		json.Unmarshal(bodyData, &errorData)
		return nil, errors.New(errorData.ErrorMessage)
	}
	return &models.ResponseData{Response: nil, StatusCode: response.StatusCode}, nil
}
