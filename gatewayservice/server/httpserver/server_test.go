// These test are supposed to be run after all the services are up & running
// tests will be made against the services that can be run in the
// docker compose file in  https://github.com/behnambm/todo
// It's recommended to run these tests against test databases and other services not to corrupt the actual data

package httpserver

import (
	"encoding/json"
	"fmt"
	"github.com/behnambm/todo/gatewayservice/types"
	"github.com/behnambm/todo/gatewayservice/types/httptypes"
	"github.com/behnambm/todo/todocommon"
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

var (
	BaseUrl = todocommon.GetEnvOrPanic("HTTP_LISTEN_URL")
)

func TestRegister(t *testing.T) {
	// register a random user in each run
	userId := rand.Intn(1000)
	data := fmt.Sprintf(`{"email": "test-user%d@gmail.com", "password": "123"}`, userId)

	resp, err := todocommon.PostJson(BaseUrl+"/user/register", strings.NewReader(data))
	if err != nil {
		t.Fatalf("unable to make request - error: %v", err)
	}

	if resp.StatusCode != 201 {
		t.Fatalf("got %d code expected 201", resp.StatusCode)
	}
}

func TestRegister_InvalidPayload(t *testing.T) {
	data := `{"email": "test-user@gmail.com"}`

	resp, err := todocommon.PostJson(BaseUrl+"/user/register", strings.NewReader(data))
	if err != nil {
		t.Fatalf("unable to make request - error: %v", err)
	}

	if resp.StatusCode != 400 {
		t.Fatalf("got %d code expected 400", err)
	}
}

func TestLogin(t *testing.T) {
	data := `{"email": "test@gmail.com", "password": "123"}`

	resp, err := todocommon.PostJson(BaseUrl+"/user/login", strings.NewReader(data))
	if err != nil {
		t.Fatalf("unable to make request - error: %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("got %d code expected 200", resp.StatusCode)
	}
}

func TestLogin_InvalidPayload(t *testing.T) {
	data := `{"email": "test@gmail.com"}`

	resp, err := todocommon.PostJson(BaseUrl+"/user/login", strings.NewReader(data))
	if err != nil {
		t.Fatalf("unable to make request - error: %v", err)
	}

	if resp.StatusCode != 400 {
		t.Fatalf("got %d code expected 400", resp.StatusCode)
	}
}

func TestTodo_GetUserTodo(t *testing.T) {
	data := `{"email": "test@gmail.com", "password": "123"}`

	resp, err := todocommon.PostJson(BaseUrl+"/user/login", strings.NewReader(data))
	if err != nil {
		t.Fatalf("unable to make request - error: %v", err)
	}

	token := httptypes.LoginOKResponse{}
	jsonErr := json.NewDecoder(resp.Body).Decode(&token)
	if jsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", jsonErr)
	}

	// Get list of todos

	todoResp, getTodoErr := todocommon.GetWithAuth(BaseUrl+"/todo/", token.Token)
	if getTodoErr != nil {
		t.Fatalf("unable to make request - error: %v", getTodoErr)
	}

	var todos []types.TodoWithItems
	todosJsonErr := json.NewDecoder(todoResp.Body).Decode(&todos)
	if todosJsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", todosJsonErr)
	}

	if len(todos) < 1 {
		t.Fatalf("unable to get todos list")
	}
	if len(todos[0].Items) < 1 {
		t.Fatalf("unable to get items list")
	}

	// Get one of the todos

	getTodoResp, todoGetErr := todocommon.GetWithAuth(fmt.Sprintf("%s/todo/%d", BaseUrl, todos[0].ID), token.Token)
	if todoGetErr != nil {
		t.Fatalf("unable to make request - error: %v", todoGetErr)
	}

	if getTodoResp.StatusCode != 200 {
		t.Fatalf("got %d code expected 200", getTodoResp.StatusCode)
	}

}

