package types

type Todo struct {
	ID          int
	Name        string
	Description string
	UserId      int
}

type Item struct {
	ID       int
	Title    string
	Priority int
	UserId   int
	TodoId   int
}

type TodoWithItems struct {
	Todo
	Items []Item
}
