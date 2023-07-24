package sqliterepo

import (
	"database/sql"
	"fmt"
	"github.com/behnambm/todo/userservice/types"
	_ "github.com/mattn/go-sqlite3"
)

type Repo struct {
	db *sql.DB
}

func New(dsn string) *Repo {
	conn, err := sql.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}

	if pingErr := conn.Ping(); pingErr != nil {
		panic(pingErr)
	}
	return &Repo{
		db: conn,
	}
}

func (r Repo) GetUserByEmail(email string) (types.User, error) {
	row := r.db.QueryRow(`SELECT * FROM user WHERE email = ?`, email)
	if row.Err() != nil {
		return types.User{}, fmt.Errorf("[Repo] GetUserByEmail - %w", row.Err())
	}

	user := types.User{}
	if err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Password); err != nil {
		return types.User{}, fmt.Errorf("[Repo] GetUserByEmail - %w", err)
	}

	return user, nil
}

func (r Repo) CreateUser(user types.User) (types.User, error) {
	res, execErr := r.db.Exec(
		`INSERT INTO user (name, email, password) VALUES (?, ?, ?)`,
		user.Name, user.Email, user.Password,
	)
	if execErr != nil {
		return types.User{}, fmt.Errorf("[Repo] CreateUser - %w", execErr)
	}

	userId, err := res.LastInsertId()
	if err != nil {
		return types.User{}, fmt.Errorf("[Repo] CreateUser - %w", err)
	}

	user.ID = int64(int(userId))

	return types.User{}, nil
}
