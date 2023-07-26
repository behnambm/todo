package service

import (
	"fmt"
	"github.com/behnambm/todo/todoservice/types"
)

type Repo interface {
	GetUserTodos(int64) ([]types.Todo, error)
	GetTodoById(int64) (types.Todo, error)
	GetUserTodosWithItems(int64) ([]types.TodoWithItems, error)
	GetTodoWithItems(int64) (types.TodoWithItems, error)
	CreateTodo(types.Todo) (types.Todo, error)
	UpdateTodo(types.Todo) (types.Todo, error)
	DeleteTodo(int64) error
	GetItemById(int64) (types.Item, error)
	CreateItem(types.Item) (types.Item, error)
	UpdateItem(types.Item) (types.Item, error)
	DeleteItem(int64) error
}

type TodoService struct {
	repo Repo
}

func New(repo Repo) TodoService {
	return TodoService{
		repo: repo,
	}
}

func (ts TodoService) GetUserTodos(userId int64) ([]types.Todo, error) {
	todos, err := ts.repo.GetUserTodos(userId)
	if err != nil {
		return nil, fmt.Errorf("[Service] GetUserTodos - %w", err)
	}

	return todos, nil
}

func (ts TodoService) GetTodo(todoId int64) (types.Todo, error) {
	todo, err := ts.repo.GetTodoById(todoId)
	if err != nil {
		return types.Todo{}, fmt.Errorf("[Service] GetTodo - %w", err)
	}

	return todo, nil
}

func (ts TodoService) GetUserTodosWithItems(userId int64) ([]types.TodoWithItems, error) {
	todos, err := ts.repo.GetUserTodosWithItems(userId)
	if err != nil {
		return nil, fmt.Errorf("[Service] GetUserTodosWithItems - %w", err)
	}

	return todos, nil
}

func (ts TodoService) GetTodoWithItems(todoId int64) (types.TodoWithItems, error) {
	todos, err := ts.repo.GetTodoWithItems(todoId)
	if err != nil {
		return types.TodoWithItems{}, fmt.Errorf("[Service] GetTodoWithItems - %w", err)
	}

	return todos, nil
}

func (ts TodoService) CreateTodo(todo types.Todo) (types.Todo, error) {
	if todo.UserId < 1 {
		return types.Todo{}, fmt.Errorf("[Service] CreateTodo - user id (%d) is not valid", todo.UserId)
	}
	newTodo, err := ts.repo.CreateTodo(todo)
	if err != nil {
		return types.Todo{}, fmt.Errorf("[Service] CreateTodo - %w", err)
	}

	return newTodo, nil
}

func (ts TodoService) UpdateTodo(todo types.Todo) (types.Todo, error) {
	if todo.ID < 1 {
		return types.Todo{}, fmt.Errorf("[Service] UpdateTodo - todo id (%d) is not valid", todo.ID)
	}

	newTodo, err := ts.repo.UpdateTodo(todo)
	if err != nil {
		return types.Todo{}, fmt.Errorf("[Service] UpdateTodo - %w", err)
	}

	return newTodo, nil
}

func (ts TodoService) DeleteTodo(todoId int64) error {
	if todoId < 1 {
		return fmt.Errorf("[Service] DeleteTodo - todo id (%d) is not valid", todoId)
	}

	err := ts.repo.DeleteTodo(todoId)
	if err != nil {
		return fmt.Errorf("[Service] DeleteTodo - %w", err)
	}

	return nil
}

func (ts TodoService) GetItem(itemId int64) (types.Item, error) {
	item, err := ts.repo.GetItemById(itemId)
	if err != nil {
		return types.Item{}, fmt.Errorf("[Service] GetItem - %w", err)
	}

	return item, nil
}

func (ts TodoService) CreateItem(item types.Item) (types.Item, error) {
	if item.UserId < 1 {
		return types.Item{}, fmt.Errorf("[Service] CreateItem - user id (%d) is not valid", item.UserId)
	}

	newItem, err := ts.repo.CreateItem(item)
	if err != nil {
		return types.Item{}, fmt.Errorf("[Service] CreateItem - %w", err)
	}

	return newItem, nil
}

func (ts TodoService) UpdateItem(item types.Item) (types.Item, error) {
	if item.ID < 1 {
		return types.Item{}, fmt.Errorf("[Service] UpdateItem - item id (%d) is not valid", item.ID)
	}

	newItem, err := ts.repo.UpdateItem(item)
	if err != nil {
		return types.Item{}, fmt.Errorf("[Service] UpdateItem - %w", err)
	}

	return newItem, nil
}

func (ts TodoService) DeleteItem(itemId int64) error {
	if itemId < 1 {
		return fmt.Errorf("[Service] DeleteItem - item id (%d) is not valid", itemId)
	}

	err := ts.repo.DeleteItem(itemId)
	if err != nil {
		return fmt.Errorf("[Service] DeleteItem - %w", err)
	}

	return nil
}
