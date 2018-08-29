package model

import (
	"time"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
)

type Todo struct {
	TodoID      string           `dynamo:"todo_id,hash"`
	ParentType  todos.ParentType `dynamo:"parent_type"`
	ParentID    string           `dynamo:"parent_id" index:"parent_id-index,hash"`
	Title       string           `dynamo:"title"`
	Description string           `dynamo:"description"`
	Status      todos.Status     `dynamo:"status"`
	Order       string           `dynamo:"order"`
	AssignedTo  string           `dynamo:"assigned_to"`
	Priority    todos.Priority   `dynamo:"priority"`
	CreatedAt   time.Time        `dynamo:"created_at"`
	UpdatedAt   time.Time        `dynamo:"updated_at"`
	DoneAt      time.Time        `dynamo:"done_at"`
	DueAt       time.Time        `dynamo:"due_at"`
}

func NewTodoFromProto(todoProto *todos.Todo) *Todo {
	return &Todo{
		TodoID:      todoProto.TodoId,
		ParentType:  todoProto.ParentType,
		ParentID:    todoProto.ParentId,
		Title:       todoProto.Title,
		Description: todoProto.Description,
		Status:      todoProto.Status,
		Order:       todoProto.Order,
		AssignedTo:  todoProto.AssignedTo,
		Priority:    todoProto.Priority,
		CreatedAt:   time.Unix(todoProto.CreatedAt, 0),
		UpdatedAt:   time.Unix(todoProto.UpdatedAt, 0),
		DoneAt:      time.Unix(todoProto.DoneAt, 0),
		DueAt:       time.Unix(todoProto.DueAt, 0),
	}
}

func (t *Todo) ToProto() *todos.Todo {
	return &todos.Todo{
		TodoId:      t.TodoID,
		ParentType:  t.ParentType,
		ParentId:    t.ParentID,
		Title:       t.Title,
		Description: t.Description,
		Status:      t.Status,
		Order:       t.Order,
		AssignedTo:  t.AssignedTo,
		Priority:    t.Priority,
		CreatedAt:   t.CreatedAt.Unix(),
		UpdatedAt:   t.UpdatedAt.Unix(),
		DoneAt:      t.DoneAt.Unix(),
		DueAt:       t.DueAt.Unix(),
	}
}
