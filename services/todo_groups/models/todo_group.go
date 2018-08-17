package models

import (
	"time"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todos"
)

type TodoGroup struct {
	TodoGroupID string    `dynamo:"todo_group_id,hash"`
	Title       string    `dynamo:"title"`
	Description string    `dynamo:"description"`
	CreatedBy   string    `dynamo:"created_by"`
	CreatedAt   time.Time `dynamo:"created_at"`
	UpdatedAt   time.Time `dynamo:"updated_at"`
}

func NewTodoGroupFromProto(todoGroupProto *todos.TodoGroup) *TodoGroup {
	return &TodoGroup{
		TodoGroupID: todoGroupProto.TodoGroupId,
		Title:       todoGroupProto.Title,
		Description: todoGroupProto.Description,
		CreatedBy:   todoGroupProto.CreatedBy,
		CreatedAt:   time.Unix(todoGroupProto.CreatedAt, 0),
		UpdatedAt:   time.Unix(todoGroupProto.UpdatedAt, 0),
	}
}

func (tg *TodoGroup) ToProto() *todos.TodoGroup {
	return &todos.TodoGroup{
		TodoGroupId: tg.TodoGroupID,
		Title:       tg.Title,
		Description: tg.Description,
		CreatedBy:   tg.CreatedBy,
		CreatedAt:   tg.CreatedAt.Unix(),
		UpdatedAt:   tg.UpdatedAt.Unix(),
	}
}
