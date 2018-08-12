package models

import "time"

type TodoGroupPermit struct {
	AccountID   string    `dynamo:"account_id,hash"`
	TodoGroupID string    `dynamo:"todo_group_id,range"`
	PermitType  string    `dynamo:"type"`
	UpdateAt    time.Time `dynamo:"updated_at"`
}
