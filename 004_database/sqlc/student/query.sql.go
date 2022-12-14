// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: query.sql

package student

import (
	"context"
	"database/sql"
)

const getStudent = `-- name: GetStudent :one
SELECT id, name
FROM students
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetStudent(ctx context.Context, id sql.NullInt32) (*Student, error) {
	row := q.db.QueryRowContext(ctx, getStudent, id)
	var i Student
	err := row.Scan(&i.ID, &i.Name)
	return &i, err
}
