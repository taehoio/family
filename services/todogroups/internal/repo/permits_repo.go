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

type PermitsRepo interface {
	Put(permit *model.TodoGroupPermit) error
	Get(string, string) (*model.TodoGroupPermit, error)
	ListByAccountID(string) ([]*model.TodoGroupPermit, error)
	Delete(string, string) error
}

type dynamodbPermitsRepo struct {
	PermitsRepo
	base.DynamodbRepo

	permitsTable *dynamo.Table
}

func NewPermitsRepo(ddb base.Dynamodb, cfg TodoGroupPermitsRepoConfig) PermitsRepo {
	todoGroupPermitsTable := ddb.DB().Table(cfg.FullTableName())

	return &dynamodbPermitsRepo{
		permitsTable: &todoGroupPermitsTable,
	}
}

func NewMockPermitsRepo() PermitsRepo {
	ddb := base.NewMockDynamodb()
	cfg := NewMockPermitsRepoConfig()

	return NewPermitsRepo(ddb, cfg)
}

func validatePermitInput(permit *model.TodoGroupPermit) error {
	if permit.AccountID == "" {
		return InvalidAccountIDError
	}
	if permit.TodoGroupID == "" {
		return InvalidTodoGroupIDError
	}
	isPermitTypeFound := false
	for _, permitType := range todogroups.PermitType_value {
		if permit.PermitType == todogroups.PermitType(permitType) {
			isPermitTypeFound = true
			break
		}
	}
	if !isPermitTypeFound {
		return InvalidPermitTypeError
	}

	return nil
}

func (r *dynamodbPermitsRepo) Table() *dynamo.Table {
	return r.permitsTable
}

func (r *dynamodbPermitsRepo) Put(permit *model.TodoGroupPermit) error {
	if err := validatePermitInput(permit); err != nil {
		return err
	}

	permit.UpdateAt = time.Now()

	return r.Table().Put(permit).Run()
}

func (r *dynamodbPermitsRepo) Get(accountID, todoGroupID string) (*model.TodoGroupPermit, error) {
	var permit *model.TodoGroupPermit
	err := r.Table().
		Get(accountIDFieldKey, accountID).
		Range(todoGroupIDFieldKey, dynamo.Equal, todoGroupID).
		One(&permit)
	if err != nil {
		return nil, err
	}
	return permit, nil
}

func (r *dynamodbPermitsRepo) ListByAccountID(accountID string) ([]*model.TodoGroupPermit, error) {
	var permit []*model.TodoGroupPermit
	err := r.Table().Get(accountIDFieldKey, accountID).All(&permit)
	if err != nil {
		return nil, err
	}
	return permit, nil
}

func (r *dynamodbPermitsRepo) Delete(accountID, todoGroupID string) error {
	return r.Table().
		Delete(accountIDFieldKey, accountID).
		Range(todoGroupIDFieldKey, todoGroupID).
		If(fmt.Sprintf("%s = ? AND %s = ?", accountIDFieldKey, todoGroupIDFieldKey), accountID, todoGroupID).
		Run()
}
