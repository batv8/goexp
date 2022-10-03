package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"path/filepath"
)

func main() {
	dbPath := "todo.sqlite"
	if home, err := os.UserHomeDir(); err == nil {
		dbPath = filepath.Join(home, "/tmp/todo.sqlite")
	}
	fmt.Println("todo database path:", dbPath)

	// step 1: connect db
	connString := fmt.Sprintf("file:%s?cache=shared", dbPath)
	db, err := sql.Open("sqlite3", connString)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(1)
	defer db.Close()

	// we got *sql.DB
	// step 2: query
	rows, err := db.Query("select id, name, done from todos")
	if err != nil {
		panic(err)
	}

	type todo struct {
		id   int
		name string
		done bool
	}

	// step 3: print all records
	fmt.Println("List of todo:")
	for rows.Next() {
		var m todo
		err := rows.Scan(&m.id, &m.name, &m.done)
		if err != nil {
			panic(err)
		}

		fmt.Printf("	*%+v\n", m)
	}
	rows.Close()
}
