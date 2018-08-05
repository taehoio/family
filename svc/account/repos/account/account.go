package account

import (
	"time"
)

type Account struct {
	AccountID      string    `dynamo:"account_id"`
	Type           string    `dynamo:"type"`
	Email          string    `dynamo:"email"`
	HashedPassword string    `dynamo:"hashed_password"`
	FullName       string    `dynamo:"full_name"`
	CreateAt       time.Time `dynamo:"created_at"`
	UpdatedAt      time.Time `dynamo:"updated_at"`
}
