package types

type Todo struct {
	ID          int64
	Name        string
	Description string
	UserId      int64
}

type Item struct {
	ID       int64
	Title    string
	Priority int
	UserId   int64
	TodoId   int64
}

type MinimalItem struct {
	ID       int64
	Title    string
	Priority int
}

type TodoWithItems struct {
	Todo
	Items []MinimalItem
}
