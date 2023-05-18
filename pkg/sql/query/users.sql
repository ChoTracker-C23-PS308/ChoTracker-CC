-- name: GetUser :one
SELECT id
     , name
     , email
     , birth_date
     , gender
     , image_url
     , created_at
     , updated_at
FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users ( id
                  , name
                  , email
                  , birth_date
                  , gender
                  , image_url)
VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id;

-- name: UpdateUser :one
UPDATE users
SET name         = $2
  , email        = $3
  , birth_date   = $4
  , gender       = $5
  , image_url       = $6
  , updated_at   = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1;

