package models

import (
	"time"

	"github.com/taeho-io/family/idl/generated/go/pb/family/accounts"
)

type Account struct {
	AccountID      string            `dynamo:"account_id,hash"`
	Type           accounts.AuthType `dynamo:"type"`
	Email          string            `dynamo:"email" index:"email-index,hash"`
	HashedPassword string            `dynamo:"hashed_password"`
	FullName       string            `dynamo:"full_name"`
	CreateAt       time.Time         `dynamo:"created_at"`
	UpdatedAt      time.Time         `dynamo:"updated_at"`
}
