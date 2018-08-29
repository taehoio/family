package model

import (
	"time"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todogroups"
)

type TodoGroupPermit struct {
	AccountID   string                `dynamo:"account_id,hash"`
	TodoGroupID string                `dynamo:"todo_group_id,range"`
	PermitType  todogroups.PermitType `dynamo:"permit_type"`
	UpdateAt    time.Time             `dynamo:"updated_at"`
}
