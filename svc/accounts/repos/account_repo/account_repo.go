package account_repo

import (
	"strings"
	"time"

	"github.com/guregu/dynamo"
	"github.com/taeho-io/family/svc/accounts/config"
	svcConfig "github.com/taeho-io/family/svc/accounts/config"
	"github.com/taeho-io/family/svc/accounts/models"
	"github.com/taeho-io/family/svc/srv/aws/dynamodb"
	"github.com/taeho-io/family/svc/srv/aws/dynamodb/table"
)

const (
	accountIDFieldKey      = "account_id"
	hashedPasswordFieldKey = "hashed_password"
	fullNameFieldKey       = "full_name"
	updatedAtFieldKey      = "updated_at"
)

type Table struct {
	table.IFace

	accountTable *dynamo.Table
}

func New(ddb dynamodb.IFace, svcCfg config.IFace) *Table {
	fullTableName := fullTableName(svcCfg)
	accountTable := ddb.DB().Table(fullTableName)

	return &Table{
		accountTable: &accountTable,
	}
}

func NewMock() *Table {
	ddb := dynamodb.NewMock()
	svcCfg := svcConfig.NewMock()

	return New(ddb, svcCfg)
}

func fullTableName(svcCfg config.IFace) string {
	prefix := svcCfg.Prefix()
	tableName := svcCfg.Settings().DynamodbAccountTableName
	return strings.Join([]string{prefix, tableName}, "-")
}

func (t *Table) Table() *dynamo.Table {
	return t.accountTable
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

func (t *Table) Put(account *models.Account) error {
	return t.Table().
		Put(account).
		Run()
}

func (t *Table) Delete(accountID string) error {
	return t.Table().
		Delete(accountIDFieldKey, accountID).
		Run()
}

func (t *Table) UpdateHashedPassword(accountID string, newHashedPassword string) (*models.Account, error) {
	var updatedAccount models.Account
	err := t.Table().
		Update(accountIDFieldKey, accountID).
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
		Set(fullNameFieldKey, newFullName).
		Set(updatedAtFieldKey, time.Now()).
		Value(&updatedAccount)
	if err != nil {
		return nil, err
	}
	return &updatedAccount, nil
}
