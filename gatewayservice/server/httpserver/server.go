package httpserver

import (
	"context"
	"github.com/behnambm/todo/gatewayservice/server/httpserver/middleware"
	"github.com/behnambm/todo/gatewayservice/types"
	"github.com/behnambm/todo/gatewayservice/types/constants"
	"github.com/behnambm/todo/gatewayservice/types/httptypes"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"strconv"
	"time"
)

type AuthService interface {
	LoginUser(string, string) (string, error)
	Register(types.User) error
	IsValidWithClaim(string) (map[string]string, bool)
}

type UserService interface {
	GetUserByID(int64) (types.User, error)
}

type TodoService interface {
	GetUserTodosWithItems(int64) ([]types.TodoWithItems, error)
	GetTodo(int64) (types.Todo, error)
	CreateTodo(types.Todo) error
	UpdateTodo(types.Todo) error
	DeleteTodo(int64) error
	GetItem(int64) (types.Item, error)
	CreateItem(types.Item) error
	UpdateItem(types.Item) error
	DeleteItem(int64) error
}

type Server struct {
	listenAddr string
	authSvc    AuthService
	userSvc    UserService
	todoSvc    TodoService
}

func New(addr string, authSvc AuthService, userSvc UserService, todoSvc TodoService) Server {
	return Server{
		listenAddr: addr,
		authSvc:    authSvc,
		userSvc:    userSvc,
		todoSvc:    todoSvc,
	}
}

func (s Server) Run(ctx context.Context) {
	e := echo.New()
	go func() {
		select {
		case <-ctx.Done():
			shutdownCtx, _ := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))

			e.Shutdown(shutdownCtx)
		}
	}()

	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.Logger())

	userRoute := e.Group("/user")
	userRoute.POST("/register", s.Register)
	userRoute.POST("/login", s.Login)

	todoRoute := e.Group("/todo", middleware.Auth(s.userSvc, s.authSvc), middleware.LoginRequired())
	todoRoute.GET("/", s.TodoList)
	todoRoute.GET("/:id", s.GetTodo)
	todoRoute.POST("/", s.CreateTodo)
	todoRoute.PUT("/:id", s.UpdateTodo)
	todoRoute.DELETE("/:id", s.DeleteTodo)

	todoRoute.GET("/item/:id", s.GetItem)
	todoRoute.POST("/item/", s.CreateItem)
	todoRoute.PUT("/item/:id", s.UpdateItem)
	todoRoute.DELETE("/item/:id", s.DeleteItem)

	e.Logger.Fatal(e.Start(s.listenAddr))
}

func (s Server) Login(c echo.Context) error {
	loginRequest := httptypes.LoginRequest{}

	if err := c.Bind(&loginRequest); err != nil || loginRequest.Email == "" || loginRequest.Password == "" {
		log.Println("[HTTP] Login - json unmarshal error -", err)

		return c.JSON(http.StatusBadRequest, httptypes.BadRequestResponse{Error: "invalid data"})
	}

	jwtToken, err := s.authSvc.LoginUser(loginRequest.Email, loginRequest.Password)
	if err != nil {
		log.Println("[HTTP] Login - failed to login user -", err)

		return c.JSON(http.StatusForbidden, httptypes.BadRequestResponse{Error: "invalid credentials"})
	}

	return c.JSON(http.StatusOK, httptypes.LoginOKResponse{Token: jwtToken})
}

func (s Server) Register(c echo.Context) error {
	registerRequest := httptypes.RegisterRequest{}

	if err := c.Bind(&registerRequest); err != nil || registerRequest.Email == "" || registerRequest.Password == "" {
		log.Println("[HTTP] Register - json unmarshal error -", err)

		return c.JSON(http.StatusBadRequest, httptypes.BadRequestResponse{Error: "invalid data"})
	}

	newUser := types.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}
	err := s.authSvc.Register(newUser)
	if err != nil {
		log.Println("[HTTP] Register - failed to register user -", err)

		return c.JSON(http.StatusBadRequest, httptypes.BadRequestResponse{Error: "register  failed"})
	}

	return c.JSON(http.StatusCreated, httptypes.Response{Message: "register was successful"})
}

