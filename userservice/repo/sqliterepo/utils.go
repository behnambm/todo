package sqliterepo

import (
	"github.com/behnambm/todo/common/utils/hash"
	"log"
)

// CreateTables will be used to create the tables that we need.
func CreateTables(repo *Repo) error {
	userTable := `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email VARCHAR(255) UNIQUE,
	    name VARCHAR(255),
	    password VARCHAR(255) NOT NULL
	)
	`
	_, err := repo.db.Exec(userTable)
	if err != nil {
		return err
	}

	log.Println("DATABASE CREATED")

	return nil
}

func SeedTables(repo *Repo) {
	userEmail := "test@gmail.com"
	name := "My Name"
	passwordHash, hashErr := hash.String("123")
	if hashErr != nil {
		log.Println(hashErr)
	}
	_, err := repo.db.Exec(
		"INSERT INTO user (email, name, password) VALUES (?, ?, ?)", userEmail, name, passwordHash,
	)
	if err != nil {
		log.Println(err)
	}

	log.Println("DATABASE SEED SUCCESSFUL")
}
