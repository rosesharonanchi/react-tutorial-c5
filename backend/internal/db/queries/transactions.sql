-- name: CreateTransaction :one
INSERT INTO transactions (user_id,amount, reason, type) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetTransactions :many
SELECT * FROM transactions
WHERE user_id = $1 AND (CAST(@type_filter AS TEXT) = '' OR type = @type_filter)
ORDER BY created_at DESC;


-- name: CreateUser :one
INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;

-- name: GetUserByID :one
SELECT * FROM users WHERE id = $1;

-- name: UpdateUserSavings :one 
UPDATE users SET total_savings = $1 WHERE id = $2 RETURNING *;

-- name: UpdateUserWithdrawals :one
UPDATE users SET total_withdrawals = $1 WHERE id = $2 RETURNING *;

-- name: AddToUserSavings :one
UPDATE users 
SET total_savings = (CAST(total_savings AS NUMERIC) + CAST(@amount AS NUMERIC))::VARCHAR
WHERE id = @id
RETURNING *;

-- name: AddToUserWithdrawals :one
UPDATE users 
SET total_withdrawals = (CAST(total_withdrawals AS NUMERIC) + CAST(@amount AS NUMERIC))::VARCHAR
WHERE id = @id
RETURNING *;