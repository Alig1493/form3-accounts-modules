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
	encoding_error := json.NewEncoder(&buffer).Encode(accountData)
	if encoding_error != nil {
		log.Fatalf("Failed adding a new account: %v", encoding_error)
		return nil, encoding_error
	}

	request, request_error := http.NewRequest("POST", postUrl, &buffer)
	request.Header.Set("Content-Type", "application/vnd.api+json")
	if request_error != nil {
		log.Fatalf("Failed to get a new request: %v", request_error)
		return nil, request_error
	}

	client := &http.Client{}
	response, repsonse_error := client.Do(request)
	if repsonse_error != nil {
		log.Fatalf("Failed to get a response: %v", repsonse_error)
		return nil, request_error
	}
	defer response.Body.Close()

	log.Println("response Status:", response.StatusCode)
	log.Println("response Headers:", response.Header)
	body_data, body_read_error := ioutil.ReadAll(response.Body)
	if body_read_error != nil {
		log.Fatalf("Failed to read the response body: %v", body_data)
		return nil, body_read_error
	}
	log.Println("response Body:", string(body_data))

	if response.StatusCode != 201 {
		return nil, errors.New(string(body_data))
	}
	return accountData, nil
}

func (*account) Fetch(accountId string) (*models.Data, error) {
	fetchUrl := accountsUrl + "/v1/organisation/accounts/" + accountId
	var (
		error_data   models.ErrorData
		account_data models.Data
	)

	request, request_error := http.NewRequest("GET", fetchUrl, nil)
	request.Header.Add("Accept", "application/vnd.api+json")
	if request_error != nil {
		log.Fatalf("Failed to get a new request: %v", request_error)
		return nil, request_error
	}
	client := &http.Client{}
	response, repsonse_error := client.Do(request)
	if repsonse_error != nil {
		log.Fatalf("Failed to get a response: %v", repsonse_error)
		return nil, request_error
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
		json.Unmarshal(body_data, &error_data)
		return nil, errors.New(error_data.ErrorMessage)
	}

	body_data_unmarshal_error := json.Unmarshal(body_data, &account_data)
	if body_data_unmarshal_error != nil {
		log.Fatalf("Failed to unmarshal body data: %v", body_data_unmarshal_error)
		return nil, body_data_unmarshal_error
	}
	return &account_data, nil
}

func (*account) Delete(accountId string) error {
	deleteUrl := accountsUrl + "/v1/organisation/accounts/" + accountId
	log.Printf("Accounts version: %v", accountsVersion)
	var error_data models.ErrorData
	request, request_error := http.NewRequest("DELETE", deleteUrl, nil)
	request.Header.Add("Accept", "application/vnd.api+json")
	query := request.URL.Query()
	query.Add("version", accountsVersion)
	request.URL.RawQuery = query.Encode()
	log.Println(request.URL.String())
	if request_error != nil {
		log.Fatalf("Failed to get a new request: %v", request_error)
		return request_error
	}
	client := &http.Client{}
	response, repsonse_error := client.Do(request)
	if repsonse_error != nil {
		log.Fatalf("Failed to get a response: %v", repsonse_error)
		return request_error
	}
	defer response.Body.Close()

	log.Println("response Status:", response.StatusCode)
	log.Println("response Headers:", response.Header)

	body_data, body_error := ioutil.ReadAll(response.Body)
	if body_error != nil {
		log.Fatalf("Failed to read the response body: %v", body_data)
		return body_error
	}
	log.Println("response Body:", string(body_data))

	if response.StatusCode != 204 {
		json.Unmarshal(body_data, &error_data)
		return errors.New(error_data.ErrorMessage)
	}
	return nil
}
