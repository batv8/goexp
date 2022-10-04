package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github/com/batv8/goexp/004_database/sqlc/student"
	"os"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

/*
Run temporal postgres on docker: docker run --rm -e "POSTGRES_PASSWORD=Password@123" -p "5432:5432" postgres:10.21
*/
func main() {
	sqlDB, err := sql.Open("pgx", "postgres://postgres:Password@123@localhost:5432")
	handleErr(err)

	q := student.New(sqlDB)
	s, err := q.GetStudent(context.Background(), sql.NullInt32{
		Int32: 1,
		Valid: true,
	})
	if err != nil && err != sql.ErrNoRows {
		handleErr(err)
	}

	if err == sql.ErrNoRows {
		fmt.Println("No rows")
		os.Exit(2)
	}

	fmt.Printf("student 1: %+v", s)
}
