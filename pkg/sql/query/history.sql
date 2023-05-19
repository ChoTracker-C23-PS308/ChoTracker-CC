
-- name: CreateHistory :one
INSERT INTO history ( id
                     , uid
                     , total_kolestrol
                     , tingkat
                     , image_url
                     , created_at
                     , updated_at)
VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
    RETURNING id;

-- name: GetHistory :many
SELECT *
FROM history
         JOIN users ON history.uid = users.id
WHERE history.uid = $1;


-- name: UpdateHistory :one
UPDATE history
SET uid                = $2
  , total_kolestrol    = $3
  , tingkat            = $4
  , image_url          = $5
  , updated_at   = CURRENT_TIMESTAMP
WHERE id = $1
    RETURNING id;

-- name: DeleteHistory :exec
DELETE
FROM history
WHERE id = $1;