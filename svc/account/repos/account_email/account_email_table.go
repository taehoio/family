package account_email

import (
	"strings"

	"github.com/guregu/dynamo"
	"github.com/taeho-io/family/svc/account/config"
	svcConfig "github.com/taeho-io/family/svc/account/config"
	"github.com/taeho-io/family/svc/srv/aws/dynamodb"
	"github.com/taeho-io/family/svc/srv/aws/dynamodb/table"
)

const (
	emailFieldKey = "email"
)

type Table struct {
	table.IFace

	accountEmailTable *dynamo.Table
}

func New(ddb dynamodb.IFace, svcCfg config.IFace) *Table {
	fullTableName := fullTableName(svcCfg)
	accountEmailTable := ddb.DB().Table(fullTableName)

	return &Table{
		accountEmailTable: &accountEmailTable,
	}
}

func NewMock() *Table {
	ddb := dynamodb.NewMock()
	svcCfg := svcConfig.NewMock()

	return New(ddb, svcCfg)
}

func fullTableName(svcCfg config.IFace) string {
	prefix := svcCfg.Prefix()
	tableName := svcCfg.Settings().DynamodbAccountEmailTableName
	return strings.Join([]string{prefix, tableName}, "-")
}

func (t *Table) Table() *dynamo.Table {
	return t.accountEmailTable
}

func (t *Table) Get(email string) (*AccountEmail, error) {
	var accountEmail AccountEmail
	if err := t.Table().
		Get(emailFieldKey, email).
		One(&accountEmail); err != nil {
		return nil, err
	}
	return &accountEmail, nil
}

func (t *Table) Put(accountEmail *AccountEmail) error {
	return t.Table().
		Put(accountEmail).
		Run()
}

func (t *Table) Delete(email string) error {
	return t.Table().
		Delete(emailFieldKey, email).
		Run()
}

func (t *Table) GetAccountIDByEmail(email string) (string, error) {
	accountEmail, err := t.Get(email)
	if err != nil {
		return "", err
	}
	return accountEmail.AccountID, nil
}
