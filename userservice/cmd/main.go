package main

import (
	"flag"
	"github.com/behnambm/todo/userservice/repo/sqliterepo"
)

func main() {
	initDBFlag := flag.Bool("initdb", false, "Create and Seed the database")
	flag.Parse()

	sqliteRepo := sqliterepo.New("storage.db")

	if *initDBFlag {
		if err := sqliterepo.CreateTables(sqliteRepo); err != nil {
			panic(err)
		}
		sqliterepo.SeedTables(sqliteRepo)

		return
	}

}
