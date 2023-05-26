package pg

import (
	"context"
	"errors"
	errorCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/error"
	"github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/sqlc"
	uModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/user"
	auRepo "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/repository/auth"
	"github.com/jackc/pgx/v4"
)

type pgUserRepository struct {
	querier sqlc.Querier
	auRepo  auRepo.Repository
}

var (
	ErrCreateUser_UserNotAuthorized = errors.New("CREATE_USER.USER_NOT_AUTHORIZED")
	ErrCreateUser_UserExist         = errors.New("CREATE_USER.USER_EXISTS")
)

// CreateUser implements user.User
func (r pgUserRepository) CreateUser(ctx context.Context, arg uModel.AddUser, au uModel.AuthUser) (string, error) {
	if !au.IsSame(arg.ID) {
		return "", ErrCreateUser_UserNotAuthorized
	}

	//auf, err := r.auRepo.GetAuthUserFull(ctx, au.ID)
	//if err != nil {
	//	return "", err
	//}
	//
	//if !auf.IsEmailVerified {
	//	return "", ErrCreateUser_UserNotAuthorized
	//}

	// check if user already exists
	available, err := r.VerifyAvailableUser(ctx, au.ID)
	if err != nil {
		return "", err
	}

	if available {
		return "", ErrCreateUser_UserExist
	}

	//id, err := r.querier.CreateUser(ctx, sqlc.CreateUserParams(arg))
	id, err := r.querier.CreateUser(ctx, sqlc.CreateUserParams(arg))
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("User not found")
	}
	return id, err
}

// GetUser implements user.User
func (r pgUserRepository) GetUser(ctx context.Context, id string, au uModel.AuthUser) (uModel.User, error) {
	if !au.IsSame(id) {
		return uModel.User{}, ErrCreateUser_UserNotAuthorized
	}
	u, err := r.querier.GetUser(ctx, id)
	if err == pgx.ErrNoRows {
		return uModel.User{}, errorCommon.NewNotFoundError("User not found")
	}
	return uModel.User(u), err
}

// VerifyAvailableUser implements user.User
func (r pgUserRepository) VerifyAvailableUser(ctx context.Context, id string) (bool, error) {
	u, err := r.querier.GetUser(ctx, id)
	// user not available
	if err == pgx.ErrNoRows || (err == nil && u.ID != id) {
		return false, nil
	}
	// error
	if err != nil {
		return false, err
	}
	// user available
	return true, nil
}

// UpdateUser implements user.User
func (r pgUserRepository) UpdateUser(ctx context.Context, arg uModel.UpdateUser, au uModel.AuthUser) (string, error) {
	if !au.IsSame(arg.ID) {
		return "", ErrCreateUser_UserNotAuthorized
	}

	id, err := r.querier.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:        arg.ID,
		Name:      arg.Name,
		Email:     arg.Email,
		BirthDate: arg.BirthDate,
		Gender:    arg.Gender,
	})
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("User not found")
	}
	return id, err
}

// UpdateUser implements user.User
func (r pgUserRepository) UpdateUserImage(ctx context.Context, arg uModel.UpdateUser, au uModel.AuthUser) (string, error) {
	if !au.IsSame(arg.ID) {
		return "", ErrCreateUser_UserNotAuthorized
	}

	id, err := r.querier.UpdateUserImage(ctx, sqlc.UpdateUserImageParams{
		ID:       arg.ID,
		ImageUrl: arg.ImageUrl,
	})
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("User not found")
	}
	return id, err
}

func NewPGUserRepository(querier sqlc.Querier) pgUserRepository {
	return pgUserRepository{querier: querier}
}
