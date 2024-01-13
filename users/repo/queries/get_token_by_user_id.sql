-- name: GetTokenByUserID :one
SELECT *
FROM tokens
WHERE user_id = $1;