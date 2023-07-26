package todocommon

const (
	MessageTypeUserRegister = "mt_us_r"
	MessageTypeTodoCreate   = "mt_td_c"
	MessageTypeTodoUpdate   = "mt_td_u"
	MessageTypeTodoDelete   = "mt_td_d"
	MessageTypeItemCreate   = "mt_it_c"
	MessageTypeItemUpdate   = "mt_it_u"
	MessageTypeItemDelete   = "mt_it_d"
)

type UserMessage struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TodoMessage struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      int64  `json:"userId"`
}

type ItemMessage struct {
	ID       int64  `json:"id"`
	Title    string `json:"title"`
	Priority int    `json:"priority"`
	UserId   int64  `json:"userId"`
	TodoId   int64  `json:"todoId"`
}
