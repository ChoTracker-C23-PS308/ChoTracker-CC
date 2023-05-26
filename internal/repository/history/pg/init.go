package pg

import (
	"context"
	"errors"
	errorCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/error"
	"github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/sqlc"
	hModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/history"
	uModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/user"
	auRepo "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/repository/auth"
	"github.com/jackc/pgx/v4"
)

var (
	ErrUser_UserNotAuthorized = errors.New("CREATE_HISTORY.USER_NOT_AUTHORIZED")
)

type pgHistoryRepository struct {
	querier sqlc.Querier
	huRepo  auRepo.Repository
}

func (r pgHistoryRepository) CreateHistory(ctx context.Context, arg hModel.AddHistory, au uModel.AuthUser) (string, error) {
	if !au.IsSame(arg.Uid) {
		return "", ErrUser_UserNotAuthorized
	}

	id, err := r.querier.CreateHistory(ctx, sqlc.CreateHistoryParams(arg))
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("History not found")
	}
	return id, err
}

func (r pgHistoryRepository) DeleteHistory(ctx context.Context, uid string, id string, au uModel.AuthUser) error {
	if !au.IsSame(uid) {
		return ErrUser_UserNotAuthorized
	}

	err := r.querier.DeleteHistory(ctx, id)
	if err == pgx.ErrNoRows {
		return errorCommon.NewNotFoundError("History not found")
	}
	return err
}

func (r pgHistoryRepository) GetHistory(ctx context.Context, uid string, au uModel.AuthUser) ([]hModel.History, error) {
	if !au.IsSame(uid) {
		return []hModel.History{}, ErrUser_UserNotAuthorized
	}

	histories, err := r.querier.GetHistory(ctx, uid)
	if err != nil {
		return nil, err
	}
	result := make([]hModel.History, len(histories))
	for i, history := range histories {
		result[i] = hModel.History(history)
	}

	return result, nil
}

func NewPGHistoryRepository(querier sqlc.Querier) pgHistoryRepository {
	return pgHistoryRepository{querier: querier}
}
