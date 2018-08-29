package repo

import (
	"fmt"
	"time"

	"github.com/guregu/dynamo"

	"github.com/taeho-io/family/services/accounts/internal/model"
	"github.com/taeho-io/family/services/base"
)

const (
	accountIDFieldKey      = "account_id"
	emailFieldKey          = "email"
	emailIndexName         = "email-index"
	hashedPasswordFieldKey = "hashed_password"
	fullNameFieldKey       = "full_name"
	updatedAtFieldKey      = "updated_at"
)

type AccountsRepo interface {
	base.DynamodbRepo

	GetByID(string) (*model.Account, error)
	GetByEmail(string) (*model.Account, error)
	Put(*model.Account) error
	DeleteByID(string) error
	UpdateHashedPassword(string, string) (*model.Account, error)
	UpdateFullName(string, string) (*model.Account, error)
}

type DynamodbAccountsRepo struct {
	AccountsRepo

	accountsTable *dynamo.Table
}

func NewAccountsRepo(ddb base.Dynamodb, cfg AccountsRepoConfig) AccountsRepo {
	fullTableName := cfg.FullTableName()
	accountsTable := ddb.DB().Table(fullTableName)

	return &DynamodbAccountsRepo{
		accountsTable: &accountsTable,
	}
}

func NewMockAccountsRepo() AccountsRepo {
	ddb := base.NewMockDynamodb()
	cfg := NewMockAccountRepoConfig()

	return NewAccountsRepo(ddb, cfg)
}

// DynamodbAccountsRepo returns dynamo.DynamodbAccountsRepo
func (r *DynamodbAccountsRepo) Table() *dynamo.Table {
	return r.accountsTable
}

// GetByID returns an account model by accountID
func (r *DynamodbAccountsRepo) GetByID(accountID string) (*model.Account, error) {
	var account model.Account

	err := r.Table().
		Get(accountIDFieldKey, accountID).
		One(&account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// GetByEmail returns an account model by email
func (r *DynamodbAccountsRepo) GetByEmail(email string) (*model.Account, error) {
	var account model.Account

	err := r.Table().
		Get(emailFieldKey, email).
		Index(emailIndexName).
		One(&account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Put stores an account model
func (r *DynamodbAccountsRepo) Put(account *model.Account) error {
	return r.Table().
		Put(account).
		Run()
}

// DeleteByID removes an account model by accountID
func (r *DynamodbAccountsRepo) DeleteByID(accountID string) error {
	return r.Table().
		Delete(accountIDFieldKey, accountID).
		If(fmt.Sprintf("%s = ?", accountIDFieldKey), accountID).
		Run()
}

// UpdateHashedPassword updates the account's hashedPassword field
func (r *DynamodbAccountsRepo) UpdateHashedPassword(accountID string, newHashedPassword string) (*model.Account, error) {
	var updatedAccount model.Account

	err := r.Table().
		Update(accountIDFieldKey, accountID).
		If(fmt.Sprintf("%s = ?", accountIDFieldKey), accountID).
		Set(hashedPasswordFieldKey, newHashedPassword).
		Set(updatedAtFieldKey, time.Now()).
		Value(&updatedAccount)
	if err != nil {
		return nil, err
	}

	return &updatedAccount, nil
}

// UpdateFullName updates the account's fullName field
func (r *DynamodbAccountsRepo) UpdateFullName(accountID string, newFullName string) (*model.Account, error) {
	var updatedAccount model.Account

	err := r.Table().
		Update(accountIDFieldKey, accountID).
		If(fmt.Sprintf("%s = ?", accountIDFieldKey), accountID).
		Set(fullNameFieldKey, newFullName).
		Set(updatedAtFieldKey, time.Now()).
		Value(&updatedAccount)
	if err != nil {
		return nil, err
	}

	return &updatedAccount, nil
}
