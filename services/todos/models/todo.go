package models

import "time"

type Todo struct {
	TodoID      string    `dynamo:"todo_id,hash"`
	ParentType  string    `dynamo:"ParentType"`
	ParentID    string    `dynamo:"parent_id"`
	Title       string    `dynamo:"title"`
	Description string    `dynamo:"description"`
	Status      string    `dynamo:"status"`
	Order       string    `dynamo:"order"`
	AssignedTo  string    `dynamo:"assigned_to"`
	Priority    string    `dynamo:"priority"`
	CreatedAt   time.Time `dynamo:"created_at"`
	UpdatedAt   time.Time `dynamo:"updated_at"`
	DoneAt      time.Time `dynamo:"done_at"`
	DueAt       time.Time `dynamo:"due_at"`
}
