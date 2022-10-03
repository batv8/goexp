create table students (
    id integer
);

-- name: GetStudent :one
SELECT * FROM students
WHERE id = $1 LIMIT 1;