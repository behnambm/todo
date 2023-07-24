package sqliterepo

import (
	"log"
)

// CreateTables will be used to create the tables that we need.
func CreateTables(repo *Repo) error {
	userTable := `
	CREATE TABLE IF NOT EXISTS todo (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255) NOT NULL,
		description TEXT,
		user_id INTEGER NOT NULL
	);
	CREATE TABLE IF NOT EXISTS item (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(255) NOT NULL,
		priority INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		todo_id INTEGER NOT NULL,
		FOREIGN KEY (todo_id) REFERENCES todo(id)
	);
	`
	_, err := repo.db.Exec(userTable)
	if err != nil {
		log.Println(err)

		return err
	}
	log.Println("DATABASE CREATED")

	return nil
}

func SeedTables(repo *Repo) {
	// Seed todo table
	todoName := "My Todo List"
	todoDescription := "This is my first todo list."

	_, err := repo.db.Exec(
		"INSERT INTO todo (name, description, user_id) VALUES (?, ?, ?)",
		todoName, todoDescription, 1,
	)
	if err != nil {
		log.Println(err)

		return
	}

	// Seed item table
	itemTitle1 := "Item 1"
	priority1 := 2
	todoID := 1

	_, err = repo.db.Exec(
		"INSERT INTO item (title, priority, user_id, todo_id) VALUES (?, ?, ?, ?)",
		itemTitle1, priority1, 1, todoID,
	)
	if err != nil {
		log.Println(err)

		return
	}

	itemTitle2 := "Item 2"
	priority2 := 1

	_, err = repo.db.Exec(
		"INSERT INTO item (title, priority, user_id, todo_id) VALUES (?, ?, ?, ?)",
		itemTitle2, priority2, 1, todoID,
	)
	if err != nil {
		log.Println(err)

		return
	}

	log.Println("DATABASE SEED SUCCESSFUL")
}
