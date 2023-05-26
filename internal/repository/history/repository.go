package history

import (
	"context"
	hModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/history"
	uModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/user"
)

type Repository interface {
	CreateHistory(ctx context.Context, arg hModel.AddHistory, au uModel.AuthUser) (string, error)
	DeleteHistory(ctx context.Context, uid string, id string, au uModel.AuthUser) error
	GetHistory(ctx context.Context, uid string, au uModel.AuthUser) ([]hModel.History, error)
}
