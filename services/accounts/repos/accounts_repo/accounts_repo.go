package accounts_repo

import (
	"strings"
	"time"

	"github.com/guregu/dynamo"

	"fmt"

	"github.com/taeho-io/family/services/accounts/config"
	"github.com/taeho-io/family/services/accounts/models"
	"github.com/taeho-io/family/services/base/aws/dynamodb"
	"github.com/taeho-io/family/services/base/aws/dynamodb/table"
)

const (
	accountIDFieldKey      = "account_id"
	emailFieldKey          = "email"
	emailIndexName         = "email-index"
	hashedPasswordFieldKey = "hashed_password"
	fullNameFieldKey       = "full_name"
	updatedAtFieldKey      = "updated_at"
)

// IFace represents accounts_repo
type IFace interface {
	table.IFace

	GetByID(string) (*models.Account, error)
	GetByEmail(string) (*models.Account, error)
	Put(*models.Account) error
	DeleteByID(string) error
	UpdateHashedPassword(string, string) (*models.Account, error)
	UpdateFullName(string, string) (*models.Account, error)
}

// Table implements IFace
type Table struct {
	IFace

	accountsTable *dynamo.Table
}

// New returns a new instance of Table
func New(ddb dynamodb.IFace, cfg config.IFace) IFace {
	fullTableName := fullTableName(cfg)
	accountsTable := ddb.DB().Table(fullTableName)

	return &Table{
		accountsTable: &accountsTable,
	}
}

// NewMock returns a mock of Table
func NewMock() IFace {
	ddb := dynamodb.NewMock()
	cfg := config.NewMock()

	return New(ddb, cfg)
}

func fullTableName(cfg config.IFace) string {
	prefix := cfg.Prefix()
	tableName := cfg.Settings().DynamodbAccountsTableName

	return strings.Join([]string{prefix, tableName}, "-")
}

// Table returns dynamo.Table
func (t *Table) Table() *dynamo.Table {
	return t.accountsTable
}

// GetByID returns an account model by accountID
func (t *Table) GetByID(accountID string) (*models.Account, error) {
	var account models.Account

	err := t.Table().
		Get(accountIDFieldKey, accountID).
		One(&account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// GetByEmail returns an account model by email
func (t *Table) GetByEmail(email string) (*models.Account, error) {
	var account models.Account

	err := t.Table().
		Get(emailFieldKey, email).
		Index(emailIndexName).
		One(&account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

// Put stores an account model
func (t *Table) Put(account *models.Account) error {
	return t.Table().
		Put(account).
		Run()
}

// DeleteByID removes an account model by accountID
func (t *Table) DeleteByID(accountID string) error {
	return t.Table().
		Delete(accountIDFieldKey, accountID).
		If(fmt.Sprintf("%s = ?", accountIDFieldKey), accountID).
		Run()
}

// UpdateHashedPassword updates the account's hashedPassword field
func (t *Table) UpdateHashedPassword(accountID string, newHashedPassword string) (*models.Account, error) {
	var updatedAccount models.Account

	err := t.Table().
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
func (t *Table) UpdateFullName(accountID string, newFullName string) (*models.Account, error) {
	var updatedAccount models.Account

	err := t.Table().
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
