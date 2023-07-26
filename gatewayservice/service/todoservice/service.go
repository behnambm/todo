package todoservice

import (
	"fmt"
	"github.com/behnambm/todo/gatewayservice/types"
)

type CommandRepo interface {
	CreateTodo(types.Todo) error
	UpdateTodo(types.Todo) error
	DeleteTodo(int64) error
	CreateItem(types.Item) error
	UpdateItem(types.Item) error
	DeleteItem(int64) error
}

type QueryRepo interface {
	GetUserTodosWithItems(int64) ([]types.TodoWithItems, error)
	GetTodo(int64) (types.Todo, error)
	GetItem(int64) (types.Item, error)
}

type TodoService struct {
	command CommandRepo
	query   QueryRepo
}

func New(cmdRepo CommandRepo, queryRepo QueryRepo) TodoService {
	return TodoService{
		command: cmdRepo,
		query:   queryRepo,
	}
}

func (ts TodoService) GetUserTodosWithItems(userId int64) ([]types.TodoWithItems, error) {
	todosWithItems, err := ts.query.GetUserTodosWithItems(userId)
	if err != nil {
		return nil, fmt.Errorf("[Todo Service] GetUserTodosWithItems - %w", err)
	}

	return todosWithItems, nil
}

func (ts TodoService) GetTodo(todoId int64) (types.Todo, error) {
	todo, err := ts.query.GetTodo(todoId)
	if err != nil {
		return types.Todo{}, fmt.Errorf("[Todo Service] GetTodo - %w", err)
	}

	return todo, nil
}

func (ts TodoService) CreateTodo(todo types.Todo) error {
	if err := ts.command.CreateTodo(todo); err != nil {
		return fmt.Errorf("[Todo Service] CreateTodo - %w", err)
	}

	return nil
}

func (ts TodoService) UpdateTodo(todo types.Todo) error {
	if err := ts.command.UpdateTodo(todo); err != nil {
		return fmt.Errorf("[Todo Service] UpdateTodo - %w", err)
	}

	return nil
}

func (ts TodoService) DeleteTodo(todoId int64) error {
	if err := ts.command.DeleteTodo(todoId); err != nil {
		return fmt.Errorf("[Todo Service] DeleteTodo - %w", err)
	}

	return nil
}

func (ts TodoService) GetItem(itemId int64) (types.Item, error) {
	item, err := ts.query.GetItem(itemId)
	if err != nil {
		return types.Item{}, fmt.Errorf("[Todo Service] GetItem - %w", err)
	}

	return item, nil
}

func (ts TodoService) CreateItem(item types.Item) error {
	if err := ts.command.CreateItem(item); err != nil {
		return fmt.Errorf("[Todo Service] CreateItem - %w", err)
	}

	return nil
}

func (ts TodoService) UpdateItem(item types.Item) error {
	if err := ts.command.UpdateItem(item); err != nil {
		return fmt.Errorf("[Todo Service] UpdateItem - %w", err)
	}

	return nil
}

func (ts TodoService) DeleteItem(itemId int64) error {
	if err := ts.command.DeleteItem(itemId); err != nil {
		return fmt.Errorf("[Todo Service] DeleteItem - %w", err)
	}

	return nil
}
