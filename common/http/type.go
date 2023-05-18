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
	Driver struct {
		Name         string `json:"name"`
		PoliceNumber string `json:"police_number"`
		VehicleModel string `json:"vehicle_model"`
		VehicleType  string `json:"vehicle_type"`
	}
	Inquiry struct {
		Price    int64  `json:"price"`
		Distance int32  `json:"distance"`
		Duration int32  `json:"duration"`
		OAddress string `json:"o_address"`
		DAddress string `json:"d_address"`
	}
	Payment struct {
		Amount   float64 `json:"amount"`
		Status   string  `json:"status"`
		Method   string  `json:"method"`
		QrString string  `json:"qr_string"`
	}
	getOrder struct {
		ID           string    `json:"id"`
		UName        string    `json:"u_name"`
		Driver       Driver    `json:"driver"`
		OrderInquiry Inquiry   `json:"order_inquiry"`
		Payment      Payment   `json:"payment"`
		Status       string    `json:"status"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)
