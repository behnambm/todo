package service

import (
	"fmt"
	"github.com/behnambm/todo/todoservice/types"
)

type UserRepo interface {
	GetUserTodos(int) ([]types.Todo, error)
	GetUserTodosWithItems(int) ([]types.TodoWithItems, error)
	GetTodoWithItems(int) (types.TodoWithItems, error)
	CreateTodo(types.Todo) (types.Todo, error)
	UpdateTodo(types.Todo) (types.Todo, error)
	CreateItem(types.Item) (types.Item, error)
	UpdateItem(types.Item) (types.Item, error)
}

type UserService struct {
	repo UserRepo
}

func New(repo UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us UserService) GetUserTodos(userId int) ([]types.Todo, error) {
	todos, err := us.repo.GetUserTodos(userId)
	if err != nil {
		return nil, fmt.Errorf("[Service] GetUserTodos - %w", err)
	}

	return todos, nil
}

func (us UserService) GetUserTodosWithItems(userId int) ([]types.TodoWithItems, error) {
	todos, err := us.repo.GetUserTodosWithItems(userId)
	if err != nil {
		return nil, fmt.Errorf("[Service] GetUserTodosWithItems - %w", err)
	}

	return todos, nil
}

func (us UserService) GetTodoWithItems(todoId int) (types.TodoWithItems, error) {
	todos, err := us.repo.GetTodoWithItems(todoId)
	if err != nil {
		return types.TodoWithItems{}, fmt.Errorf("[Service] GetTodoWithItems - %w", err)
	}

	return todos, nil
}

func (us UserService) CreateTodo(todo types.Todo) (types.Todo, error) {
	if todo.UserId < 1 {
		return types.Todo{}, fmt.Errorf("[Service] CreateTodo - user id (%d) is not valid", todo.UserId)
	}
	newTodo, err := us.repo.CreateTodo(todo)
	if err != nil {
		return types.Todo{}, fmt.Errorf("[Service] CreateTodo - %w", err)
	}

	return newTodo, nil
}

func (us UserService) UpdateTodo(todo types.Todo) (types.Todo, error) {
	if todo.ID < 1 {
		return types.Todo{}, fmt.Errorf("[Service] UpdateTodo - todo id (%d) is not valid", todo.ID)
	}

	newTodo, err := us.repo.UpdateTodo(todo)
	if err != nil {
		return types.Todo{}, fmt.Errorf("[Service] UpdateTodo - %w", err)
	}

	return newTodo, nil
}

func (us UserService) CreateItem(item types.Item) (types.Item, error) {
	if item.UserId < 1 {
		return types.Item{}, fmt.Errorf("[Service] CreateItem - user id (%d) is not valid", item.UserId)
	}

	newItem, err := us.repo.CreateItem(item)
	if err != nil {
		return types.Item{}, fmt.Errorf("[Service] CreateItem - %w", err)
	}

	return newItem, nil
}

func (us UserService) UpdateItem(item types.Item) (types.Item, error) {
	if item.ID < 1 {
		return types.Item{}, fmt.Errorf("[Service] UpdateItem - item id (%d) is not valid", item.UserId)
	}
	if item.TodoId < 1 {
		return types.Item{}, fmt.Errorf("[Service] UpdateItem - todo id (%d) is not valid", item.UserId)
	}

	newItem, err := us.repo.UpdateItem(item)
	if err != nil {
		return types.Item{}, fmt.Errorf("[Service] UpdateItem - %w", err)
	}

	return newItem, nil
}
