-- name: ListCats :many
SELECT * FROM cats ORDER BY `id` DESC LIMIT 10;

-- name: CreateCat :execresult
INSERT INTO `cats` (`name`, `age`, `breed`) 
VALUES (?, ?, ?);