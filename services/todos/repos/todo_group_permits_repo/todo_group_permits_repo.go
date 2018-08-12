package todo_group_permits_repo

import (
	"strings"

	"github.com/guregu/dynamo"

	"github.com/taeho-io/family/services/base/aws/dynamodb"
	"github.com/taeho-io/family/services/base/aws/dynamodb/table"
	"github.com/taeho-io/family/services/todos/config"
	"github.com/taeho-io/family/services/todos/models"
)

const (
	accountIDFieldKey = "account_id"
)

type Table struct {
	table.IFace

	todoGroupPermitsTable *dynamo.Table
}

func New(ddb dynamodb.IFace, cfg config.IFace) *Table {
	fullTableName := fullTableName(cfg)
	todoGroupPermitsTable := ddb.DB().Table(fullTableName)

	return &Table{
		todoGroupPermitsTable: &todoGroupPermitsTable,
	}
}

func NewMock() *Table {
	ddb := dynamodb.NewMock()
	cfg := config.NewMock()

	return New(ddb, cfg)
}

func fullTableName(cfg config.IFace) string {
	prefix := cfg.Prefix()
	tableName := cfg.Settings().DynamodbTodoGroupPermitsTableName
	return strings.Join([]string{prefix, tableName}, "-")
}

func (t *Table) Table() *dynamo.Table {
	return t.todoGroupPermitsTable
}

func (t *Table) GetByAccountID(accountID string) (*models.TodoGroupPermit, error) {
	var todoGroupPermit models.TodoGroupPermit
	err := t.Table().Get(accountIDFieldKey, accountID).One(&todoGroupPermit)
	if err != nil {
		return nil, err
	}
	return &todoGroupPermit, nil
}
