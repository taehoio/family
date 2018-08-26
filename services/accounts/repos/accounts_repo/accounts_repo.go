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

type Table struct {
	table.IFace

	accountsTable *dynamo.Table
}

func New(ddb dynamodb.IFace, cfg config.IFace) *Table {
	fullTableName := fullTableName(cfg)
	accountsTable := ddb.DB().Table(fullTableName)

	return &Table{
		accountsTable: &accountsTable,
	}
}

func NewMock() *Table {
	ddb := dynamodb.NewMock()
	cfg := config.NewMock()

	return New(ddb, cfg)
}

func fullTableName(cfg config.IFace) string {
	prefix := cfg.Prefix()
	tableName := cfg.Settings().DynamodbAccountsTableName

	return strings.Join([]string{prefix, tableName}, "-")
}

func (t *Table) Table() *dynamo.Table {
	return t.accountsTable
}

func (t *Table) Get(accountID string) (*models.Account, error) {
	var account models.Account

	err := t.Table().
		Get(accountIDFieldKey, accountID).
		One(&account)
	if err != nil {
		return nil, err
	}

	return &account, nil
}

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

func (t *Table) Put(account *models.Account) error {
	return t.Table().
		Put(account).
		Run()
}

func (t *Table) Delete(accountID string) error {
	return t.Table().
		Delete(accountIDFieldKey, accountID).
		If(fmt.Sprintf("%s = ?", accountIDFieldKey), accountID).
		Run()
}

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
