package models

import (
	"time"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
)

type TodoGroupPermit struct {
	AccountID   string                 `dynamo:"account_id,hash"`
	TodoGroupID string                 `dynamo:"todo_group_id,range"`
	PermitType  todo_groups.PermitType `dynamo:"permit_type"`
	UpdateAt    time.Time              `dynamo:"updated_at"`
}