func (s Server) TodoList(c echo.Context) error {
	user, ok := c.Get(constants.CurrentUserKey).(types.User)
	if !ok {
		log.Println("[HTTP] TodoList - failed to get current user")

		return c.JSON(http.StatusUnauthorized, httptypes.BadRequestResponse{Error: "invalid user credentials"})
	}

	todosWithItems, err := s.todoSvc.GetUserTodosWithItems(user.ID)
	if err != nil {
		log.Println("[HTTP] TodoList - failed to get list of todos -", err)

		return c.JSON(http.StatusForbidden, httptypes.BadRequestResponse{Error: "unable to get list of todos"})
	}
	if todosWithItems == nil {
		return c.JSON(http.StatusOK, struct{}{})
	}

	return c.JSON(http.StatusOK, todosWithItems)
}

func (s Server) GetTodo(c echo.Context) error {
	// get URL parameter
	todoIdStr := c.Param("id")

	todoId, err := strconv.ParseInt(todoIdStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httptypes.BadRequestResponse{Error: "invalid todo id"})
	}

	user, ok := c.Get(constants.CurrentUserKey).(types.User)
	if !ok {
		log.Println("[HTTP] GetTodo - failed to get current user")

		return c.JSON(http.StatusUnauthorized, httptypes.BadRequestResponse{Error: "invalid user credentials"})
	}

	todo, getErr := s.todoSvc.GetTodo(todoId)
	if getErr != nil {
		log.Println("[HTTP] GetTodo - failed to get todo -", getErr)

		return c.JSON(http.StatusNotFound, httptypes.BadRequestResponse{Error: "todo not found"})
	}

	if user.ID != todo.UserId {
		return c.JSON(http.StatusForbidden, httptypes.BadRequestResponse{Error: "access denied"})
	}

	return c.JSON(http.StatusOK, todo)
}

func (s Server) CreateTodo(c echo.Context) error {
	createTodoRequest := httptypes.CreateTodoRequest{}

	err := c.Bind(&createTodoRequest)
	if err != nil || createTodoRequest.Name == "" || createTodoRequest.Description == "" {
		log.Println("[HTTP] CreateTodo - json unmarshal error -", err)

		return c.JSON(http.StatusBadRequest, httptypes.BadRequestResponse{Error: "invalid data"})
	}

	user, ok := c.Get(constants.CurrentUserKey).(types.User)
	if !ok {
		log.Println("[HTTP] CreateTodo - failed to get current user")

		return c.JSON(http.StatusUnauthorized, httptypes.BadRequestResponse{Error: "invalid user credentials"})
	}

	todo := types.Todo{
		Name:        createTodoRequest.Name,
		Description: createTodoRequest.Description,
		UserId:      user.ID,
	}

	createErr := s.todoSvc.CreateTodo(todo)
	if createErr != nil {
		log.Println("[HTTP] CreateTodo - failed to create todo -", createErr)

		return c.JSON(http.StatusForbidden, httptypes.BadRequestResponse{Error: "todo creation failed"})
	}

	return c.JSON(http.StatusCreated, httptypes.Response{Message: "todo creation was successful"})
}

