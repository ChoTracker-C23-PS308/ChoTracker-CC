package history

import (
	"context"
	hModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/history"
)

type Repository interface {
	CreateHistory(ctx context.Context, arg hModel.AddHistory, au hModel.AuthHistory) (string, error)
	DeleteHistory(ctx context.Context, id string, au hModel.AuthHistory) error
	GetHistory(ctx context.Context, uid string, au hModel.AuthHistory) ([]hModel.History, error)
}
