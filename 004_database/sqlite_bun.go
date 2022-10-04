package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"os"
	"path/filepath"
	"runtime"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	dbPath := "todo.sqlite"
	if home, err := os.UserHomeDir(); err == nil {
		dbPath = filepath.Join(home, "/tmp/todo.sqlite")
	}
	fmt.Println("todo database path:", dbPath)

	connString := fmt.Sprintf("file:%s?cache=shared", dbPath)

	baseDB, err := sql.Open("sqlite3", connString)
	handleErr(err)

	baseDB.SetMaxOpenConns(4 * runtime.NumCPU())
	baseDB.SetMaxIdleConns(4 * runtime.NumCPU())

	type todo struct {
		bun.BaseModel `bun:"table:todos"`

		Id   int
		Name string
		Done bool
	}
	bunDB := bun.NewDB(baseDB, sqlitedialect.New())
	var m todo
	_, err = bunDB.NewSelect().Model(&m).Where("id =  ?", 1).Exec(context.Background(), &m)
	handleErr(err)

	fmt.Printf("todo 1: %+v", m)
}