func (s Server) UpdateTodo(c echo.Context) error {
	// get URL parameter
	todoIdStr := c.Param("id")
	todoId, err := strconv.ParseInt(todoIdStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httptypes.BadRequestResponse{Error: "invalid todo id"})
	}

	todoUpdateRequest := httptypes.UpdateTodoRequest{}

	bindErr := c.Bind(&todoUpdateRequest)
	if bindErr != nil || todoUpdateRequest.Name == "" || todoUpdateRequest.Description == "" {
		log.Println("[HTTP] UpdateTodo - json unmarshal error -", bindErr)

		return c.JSON(http.StatusBadRequest, httptypes.BadRequestResponse{Error: "invalid data"})
	}

	user, ok := c.Get(constants.CurrentUserKey).(types.User)
	if !ok {
		log.Println("[HTTP] UpdateTodo - failed to get current user")

		return c.JSON(http.StatusUnauthorized, httptypes.BadRequestResponse{Error: "invalid user credentials"})
	}

	todo, getErr := s.todoSvc.GetTodo(todoId)
	if getErr != nil {
		return c.JSON(http.StatusNotFound, httptypes.BadRequestResponse{Error: "todo does not exist"})
	}

	if todo.UserId != user.ID {
		return c.JSON(http.StatusForbidden, httptypes.BadRequestResponse{Error: "access denied"})
	}

	updatedTodo := types.Todo{
		ID:          todo.ID,
		Name:        todoUpdateRequest.Name,
		Description: todoUpdateRequest.Description,
	}

	updateErr := s.todoSvc.UpdateTodo(updatedTodo)
	if updateErr != nil {
		log.Println("[HTTP] UpdateTodo - failed to update todo -", updateErr)

		return c.JSON(http.StatusForbidden, httptypes.BadRequestResponse{Error: "todo update failed"})
	}

	return c.JSON(http.StatusOK, httptypes.Response{Message: "todo update was successful"})
}

func (s Server) DeleteTodo(c echo.Context) error {
	// get URL parameter
	todoIdStr := c.Param("id")
	todoId, err := strconv.ParseInt(todoIdStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httptypes.BadRequestResponse{Error: "invalid todo id"})
	}

	user, ok := c.Get(constants.CurrentUserKey).(types.User)
	if !ok {
		log.Println("[HTTP] DeleteTodo - failed to get current user")

		return c.JSON(http.StatusUnauthorized, httptypes.BadRequestResponse{Error: "invalid user credentials"})
	}

	todo, getErr := s.todoSvc.GetTodo(todoId)
	if getErr != nil {
		return c.JSON(http.StatusNotFound, httptypes.BadRequestResponse{Error: "todo does not exist"})
	}

	if todo.UserId != user.ID {
		return c.JSON(http.StatusForbidden, httptypes.BadRequestResponse{Error: "access denied"})
	}

	deleteErr := s.todoSvc.DeleteTodo(todo.ID)
	if deleteErr != nil {
		log.Println("[HTTP] DeleteTodo - failed to delete todo -", deleteErr)

		return c.JSON(http.StatusForbidden, httptypes.BadRequestResponse{Error: "todo delete failed"})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func (s Server) GetItem(c echo.Context) error {
	// get URL parameter
	itemIdStr := c.Param("id")

	itemId, err := strconv.ParseInt(itemIdStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httptypes.BadRequestResponse{Error: "invalid item id"})
	}

	user, ok := c.Get(constants.CurrentUserKey).(types.User)
	if !ok {
		log.Println("[HTTP] GetItem - failed to get current user")

		return c.JSON(http.StatusUnauthorized, httptypes.BadRequestResponse{Error: "invalid user credentials"})
	}

	item, getErr := s.todoSvc.GetItem(itemId)
	if getErr != nil {
		log.Println("[HTTP] GetItem - failed to get item -", getErr)

		return c.JSON(http.StatusNotFound, httptypes.BadRequestResponse{Error: "todo not found"})
	}

	if user.ID != item.UserId {
		return c.JSON(http.StatusForbidden, httptypes.BadRequestResponse{Error: "access denied"})
	}

	return c.JSON(http.StatusOK, item)
}

