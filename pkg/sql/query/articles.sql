-- name: GetArticle :one
SELECT id
     , author_id
     , judul_article
     , isi_article
     , author
     , image_url
     , created_at
     , updated_at
FROM articles
WHERE id = $1;

-- name: CreateArticle :one
INSERT INTO articles ( id
                     , author_id
                     , judul_article
                     , isi_article
                     , author
                     , image_url)
VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING id;

-- name: UpdateArticle :one
UPDATE articles
SET author_id        = $2
  , judul_article    = $3
  , isi_article      = $4
  , author           = $5
  , image_url        = $6
  , updated_at   = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id;

-- name: DeleteArticle :exec
DELETE
FROM articles
WHERE id = $1;