func TestTodo_CreateUserTodo(t *testing.T) {
	expectedTodoName := "my new todo"

	data := `{"email": "test@gmail.com", "password": "123"}`

	resp, err := todocommon.PostJson(BaseUrl+"/user/login", strings.NewReader(data))
	if err != nil {
		t.Fatalf("unable to make request - error: %v", err)
	}

	token := httptypes.LoginOKResponse{}
	jsonErr := json.NewDecoder(resp.Body).Decode(&token)
	if jsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", jsonErr)
	}

	// Create new todo

	todoData := fmt.Sprintf(`{"name": "%s", "description": "a description for this todo"}`, expectedTodoName)
	createTodoResp, createTodoErr := todocommon.PostJsonWithAuth(BaseUrl+"/todo/", token.Token, strings.NewReader(todoData))
	if createTodoErr != nil {
		t.Fatalf("unable to make request - error: %v", createTodoErr)
	}

	if createTodoResp.StatusCode != 201 {
		t.Fatalf("got %d code expected 201", createTodoResp.StatusCode)
	}

	// Get list of todos to check if the new todos is created

	todoResp, getTodoErr := todocommon.GetWithAuth(BaseUrl+"/todo/", token.Token)
	if getTodoErr != nil {
		t.Fatalf("unable to make request - error: %v", getTodoErr)
	}

	var todos []types.TodoWithItems
	todosJsonErr := json.NewDecoder(todoResp.Body).Decode(&todos)
	if todosJsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", todosJsonErr)
	}

	if len(todos) < 1 {
		t.Fatalf("unable to get todos list")
	}
	if len(todos[0].Items) < 1 {
		t.Fatalf("unable to get items list")
	}
	for _, todo := range todos {
		if todo.Name == expectedTodoName {
			// we found our newly created todo
			// let's return

			return
		}
	}
	t.Fatalf("cannot get newly created todo")
}

func TestTodo_UpdateUserTodo(t *testing.T) {
	expectedTodoName := "my new todo"
	toBeUpdatedTodoId := 0
	expectedTodoNameAfterUpdate := "my updated new todo"

	data := `{"email": "test@gmail.com", "password": "123"}`

	resp, err := todocommon.PostJson(BaseUrl+"/user/login", strings.NewReader(data))
	if err != nil {
		t.Fatalf("unable to make request - error: %v", err)
	}

	token := httptypes.LoginOKResponse{}
	jsonErr := json.NewDecoder(resp.Body).Decode(&token)
	if jsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", jsonErr)
	}

	// Create new todo

	todoData := fmt.Sprintf(`{"name": "%s", "description": "a description for this todo"}`, expectedTodoName)
	createTodoResp, createTodoErr := todocommon.PostJsonWithAuth(BaseUrl+"/todo/", token.Token, strings.NewReader(todoData))
	if createTodoErr != nil {
		t.Fatalf("unable to make request - error: %v", createTodoErr)
	}

	if createTodoResp.StatusCode != 201 {
		t.Fatalf("got %d code expected 201", createTodoResp.StatusCode)
	}

	// Get list of todos to get the ID of a new todo in order to update it

	todoResp, getTodoErr := todocommon.GetWithAuth(BaseUrl+"/todo/", token.Token)
	if getTodoErr != nil {
		t.Fatalf("unable to make request - error: %v", getTodoErr)
	}

	var todos []types.TodoWithItems
	todosJsonErr := json.NewDecoder(todoResp.Body).Decode(&todos)
	if todosJsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", todosJsonErr)
	}

	if len(todos) < 1 {
		t.Fatalf("unable to get todos list")
	}
	if len(todos[0].Items) < 1 {
		t.Fatalf("unable to get items list")
	}
	for _, todo := range todos {
		if todo.Name == expectedTodoName {
			toBeUpdatedTodoId = int(todo.ID)
			break
		}
	}

	// update the new todo

	todoUpdateData := fmt.Sprintf(
		`{"name": "%s", "description": "a description for this todo"}`, expectedTodoNameAfterUpdate,
	)
	updateResp, updateTodoErr := todocommon.PutJsonWithAuth(
		BaseUrl+"/todo/"+strconv.Itoa(toBeUpdatedTodoId), token.Token, strings.NewReader(todoUpdateData),
	)
	if updateTodoErr != nil {
		t.Fatalf("unable to make request - error: %v", updateTodoErr)
	}

	if updateResp.StatusCode != 200 {
		t.Fatalf("got %d code expected 200", updateResp.StatusCode)
	}

	// Get todo to check if the todo updated or not

	getTodoResp, todoGetErr := todocommon.GetWithAuth(fmt.Sprintf("%s/todo/%d", BaseUrl, toBeUpdatedTodoId), token.Token)
	if todoGetErr != nil {
		t.Fatalf("unable to make request - error: %v", todoGetErr)
	}

	if getTodoResp.StatusCode != 200 {
		t.Fatalf("got %d code expected 200", getTodoResp.StatusCode)
	}

	var updatedTodo types.Todo
	updatedTodoJsonErr := json.NewDecoder(getTodoResp.Body).Decode(&updatedTodo)
	if updatedTodoJsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", todosJsonErr)
	}

	if updatedTodo.Name != expectedTodoNameAfterUpdate {
		t.Fatalf("got %s Name excepted %s", updatedTodo.Name, expectedTodoNameAfterUpdate)
	}
}

