package repo

import (
	"fmt"
	"time"

	"github.com/guregu/dynamo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"
	"github.com/taeho-io/family/services/base"
	"github.com/taeho-io/family/services/todogroups/internal/model"
)

var (
	accountIDFieldKey = "account_id"

	InvalidAccountIDError  = status.Error(codes.InvalidArgument, "invalid account_id")
	InvalidPermitTypeError = status.Error(codes.InvalidArgument, "invalid permit_type")
)

type TodoGroupPermitsRepo interface {
	base.DynamodbRepo

	Put(permit *model.TodoGroupPermit) error
	Get(string, string) (*model.TodoGroupPermit, error)
	ListByAccountID(string) ([]*model.TodoGroupPermit, error)
	Delete(string, string) error
}

type DynamodbTodoGroupPermitsRepo struct {
	TodoGroupPermitsRepo

	todoGroupPermitsTable *dynamo.Table
}

func NewTodoGroupPermitsRepo(ddb base.Dynamodb, cfg TodoGroupPermitsRepoConfig) TodoGroupPermitsRepo {
	todoGroupPermitsTable := ddb.DB().Table(cfg.FullTableName())

	return &DynamodbTodoGroupPermitsRepo{
		todoGroupPermitsTable: &todoGroupPermitsTable,
	}
}

func NewMockTodoGroupPermitsRepo() TodoGroupPermitsRepo {
	ddb := base.NewMockDynamodb()
	cfg := NewMockTodoGroupPermitsRepoConfig()

	return NewTodoGroupPermitsRepo(ddb, cfg)
}

func validateTodoGroupPermitInput(todoGroupPermit *model.TodoGroupPermit) error {
	if todoGroupPermit.AccountID == "" {
		return InvalidAccountIDError
	}
	if todoGroupPermit.TodoGroupID == "" {
		return InvalidTodoGroupIDError
	}
	isPermitTypeFound := false
	for _, permitType := range todogroups.PermitType_value {
		if todoGroupPermit.PermitType == todogroups.PermitType(permitType) {
			isPermitTypeFound = true
			break
		}
	}
	if !isPermitTypeFound {
		return InvalidPermitTypeError
	}

	return nil
}

func (r *DynamodbTodoGroupPermitsRepo) Table() *dynamo.Table {
	return r.todoGroupPermitsTable
}

func (r *DynamodbTodoGroupPermitsRepo) Put(todoGroupPermit *model.TodoGroupPermit) error {
	if err := validateTodoGroupPermitInput(todoGroupPermit); err != nil {
		return err
	}

	todoGroupPermit.UpdateAt = time.Now()

	return r.Table().Put(todoGroupPermit).Run()
}

func (r *DynamodbTodoGroupPermitsRepo) Get(accountID, todoGroupID string) (*model.TodoGroupPermit, error) {
	var todoGroupPermit *model.TodoGroupPermit
	err := r.Table().
		Get(accountIDFieldKey, accountID).
		Range(todoGroupIDFieldKey, dynamo.Equal, todoGroupID).
		One(&todoGroupPermit)
	if err != nil {
		return nil, err
	}
	return todoGroupPermit, nil
}

func (r *DynamodbTodoGroupPermitsRepo) ListByAccountID(accountID string) ([]*model.TodoGroupPermit, error) {
	var todoGroupPermits []*model.TodoGroupPermit
	err := r.Table().Get(accountIDFieldKey, accountID).All(&todoGroupPermits)
	if err != nil {
		return nil, err
	}
	return todoGroupPermits, nil
}

func (r *DynamodbTodoGroupPermitsRepo) Delete(accountID, todoGroupID string) error {
	return r.Table().
		Delete(accountIDFieldKey, accountID).
		Range(todoGroupIDFieldKey, todoGroupID).
		If(fmt.Sprintf("%s = ? AND %s = ?", accountIDFieldKey, todoGroupIDFieldKey), accountID, todoGroupID).
		Run()
}
