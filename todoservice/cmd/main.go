package main

import (
	"flag"
	"github.com/behnambm/todo/todoservice/repo/sqliterepo"
	"log"
)

func main() {
	initDBFlag := flag.Bool("initdb", false, "Create and Seed the database")
	flag.Parse()

	sqliteRepo := sqliterepo.New("storage.db")

	if *initDBFlag {
		if err := sqliterepo.CreateTables(sqliteRepo); err != nil {
			log.Println(err)
		}
		sqliterepo.SeedTables(sqliteRepo)
		return
	}
}
