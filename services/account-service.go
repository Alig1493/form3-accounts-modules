package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/Alig1493/from3-accounts-modules/models"
)

var (
	accountsUrl     = "http://" + os.Getenv("ACCOUNT_SERVICE_NAME") + ":" + os.Getenv("ACCOUNT_SERVICE_PORT")
	accountsVersion = os.Getenv("ACCOUNT_SERVICE_VERSION")
)

type AccountService interface {
	// implement the Create, Fetch, and Delete operations on the accounts resource.
	Create(accountData *models.Data) (*models.Data, error)
	Fetch(accountId string) (*models.Data, error)
	Delete(accountId string) error
}

type account struct{}

func NewAccountService() AccountService {
	return &account{}
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

func (*account) Create(accountData *models.Data) (*models.Data, error) {
	postUrl := accountsUrl + "/v1/organisation/accounts"
	var buffer bytes.Buffer
	encodingError := json.NewEncoder(&buffer).Encode(accountData)
	if encodingError != nil {
		log.Fatalf("Failed adding a new account: %v", encodingError)
		return nil, encodingError
	}

	request, requestError := http.NewRequest("POST", postUrl, &buffer)
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
		return nil, errors.New(string(bodyData))
	}
	return accountData, nil
}

func (*account) Fetch(accountId string) (*models.Data, error) {
	fetchUrl := accountsUrl + "/v1/organisation/accounts/" + accountId
	var (
		errorData   models.ErrorData
		accountData models.Data
	)

	request, requestError := http.NewRequest("GET", fetchUrl, nil)
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

	body_data, body_error := ioutil.ReadAll(response.Body)
	if body_error != nil {
		log.Fatalf("Failed to read the response body: %v", body_data)
		return nil, body_error
	}
	log.Println("response Body:", string(body_data))
	if response.StatusCode != 200 {
		json.Unmarshal(body_data, &errorData)
		return nil, errors.New(errorData.ErrorMessage)
	}

	body_data_unmarshal_error := json.Unmarshal(body_data, &accountData)
	if body_data_unmarshal_error != nil {
		log.Fatalf("Failed to unmarshal body data: %v", body_data_unmarshal_error)
		return nil, body_data_unmarshal_error
	}
	return &accountData, nil
}

func (*account) Delete(accountId string) error {
	deleteUrl := accountsUrl + "/v1/organisation/accounts/" + accountId
	log.Printf("Accounts version: %v", accountsVersion)
	var errorData models.ErrorData
	request, requestError := http.NewRequest("DELETE", deleteUrl, nil)
	request.Header.Add("Accept", "application/vnd.api+json")
	query := request.URL.Query()
	query.Add("version", accountsVersion)
	request.URL.RawQuery = query.Encode()
	log.Println(request.URL.String())
	if requestError != nil {
		log.Fatalf("Failed to get a new request: %v", requestError)
		return requestError
	}
	client := &http.Client{}
	response, repsonseError := client.Do(request)
	if repsonseError != nil {
		log.Fatalf("Failed to get a response: %v", repsonseError)
		return repsonseError
	}
	defer response.Body.Close()

	log.Println("response Status:", response.StatusCode)
	log.Println("response Headers:", response.Header)

	bodyData, bodyError := ioutil.ReadAll(response.Body)
	if bodyError != nil {
		log.Fatalf("Failed to read the response body: %v", bodyError)
		return bodyError
	}
	log.Println("response Body:", string(bodyData))

	if response.StatusCode != 204 {
		json.Unmarshal(bodyData, &errorData)
		return errors.New(errorData.ErrorMessage)
	}
	return nil
}
