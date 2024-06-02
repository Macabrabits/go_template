-- name: ListCats :many
SELECT * FROM cats ORDER BY `id` DESC LIMIT 10;

-- name: GetCat :one
SELECT * FROM cats 
WHERE id = sqlc.arg('id');

-- name: CreateCat :execresult
INSERT INTO `cats` (`name`, `age`, `breed`) 
VALUES (?, ?, ?);