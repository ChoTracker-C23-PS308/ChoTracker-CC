// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: users.sql

package sqlc

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users ( id
                  , name
                  , email
                  , birth_date
                  , gender
                  , image_url)
VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id
`

type CreateUserParams struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	BirthDate string `db:"birth_date"`
	Gender    string `db:"gender"`
	ImageUrl  string `db:"image_url"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (string, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.BirthDate,
		arg.Gender,
		arg.ImageUrl,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id
     , name
     , email
     , birth_date
     , gender
     , image_url
     , created_at
     , updated_at
FROM users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.BirthDate,
		&i.Gender,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET name         = $2
  , email        = $3
  , birth_date   = $4
  , gender       = $5
  , updated_at   = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id
`

type UpdateUserParams struct {
	ID        string `db:"id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
	BirthDate string `db:"birth_date"`
	Gender    string `db:"gender"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (string, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.Name,
		arg.Email,
		arg.BirthDate,
		arg.Gender,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const updateUserImage = `-- name: UpdateUserImage :one
UPDATE users
SET image_url         = $2
  , updated_at   = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id
`

type UpdateUserImageParams struct {
	ID       string `db:"id"`
	ImageUrl string `db:"image_url"`
}

func (q *Queries) UpdateUserImage(ctx context.Context, arg UpdateUserImageParams) (string, error) {
	row := q.db.QueryRow(ctx, updateUserImage, arg.ID, arg.ImageUrl)
	var id string
	err := row.Scan(&id)
	return id, err
}