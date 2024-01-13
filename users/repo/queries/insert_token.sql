-- name: InsertToken :exec
INSERT INTO tokens (user_id, token)
VALUES ($1, $2) ON CONFLICT (user_id) DO
UPDATE
SET token = EXCLUDED.token;