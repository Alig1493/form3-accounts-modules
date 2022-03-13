package services

import (
	"os"

	"github.com/Alig1493/from3-accounts-modules/models"
	"github.com/Alig1493/from3-accounts-modules/repositories"
)

var (
	accountsUrl       = "http://" + os.Getenv("ACCOUNT_SERVICE_NAME") + ":" + os.Getenv("ACCOUNT_SERVICE_PORT")
	accountsVersion   = os.Getenv("ACCOUNT_SERVICE_VERSION")
	accountRepository repositories.AccountRepository
)

type AccountService interface {
	// implement the Create, Fetch, and Delete operations on the accounts resource.
	Create(accountData *models.Data) (*models.Data, error)
	Fetch(accountId string) (*models.Data, error)
	Delete(accountId string) error
}

type account struct{}

func NewAccountService(repository repositories.AccountRepository) AccountService {
	accountRepository = repository
	return &account{}
}

func (*account) Create(accountData *models.Data) (*models.Data, error) {
	postUrl := accountsUrl + "/v1/organisation/accounts"
	return accountRepository.Create(postUrl, accountData)
}

func (*account) Fetch(accountId string) (*models.Data, error) {
	fetchUrl := accountsUrl + "/v1/organisation/accounts/" + accountId
	return accountRepository.Fetch(fetchUrl, accountId)
}

func (*account) Delete(accountId string) error {
	deleteUrl := accountsUrl + "/v1/organisation/accounts/" + accountId
	return accountRepository.Delete(deleteUrl, accountsVersion, accountId)
}