func TestTodo_DeleteUserTodo(t *testing.T) {
	expectedTodoName := "my new todo"
	toBeDeletedTodoId := 0

	data := `{"email": "test@gmail.com", "password": "123"}`

	resp, err := todocommon.PostJson(BaseUrl+"/user/login", strings.NewReader(data))
	if err != nil {
		t.Fatalf("unable to make request - error: %v", err)
	}

	token := httptypes.LoginOKResponse{}
	jsonErr := json.NewDecoder(resp.Body).Decode(&token)
	if jsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", jsonErr)
	}

	// Create new todo

	todoData := fmt.Sprintf(`{"name": "%s", "description": "a description for this todo"}`, expectedTodoName)
	createTodoResp, createTodoErr := todocommon.PostJsonWithAuth(BaseUrl+"/todo/", token.Token, strings.NewReader(todoData))
	if createTodoErr != nil {
		t.Fatalf("unable to make request - error: %v", createTodoErr)
	}

	if createTodoResp.StatusCode != 201 {
		t.Fatalf("got %d code expected 201", createTodoResp.StatusCode)
	}

	// Get list of todos to get the ID of a new todo in order to update it

	todoResp, getTodoErr := todocommon.GetWithAuth(BaseUrl+"/todo/", token.Token)
	if getTodoErr != nil {
		t.Fatalf("unable to make request - error: %v", getTodoErr)
	}

	var todos []types.TodoWithItems
	todosJsonErr := json.NewDecoder(todoResp.Body).Decode(&todos)
	if todosJsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", todosJsonErr)
	}

	if len(todos) < 1 {
		t.Fatalf("unable to get todos list")
	}
	if len(todos[0].Items) < 1 {
		t.Fatalf("unable to get items list")
	}
	for _, todo := range todos {
		if todo.Name == expectedTodoName {
			toBeDeletedTodoId = int(todo.ID)
			break
		}
	}

	// delete the new todo

	deleteResp, deleteTodoErr := todocommon.DeleteWithAuth(BaseUrl+"/todo/"+strconv.Itoa(toBeDeletedTodoId), token.Token)
	if deleteTodoErr != nil {
		t.Fatalf("unable to make request - error: %v", deleteTodoErr)
	}

	if deleteResp.StatusCode != 204 {
		t.Fatalf("got %d code expected 204", deleteResp.StatusCode)
	}

	// Get todo to check if the todo deleted or not

	getTodoResp, todoGetErr := todocommon.GetWithAuth(fmt.Sprintf("%s/todo/%d", BaseUrl, toBeDeletedTodoId), token.Token)
	if todoGetErr != nil {
		t.Fatalf("unable to make request - error: %v", todoGetErr)
	}

	if getTodoResp.StatusCode != 404 {
		t.Fatalf("got %d code expected 404", getTodoResp.StatusCode)
	}
}

