package user

import (
	"context"
	uModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/user"
)

type Repository interface {
	CreateUser(ctx context.Context, arg uModel.AddUser, au uModel.AuthUser) (string, error)
	//DeleteUser(ctx context.Context, id string) error
	GetUser(ctx context.Context, id string, au uModel.AuthUser) (uModel.User, error)
	//GetUserHistory(ctx context.Context, id string) ([]oModel.Order, error)
	VerifyAvailableUser(ctx context.Context, id string) (bool, error)
	//ListUsers(ctx context.Context) ([]uModel.User, error)
	UpdateUser(ctx context.Context, arg uModel.UpdateUser, au uModel.AuthUser) (string, error)
}
