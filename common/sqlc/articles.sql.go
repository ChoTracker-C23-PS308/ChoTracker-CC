// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: articles.sql

package sqlc

import (
	"context"
)

const createArticle = `-- name: CreateArticle :one
INSERT INTO articles ( id
                     , author_id
                     , judul_article
                     , isi_article
                     , author
                     , image_url)
VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id
`

type CreateArticleParams struct {
	ID           string `db:"id"`
	AuthorID     string `db:"author_id"`
	JudulArticle string `db:"judul_article"`
	IsiArticle   string `db:"isi_article"`
	Author       string `db:"author"`
	ImageUrl     string `db:"image_url"`
}

func (q *Queries) CreateArticle(ctx context.Context, arg CreateArticleParams) (string, error) {
	row := q.db.QueryRow(ctx, createArticle,
		arg.ID,
		arg.AuthorID,
		arg.JudulArticle,
		arg.IsiArticle,
		arg.Author,
		arg.ImageUrl,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const deleteArticle = `-- name: DeleteArticle :exec
DELETE
FROM articles
WHERE id = $1
`

func (q *Queries) DeleteArticle(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteArticle, id)
	return err
}

const getAllArticle = `-- name: GetAllArticle :many
SELECT id
     , author_id
     , judul_article
     , isi_article
     , author
     , image_url
     , created_at
     , updated_at
FROM articles
`

func (q *Queries) GetAllArticle(ctx context.Context) ([]Article, error) {
	rows, err := q.db.Query(ctx, getAllArticle)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Article{}
	for rows.Next() {
		var i Article
		if err := rows.Scan(
			&i.ID,
			&i.AuthorID,
			&i.JudulArticle,
			&i.IsiArticle,
			&i.Author,
			&i.ImageUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getArticle = `-- name: GetArticle :one
SELECT id
     , author_id
     , judul_article
     , isi_article
     , author
     , image_url
     , created_at
     , updated_at
FROM articles
WHERE id = $1
`

func (q *Queries) GetArticle(ctx context.Context, id string) (Article, error) {
	row := q.db.QueryRow(ctx, getArticle, id)
	var i Article
	err := row.Scan(
		&i.ID,
		&i.AuthorID,
		&i.JudulArticle,
		&i.IsiArticle,
		&i.Author,
		&i.ImageUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateArticle = `-- name: UpdateArticle :one
UPDATE articles
SET author_id        = $2
  , judul_article    = $3
  , isi_article      = $4
  , author           = $5
  , image_url        = $6
  , updated_at   = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id
`

type UpdateArticleParams struct {
	ID           string `db:"id"`
	AuthorID     string `db:"author_id"`
	JudulArticle string `db:"judul_article"`
	IsiArticle   string `db:"isi_article"`
	Author       string `db:"author"`
	ImageUrl     string `db:"image_url"`
}

func (q *Queries) UpdateArticle(ctx context.Context, arg UpdateArticleParams) (string, error) {
	row := q.db.QueryRow(ctx, updateArticle,
		arg.ID,
		arg.AuthorID,
		arg.JudulArticle,
		arg.IsiArticle,
		arg.Author,
		arg.ImageUrl,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}