func TestItem_GetItem(t *testing.T) {
	itemIdToGet := 0

	data := `{"email": "test@gmail.com", "password": "123"}`

	resp, err := todocommon.PostJson(BaseUrl+"/user/login", strings.NewReader(data))
	if err != nil {
		t.Fatalf("unable to make request - error: %v", err)
	}

	token := httptypes.LoginOKResponse{}
	jsonErr := json.NewDecoder(resp.Body).Decode(&token)
	if jsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", jsonErr)
	}

	// Get list of todos

	todoResp, getTodoErr := todocommon.GetWithAuth(BaseUrl+"/todo/", token.Token)
	if getTodoErr != nil {
		t.Fatalf("unable to make request - error: %v", getTodoErr)
	}

	var todos []types.TodoWithItems
	todosJsonErr := json.NewDecoder(todoResp.Body).Decode(&todos)
	if todosJsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", todosJsonErr)
	}

	if len(todos) < 1 {
		t.Fatalf("unable to get todos list")
	}
	if len(todos[0].Items) < 1 {
		t.Fatalf("unable to get items list")
	}
	itemIdToGet = int(todos[0].Items[0].ID)

	// Get one of the items

	getItemResp, itemGetErr := todocommon.GetWithAuth(fmt.Sprintf("%s/todo/item/%d", BaseUrl, itemIdToGet), token.Token)
	if itemGetErr != nil {
		t.Fatalf("unable to make request - error: %v", itemGetErr)
	}

	if getItemResp.StatusCode != 200 {
		t.Fatalf("got %d code expected 200", getItemResp.StatusCode)
	}

}

func TestItem_CreateItem(t *testing.T) {
	todoIdToCreateItem := 0

	data := `{"email": "test@gmail.com", "password": "123"}`

	resp, err := todocommon.PostJson(BaseUrl+"/user/login", strings.NewReader(data))
	if err != nil {
		t.Fatalf("unable to make request - error: %v", err)
	}

	token := httptypes.LoginOKResponse{}
	jsonErr := json.NewDecoder(resp.Body).Decode(&token)
	if jsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", jsonErr)
	}

	// Get list of todos

	todoResp, getTodoErr := todocommon.GetWithAuth(BaseUrl+"/todo/", token.Token)
	if getTodoErr != nil {
		t.Fatalf("unable to make request - error: %v", getTodoErr)
	}

	var todos []types.TodoWithItems
	todosJsonErr := json.NewDecoder(todoResp.Body).Decode(&todos)
	if todosJsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", todosJsonErr)
	}

	if len(todos) < 1 {
		t.Fatalf("unable to get todos list")
	}
	if len(todos[0].Items) < 1 {
		t.Fatalf("unable to get items list")
	}
	todoIdToCreateItem = int(todos[0].Items[0].ID)

	// Create new item

	itemData := fmt.Sprintf(`{"title": "My first item", "todoId": %d, "priority": 1}`, todoIdToCreateItem)
	itemCreateResp, itemCreateErr := todocommon.PostJsonWithAuth(
		fmt.Sprintf("%s/todo/item/", BaseUrl), token.Token, strings.NewReader(itemData),
	)
	if itemCreateErr != nil {
		t.Fatalf("unable to make request - error: %v", itemCreateErr)
	}

	if itemCreateResp.StatusCode != 201 {
		t.Fatalf("got %d code expected 201", itemCreateResp.StatusCode)
	}
}

