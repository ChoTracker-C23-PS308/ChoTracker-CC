package http

import "time"

type (
	Error struct {
		Message string            `json:"message"`
		Errors  map[string]string `json:"errors"`
	}

	Response struct {
		Data any `json:"data"`
	}

	User struct {
		ID          string    `json:"id"`
		Name        string    `json:"name"`
		Email       string    `json:"email"`
		PhoneNumber string    `json:"phone_number"`
		Nim         string    `json:"nim"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
	AddUser struct {
		ID        string `json:"id" binding:"required"`
		Email     string `json:"email" binding:"required,email,contains=@student.unsri.ac.id"`
		Name      string `json:"name" binding:"required"`
		BirthDate string `json:"birth_date" binding:"required"`
		Gender    string `json:"gender" binding:"required"`
		ImageUrl  string `json:"image_url" binding:"required"`
	}
	UpdateUser struct {
		ID        string `json:"id" binding:"required"`
		Email     string `json:"email" binding:"required,email,contains=@student.unsri.ac.id"`
		Name      string `json:"name" binding:"required"`
		BirthDate string `json:"birth_date" binding:"required"`
		Gender    string `json:"gender" binding:"required"`
		ImageUrl  string `json:"image_url" binding:"required"`
	}
)
