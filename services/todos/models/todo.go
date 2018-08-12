package models

import "time"

type Todo struct {
	TodoID      string    `dynamo:"todo_id,hash"`
	TodoGroupID string    `dynamo:"todo_group_id"`
	Title       string    `dynamo:"title"`
	Description string    `dynamo:"description"`
	Status      string    `dynamo:"status"`
	CreatedAt   time.Time `dynamo:"created_at"`
	UpdatedAt   time.Time `dynamo:"updated_at"`
	DoneAt      time.Time `dynamo:"done_at"`
	DueAt       time.Time `dynamo:"due_at"`
	Priority    string    `dynamo:"priority"`
}
