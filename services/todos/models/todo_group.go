package models

import "time"

type TodoGroup struct {
	TodoGroupID    string    `dynamo:"todo_group_id,hash"`
	Title          string    `dynamo:"title"`
	Description    string    `dynamo:"description"`
	OrderedTodoIDs []string  `dynamo:"ordered_todo_ids,set"`
	CreatedAt      time.Time `dynamo:"created_at"`
	UpdatedAt      time.Time `dynamo:"updated_at"`
}