func TestItem_UpdateItem(t *testing.T) {
	toBeUpdatedItemId := 0
	toBeUpdatedItemTodoId := 0
	newItemTitleAfterUpdate := "new item title after update"

	data := `{"email": "test@gmail.com", "password": "123"}`

	resp, err := todocommon.PostJson(BaseUrl+"/user/login", strings.NewReader(data))
	if err != nil {
		t.Fatalf("unable to make request - error: %v", err)
	}

	token := httptypes.LoginOKResponse{}
	jsonErr := json.NewDecoder(resp.Body).Decode(&token)
	if jsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", jsonErr)
	}

	// Get list of todos

	todoResp, getTodoErr := todocommon.GetWithAuth(BaseUrl+"/todo/", token.Token)
	if getTodoErr != nil {
		t.Fatalf("unable to make request - error: %v", getTodoErr)
	}

	var todos []types.TodoWithItems
	todosJsonErr := json.NewDecoder(todoResp.Body).Decode(&todos)
	if todosJsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", todosJsonErr)
	}

	if len(todos) < 1 {
		t.Fatalf("unable to get todos list")
	}
	if len(todos[0].Items) < 1 {
		t.Fatalf("unable to get items list")
	}
	for _, todo := range todos {
		if len(todo.Items) > 1 {
			for _, item := range todo.Items {
				toBeUpdatedItemId = int(item.ID)
				toBeUpdatedItemTodoId = int(todo.ID)
				break
			}
		}
	}

	// update the item

	itemData := fmt.Sprintf(`{"title": "%s", "todoId": %d, "priority": 1}`, newItemTitleAfterUpdate, toBeUpdatedItemTodoId)
	itemUpdateResp, itemUpdateErr := todocommon.PutJsonWithAuth(
		fmt.Sprintf("%s/todo/item/%d", BaseUrl, toBeUpdatedItemId), token.Token, strings.NewReader(itemData),
	)
	if itemUpdateErr != nil {
		t.Fatalf("unable to make request - error: %v", itemUpdateErr)
	}

	if itemUpdateResp.StatusCode != 200 {
		t.Fatalf("got %d code expected 200", itemUpdateResp.StatusCode)
	}

	// Get list of todos to check if the item updated or not

	todoResp, getTodoErr = todocommon.GetWithAuth(BaseUrl+"/todo/", token.Token)
	if getTodoErr != nil {
		t.Fatalf("unable to make request - error: %v", getTodoErr)
	}

	todosJsonErr = json.NewDecoder(todoResp.Body).Decode(&todos)
	if todosJsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", todosJsonErr)
	}
	if len(todos) < 1 {
		t.Fatalf("unable to get todos list")
	}
	if len(todos[0].Items) < 1 {
		t.Fatalf("unable to get items list")
	}
	for _, todo := range todos {
		if len(todo.Items) > 1 {
			for _, item := range todo.Items {
				if item.Title == newItemTitleAfterUpdate {
					return
				}
			}
		}
	}
	t.Fatalf("cannot get newly updated todo")
}

func TestItem_DeleteItem(t *testing.T) {
	toBeDeletedItemId := 0

	data := `{"email": "test@gmail.com", "password": "123"}`

	resp, err := todocommon.PostJson(BaseUrl+"/user/login", strings.NewReader(data))
	if err != nil {
		t.Fatalf("unable to make request - error: %v", err)
	}

	token := httptypes.LoginOKResponse{}
	jsonErr := json.NewDecoder(resp.Body).Decode(&token)
	if jsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", jsonErr)
	}

	// Get list of todos

	todoResp, getTodoErr := todocommon.GetWithAuth(BaseUrl+"/todo/", token.Token)
	if getTodoErr != nil {
		t.Fatalf("unable to make request - error: %v", getTodoErr)
	}

	var todos []types.TodoWithItems
	todosJsonErr := json.NewDecoder(todoResp.Body).Decode(&todos)
	if todosJsonErr != nil {
		t.Fatalf("unable to unmarshal - error: %v", todosJsonErr)
	}

	if len(todos) < 1 {
		t.Fatalf("unable to get todos list")
	}
	if len(todos[0].Items) < 1 {
		t.Fatalf("unable to get items list")
	}
	for _, todo := range todos {
		for _, item := range todo.Items {
			toBeDeletedItemId = int(item.ID)
			break
		}
	}

	// delete the item

	itemDeleteResp, itemDeleteErr := todocommon.DeleteWithAuth(
		fmt.Sprintf("%s/todo/item/%d", BaseUrl, toBeDeletedItemId), token.Token,
	)
	if itemDeleteErr != nil {
		t.Fatalf("unable to m ake request - error: %v", itemDeleteErr)
	}

	if itemDeleteResp.StatusCode != 204 {
		t.Fatalf("got %d code expected 204", itemDeleteResp.StatusCode)
	}
}
