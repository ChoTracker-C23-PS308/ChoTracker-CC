package firebase

import (
	"context"
	"firebase.google.com/go/v4/auth"
	uModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/user"
	auRepo "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/repository/auth"
)

type fAuthRepository struct {
	fAuth *auth.Client
}

// GetAuthUserFull implements answer.Repository
func (r fAuthRepository) GetAuthUserFull(ctx context.Context, id string) (uModel.AuthUserFull, error) {
	ur, err := r.fAuth.GetUser(ctx, id)
	if err != nil {
		return uModel.AuthUserFull{}, err
	}
	return uModel.AuthUserFull{
		IsEmailVerified: ur.EmailVerified,
	}, nil
}

func NewFirebaseAuthRepository(fAuth *auth.Client) auRepo.Repository {
	return fAuthRepository{fAuth: fAuth}
}
