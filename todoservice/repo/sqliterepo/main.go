package sqliterepo

import (
	"database/sql"
	"github.com/behnambm/todo/todoservice/types"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Repo struct {
	db *sql.DB
}

func New(dsn string) *Repo {
	conn, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	if pingErr := conn.Ping(); pingErr != nil {
		log.Fatalln(pingErr)
	}

	return &Repo{
		db: conn,
	}
}

func (r Repo) GetTodoById(todoId int64) (types.Todo, error) {
	row := r.db.QueryRow(`SELECT id, name, description, user_id FROM todo WHERE id = ?`, todoId)
	if row.Err() != nil {
		return types.Todo{}, row.Err()
	}

	todo := types.Todo{}
	if err := row.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.UserId); err != nil {
		return types.Todo{}, err
	}

	return todo, nil
}

func (r Repo) GetUserTodos(userId int64) ([]types.Todo, error) {
	rows, err := r.db.Query(`SELECT id, name, description, user_id FROM todo WHERE user_id = ?`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []types.Todo

	for rows.Next() {
		todo := types.Todo{}
		if err := rows.Scan(&todo.ID, &todo.Name, &todo.Description, &todo.UserId); err != nil {
			continue
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

func (r Repo) GetTodoItems(todoId int64) ([]types.MinimalItem, error) {
	rows, err := r.db.Query(`SELECT id, title, priority FROM item WHERE todo_id = ?`, todoId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []types.MinimalItem

	for rows.Next() {
		item := types.MinimalItem{}
		if scanErr := rows.Scan(&item.ID, &item.Title, &item.Priority); scanErr != nil {
			continue
		}
		items = append(items, item)
	}

	return items, nil
}

func (r Repo) GetUserTodosWithItems(userId int64) ([]types.TodoWithItems, error) {
	var todosWithItems []types.TodoWithItems

	userTodos, err := r.GetUserTodos(userId)
	if err != nil {
		return nil, err
	}

	for _, todo := range userTodos {
		items, itemErr := r.GetTodoItems(todo.ID)
		if itemErr != nil {
			continue
		}
		todoWithItems := types.TodoWithItems{
			Todo:  todo,
			Items: items,
		}
		todosWithItems = append(todosWithItems, todoWithItems)
	}

	return todosWithItems, nil
}

func (r Repo) GetTodoWithItems(todoId int64) (types.TodoWithItems, error) {
	todo, err := r.GetTodoById(todoId)
	if err != nil {
		return types.TodoWithItems{}, err
	}

	items, itemErr := r.GetTodoItems(todo.ID)
	if itemErr != nil {
		return types.TodoWithItems{}, err
	}

	todoWithItems := types.TodoWithItems{
		Todo:  todo,
		Items: items,
	}

	return todoWithItems, nil
}

func (r Repo) CreateTodo(todo types.Todo) (types.Todo, error) {
	res, err := r.db.Exec(
		`INSERT INTO todo (name, description, user_id) VALUES (?, ?, ?)`,
		todo.Name, todo.Description, todo.UserId,
	)
	if err != nil {
		return types.Todo{}, err
	}

	todoId, idErr := res.LastInsertId()
	if idErr != nil {
		return types.Todo{}, idErr
	}
	todo.ID = todoId

	return todo, nil
}

func (r Repo) UpdateTodo(todo types.Todo) (types.Todo, error) {
	_, err := r.db.Exec(
		`UPDATE todo SET name = ?, description = ? WHERE id = ?;`,
		todo.Name, todo.Description, todo.ID,
	)
	if err != nil {
		return types.Todo{}, err
	}

	return todo, nil
}

func (r Repo) DeleteTodo(todoId int64) error {
	_, err := r.db.Exec(`DELETE FROM todo WHERE id = ?;`, todoId)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) GetItemById(itemId int64) (types.Item, error) {
	row := r.db.QueryRow(`SELECT id, title, priority, user_id , todo_id FROM item WHERE id = ?`, itemId)
	if row.Err() != nil {
		return types.Item{}, row.Err()
	}

	item := types.Item{}
	if err := row.Scan(&item.ID, &item.Title, &item.Priority, &item.UserId, &item.TodoId); err != nil {
		return types.Item{}, err
	}

	return item, nil
}

func (r Repo) CreateItem(item types.Item) (types.Item, error) {
	res, err := r.db.Exec(
		`INSERT INTO item (title, priority, user_id, todo_id) VALUES (?, ?, ?, ?)`,
		item.Title, item.Priority, item.UserId, item.TodoId,
	)
	if err != nil {
		return types.Item{}, err
	}

	itemId, idErr := res.LastInsertId()
	if idErr != nil {
		return types.Item{}, idErr
	}

	item.ID = itemId

	return item, nil
}

func (r Repo) UpdateItem(item types.Item) (types.Item, error) {
	_, err := r.db.Exec(
		`UPDATE item SET title = ?, priority = ? WHERE id = ?;`,
		item.Title, item.Priority, item.ID,
	)
	if err != nil {
		return types.Item{}, err
	}

	return item, nil
}

func (r Repo) DeleteItem(itemId int64) error {
	_, err := r.db.Exec(`DELETE FROM item WHERE id = ?;`, itemId)
	if err != nil {
		return err
	}

	return nil
}
