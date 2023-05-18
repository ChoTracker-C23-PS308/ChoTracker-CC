package pg

import (
	"context"
	"fmt"
	"github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/sqlc"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Store struct {
	sqlc.Querier
	Db *pgxpool.Pool
}

func New(dbURL string) *Store {
	pool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		panic(err)
	}
	return &Store{
		Querier: sqlc.New(pool),
		Db:      pool,
	}
}

func (store *Store) ExecTx(ctx context.Context, fn func(sqlc.Querier) error) error {
	tx, err := store.Db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	q := sqlc.New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit(ctx)
}
