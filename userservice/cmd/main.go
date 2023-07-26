package main

import (
	"flag"
	"github.com/behnambm/todo/userservice/repo/sqliterepo"
	"log"
)

func main() {
	initDBFlag := flag.Bool("initdb", false, "Create and Seed the database")
	flag.Parse()

	sqliteRepo := sqliterepo.New("storage.db")

	if *initDBFlag {
		if err := sqliterepo.CreateTables(sqliteRepo); err != nil {
			log.Fatalln(err)
		}
		sqliterepo.SeedTables(sqliteRepo)

		return
	}

}
