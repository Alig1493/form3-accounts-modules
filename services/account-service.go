package services

import (
	"github.com/Alig1493/from3-accounts-modules/models"
	"github.com/Alig1493/from3-accounts-modules/repositories"
	"github.com/Alig1493/from3-accounts-modules/utils"
)

var (
	accountRepository repositories.AccountRepository
)

type AccountService interface {
	// implement the Create, Fetch, and Delete operations on the accounts resource.
	Create(accountData *models.Data) (*models.ResponseData, error)
	Fetch(accountId string) (*models.ResponseData, error)
	Delete(accountId string) (*models.ResponseData, error)
}

type account struct{}

func NewAccountService(repository repositories.AccountRepository) AccountService {
	accountRepository = repository
	return &account{}
}

func (*account) Create(accountData *models.Data) (*models.ResponseData, error) {
	postUrl := utils.GetUrl() + "/v1/organisation/accounts"
	return accountRepository.Create(postUrl, accountData)
}

func (*account) Fetch(accountId string) (*models.ResponseData, error) {
	fetchUrl := utils.GetUrl() + "/v1/organisation/accounts/" + accountId
	return accountRepository.Fetch(fetchUrl, accountId)
}

func (*account) Delete(accountId string) (*models.ResponseData, error) {
	deleteUrl := utils.GetUrl() + "/v1/organisation/accounts/" + accountId
	return accountRepository.Delete(deleteUrl, utils.GetVersion(), accountId)
}
