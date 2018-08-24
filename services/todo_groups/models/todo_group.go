package models

import (
	"time"

	"github.com/taeho-io/family/idl/generated/go/pb/family/todo_groups"
)

type TodoGroup struct {
	TodoGroupID string    `dynamo:"todo_group_id,hash"`
	Title       string    `dynamo:"title"`
	Description string    `dynamo:"description"`
	CreatedBy   string    `dynamo:"created_by"`
	CreatedAt   time.Time `dynamo:"created_at"`
	UpdatedAt   time.Time `dynamo:"updated_at"`
	Order       string    `dynamo:"order"`
}

func NewTodoGroupFromProto(todoGroupProto *todo_groups.TodoGroup) *TodoGroup {
	return &TodoGroup{
		TodoGroupID: todoGroupProto.TodoGroupId,
		Title:       todoGroupProto.Title,
		Description: todoGroupProto.Description,
		CreatedBy:   todoGroupProto.CreatedBy,
		CreatedAt:   time.Unix(todoGroupProto.CreatedAt, 0),
		UpdatedAt:   time.Unix(todoGroupProto.UpdatedAt, 0),
		Order:       todoGroupProto.Order,
	}
}

func (tg *TodoGroup) ToProto() *todo_groups.TodoGroup {
	return &todo_groups.TodoGroup{
		TodoGroupId: tg.TodoGroupID,
		Title:       tg.Title,
		Description: tg.Description,
		CreatedBy:   tg.CreatedBy,
		CreatedAt:   tg.CreatedAt.Unix(),
		UpdatedAt:   tg.UpdatedAt.Unix(),
		Order:       tg.Order,
	}
}
