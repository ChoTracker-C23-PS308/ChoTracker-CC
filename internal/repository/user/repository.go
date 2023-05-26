package user

import (
	"context"
	uModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/user"
)

type Repository interface {
	CreateUser(ctx context.Context, arg uModel.AddUser, au uModel.AuthUser) (string, error)
	GetUser(ctx context.Context, id string, au uModel.AuthUser) (uModel.User, error)
	VerifyAvailableUser(ctx context.Context, id string) (bool, error)
	UpdateUser(ctx context.Context, arg uModel.UpdateUser, au uModel.AuthUser) (string, error)
	UpdateUserImage(ctx context.Context, arg uModel.UpdateUser, au uModel.AuthUser) (string, error)
}
