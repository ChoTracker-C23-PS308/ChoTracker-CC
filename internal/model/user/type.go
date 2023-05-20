package user

import "time"

type (
	AuthUser struct {
		ID   string
		Role uint32
	}
	AuthUserFull struct {
		IsEmailVerified bool
	}
	User struct {
		ID        string
		Name      string
		Email     string
		BirthDate string
		Gender    string
		ImageUrl  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	AddUser struct {
		ID        string
		Name      string
		Email     string
		BirthDate string
		Gender    string
		ImageUrl  string
	}
	UpdateUser struct {
		ID        string
		Name      string
		Email     string
		BirthDate string
		Gender    string
		ImageUrl  string
	}
	AddImage struct {
	}
)
