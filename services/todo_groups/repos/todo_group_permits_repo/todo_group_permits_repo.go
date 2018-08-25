package todo_group_permits_repo

import (
	"strings"
	"time"

	"github.com/guregu/dynamo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"fmt"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
	"github.com/taeho-io/family/services/base/aws/dynamodb"
	"github.com/taeho-io/family/services/base/aws/dynamodb/table"
	"github.com/taeho-io/family/services/todo_groups/config"
	"github.com/taeho-io/family/services/todo_groups/models"
)

var (
	accountIDFieldKey   = "account_id"
	todoGroupIDFieldKey = "todo_group_id"

	InvalidAccountIDError   = status.Error(codes.InvalidArgument, "invalid account_id")
	InvalidTodoGroupIDError = status.Error(codes.InvalidArgument, "invalid todo_group_id")
	InvalidPermitTypeError  = status.Error(codes.InvalidArgument, "invalid permit_type")
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

func (t *Table) validateTodoGroupPermitInput(todoGroupPermit *models.TodoGroupPermit) error {
	if todoGroupPermit.AccountID == "" {
		return InvalidAccountIDError
	}
	if todoGroupPermit.TodoGroupID == "" {
		return InvalidTodoGroupIDError
	}
	isPermitTypeFound := false
	for _, permitType := range todo_groups.PermitType_value {
		if todoGroupPermit.PermitType == todo_groups.PermitType(permitType) {
			isPermitTypeFound = true
			break
		}
	}
	if !isPermitTypeFound {
		return InvalidPermitTypeError
	}

	return nil
}

func (t *Table) Put(todoGroupPermit *models.TodoGroupPermit) error {
	if err := t.validateTodoGroupPermitInput(todoGroupPermit); err != nil {
		return err
	}

	todoGroupPermit.UpdateAt = time.Now()

	return t.Table().Put(todoGroupPermit).Run()
}

func (t *Table) Get(accountID, todoGroupID string) (*models.TodoGroupPermit, error) {
	var todoGroupPermit *models.TodoGroupPermit
	err := t.Table().
		Get(accountIDFieldKey, accountID).
		Range(todoGroupIDFieldKey, dynamo.Equal, todoGroupID).
		One(&todoGroupPermit)
	if err != nil {
		return nil, err
	}
	return todoGroupPermit, nil
}

func (t *Table) ListByAccountID(accountID string) ([]*models.TodoGroupPermit, error) {
	var todoGroupPermits []*models.TodoGroupPermit
	err := t.Table().Get(accountIDFieldKey, accountID).All(&todoGroupPermits)
	if err != nil {
		return nil, err
	}
	return todoGroupPermits, nil
}

func (t *Table) Delete(accountID, todoGroupID string) error {
	return t.Table().
		Delete(accountIDFieldKey, accountID).
		Range(todoGroupIDFieldKey, todoGroupID).
		If(fmt.Sprintf("%s = ? AND %s = ?", accountIDFieldKey, todoGroupIDFieldKey), accountID, todoGroupID).
		Run()
}
