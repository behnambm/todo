package httptypes

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Message string `json:"message"`
}

type CreateTodoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateTodoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateItemRequest struct {
	Title    string `json:"title"`
	Priority int    `json:"priority"`
	TodoId   int64  `json:"todoId"`
}

type UpdateItemRequest struct {
	Title    string `json:"title"`
	Priority int    `json:"priority"`
}

type BadRequestResponse struct {
	Error string `json:"error"`
}

type LoginOKResponse struct {
	Token string `json:"token"`
}
