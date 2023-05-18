package auth

import (
	"context"
	uModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/user"
)

type Repository interface {
	GetAuthUserFull(ctx context.Context, id string) (uModel.AuthUserFull, error)
}
