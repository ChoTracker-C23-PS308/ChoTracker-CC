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

	id, err := r.querier.CreateUser(ctx, sqlc.CreateUserParams(arg))
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("User not found")
	}
	return id, err
}

// // DeleteUser implements user.User
//
//	func (r pgUserRepository) DeleteUser(ctx context.Context, id string) error {
//		err := r.querier.DeleteUser(ctx, id)
//		if err == pgx.ErrNoRows {
//			return errorCommon.NewNotFoundError("User not found")
//		}
//		return err
//	}
//
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

// // ListUsers implements user.User
//
//	func (r pgUserRepository) ListUsers(ctx context.Context) ([]uModel.User, error) {
//		us, err := r.querier.ListUsers(ctx)
//		ums := make([]uModel.User, 0)
//		for _, u := range us {
//			ums = append(ums, uModel.User(u))
//		}
//		return ums, err
//	}
//
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
		ImageUrl:  arg.ImageUrl,
	})
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("User not found")
	}
	return id, err
}

func NewPGUserRepository(querier sqlc.Querier) pgUserRepository {
	return pgUserRepository{querier: querier}
}

//// GetUserHistory implements BLM
//func (r pgUserRepository) GetUserHistory(ctx context.Context, id string) ([]oModel.Order, error) {
//	o, err := r.querier.GetOrderHistory(ctx, id)
//	if err == pgx.ErrNoRows {
//		var temp []oModel.Order
//		return temp, errorCommon.NewNotFoundError("User not found")
//	}
//
//	var oh []oModel.Order
//
//	for i := 0; i < len(o); i++ {
//		d, err := r.querier.GetUser(ctx, o[i].UserID_2)
//		if err == pgx.ErrNoRows {
//			var temp []oModel.Order
//			return temp, errorCommon.NewNotFoundError("Driver not found")
//		}
//
//		oh = append(oh, oModel.Order{
//			ID:    o[i].ID,
//			DName: d.Name,
//			User: uModel.User{
//				ID:   o[i].UserID,
//				Name: o[i].Name,
//			},
//			Driver: dModel.Driver{
//				ID:           o[i].DriverID,
//				PoliceNumber: o[i].PoliceNumber,
//				VehicleModel: o[i].VehicleModel,
//				VehicleType:  o[i].VehicleType,
//			},
//			OrderInquiry: oModel.OrderInquiry{
//				ID:       o[i].OrderInquiryID,
//				Price:    o[i].Price,
//				Distance: o[i].Distance,
//				Duration: o[i].Duration,
//				Origin: oModel.Location{
//					Address: o[i].OriginAddress,
//				},
//				Destination: oModel.Location{
//					Address: o[i].DestinationAddress,
//				},
//				Routes: o[i].Routes,
//			},
//			Payment: pModel.Payment{
//				ID:       o[i].PaymentID,
//				Amount:   o[i].Amount,
//				Status:   pModel.Status(o[i].Status_2),
//				Method:   pModel.Method(o[i].Method),
//				QrString: o[i].QrStr,
//			},
//			Status:    oModel.Status(o[i].Status),
//			CreatedAt: o[i].CreatedAt,
//			UpdatedAt: o[i].UpdatedAt,
//		})
//	}
//
//	return []oModel.Order(oh), err
//
//	//{
//	//	"id" : 1,
//	//	"address" : "Gang Buntu",
//	//	"status" : "Sedang diperjalanan",
//	//	"update_at" : "2023-02-10T13:45:00.000Z"
//	//
//	//}
//}
