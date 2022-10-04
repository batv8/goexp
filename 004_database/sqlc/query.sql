create table students
(
    id   integer,
    name text
);

-- name: GetStudent :one
SELECT id, name
FROM students
WHERE id = $1 LIMIT 1;