func (s Server) CreateItem(c echo.Context) error {
	createItemRequest := httptypes.CreateItemRequest{}

	err := c.Bind(&createItemRequest)
	if err != nil || createItemRequest.Title == "" || createItemRequest.TodoId == 0 {
		log.Println("[HTTP] CreateItem - json unmarshal error -", err)

		return c.JSON(http.StatusBadRequest, httptypes.BadRequestResponse{Error: "invalid data"})
	}

	user, ok := c.Get(constants.CurrentUserKey).(types.User)
	if !ok {
		log.Println("[HTTP] CreateTodo - failed to get current user")

		return c.JSON(http.StatusUnauthorized, httptypes.BadRequestResponse{Error: "invalid user credentials"})
	}

	todo, getErr := s.todoSvc.GetTodo(createItemRequest.TodoId)
	if getErr != nil {
		return c.JSON(http.StatusNotFound, httptypes.BadRequestResponse{Error: "todo does not exist"})
	}

	if todo.UserId != user.ID {
		return c.JSON(http.StatusForbidden, httptypes.BadRequestResponse{Error: "access denied"})
	}

	item := types.Item{
		Title:    createItemRequest.Title,
		Priority: createItemRequest.Priority,
		TodoId:   todo.ID,
		UserId:   user.ID,
	}

	createErr := s.todoSvc.CreateItem(item)
	if createErr != nil {
		log.Println("[HTTP] CreateItem - failed to create item -", createErr)

		return c.JSON(http.StatusForbidden, httptypes.BadRequestResponse{Error: "item creation failed"})
	}

	return c.JSON(http.StatusCreated, httptypes.Response{Message: "item creation was successful"})
}

func (s Server) UpdateItem(c echo.Context) error {
	// get URL parameter
	itemIdStr := c.Param("id")
	itemId, err := strconv.ParseInt(itemIdStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httptypes.BadRequestResponse{Error: "invalid item id"})
	}

	updateItemRequest := httptypes.UpdateItemRequest{}
	bindErr := c.Bind(&updateItemRequest)
	if bindErr != nil || updateItemRequest.Title == "" {
		log.Println("[HTTP] UpdateItem - json unmarshal error -", bindErr)

		return c.JSON(http.StatusBadRequest, httptypes.BadRequestResponse{Error: "invalid data"})
	}

	user, ok := c.Get(constants.CurrentUserKey).(types.User)
	if !ok {
		log.Println("[HTTP] UpdateItem - failed to get current user")

		return c.JSON(http.StatusUnauthorized, httptypes.BadRequestResponse{Error: "invalid user credentials"})
	}

	item, getErr := s.todoSvc.GetItem(itemId)
	if getErr != nil {
		return c.JSON(http.StatusNotFound, httptypes.BadRequestResponse{Error: "item does not exist"})
	}

	if item.UserId != user.ID {
		return c.JSON(http.StatusForbidden, httptypes.BadRequestResponse{Error: "access denied"})
	}

	updatedItem := types.Item{
		ID:       item.ID,
		Title:    updateItemRequest.Title,
		Priority: updateItemRequest.Priority,
	}

	updateErr := s.todoSvc.UpdateItem(updatedItem)
	if updateErr != nil {
		log.Println("[HTTP] UpdateItem - failed to update item -", updateErr)

		return c.JSON(http.StatusForbidden, httptypes.BadRequestResponse{Error: "item update failed"})
	}

	return c.JSON(http.StatusOK, httptypes.Response{Message: "item update was successful"})
}

func (s Server) DeleteItem(c echo.Context) error {
	// get URL parameter
	itemIdStr := c.Param("id")
	itemId, err := strconv.ParseInt(itemIdStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, httptypes.BadRequestResponse{Error: "invalid item id"})
	}

	user, ok := c.Get(constants.CurrentUserKey).(types.User)
	if !ok {
		log.Println("[HTTP] DeleteItem - failed to get current user")

		return c.JSON(http.StatusUnauthorized, httptypes.BadRequestResponse{Error: "invalid user credentials"})
	}

	item, getErr := s.todoSvc.GetItem(itemId)
	if getErr != nil {
		return c.JSON(http.StatusNotFound, httptypes.BadRequestResponse{Error: "item does not exist"})
	}

	if item.UserId != user.ID {
		return c.JSON(http.StatusForbidden, httptypes.BadRequestResponse{Error: "access denied"})
	}

	deleteErr := s.todoSvc.DeleteItem(item.ID)
	if deleteErr != nil {
		log.Println("[HTTP] DeleteItem - failed to delete item -", deleteErr)

		return c.JSON(http.StatusForbidden, httptypes.BadRequestResponse{Error: "item deletion failed"})
	}

	return c.JSON(http.StatusNoContent, nil)
}
