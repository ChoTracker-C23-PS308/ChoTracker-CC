package http

import "time"

type (
	Error struct {
		Message string            `json:"message"`
		Errors  map[string]string `json:"errors"`
	}

	Response struct {
		Data    any    `json:"data"`
		Message string `json:"message"`
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
		Email     string `json:"email" binding:"required"`
		Name      string `json:"name" binding:"required"`
		BirthDate string `json:"birth_date" binding:"required"`
		Gender    string `json:"gender" binding:"required"`
		ImageUrl  string `json:"image_url" binding:"required"`
	}

	UpdateUser struct {
		ID        string `json:"id" binding:"required"`
		Email     string `json:"email" binding:"required"`
		Name      string `json:"name" binding:"required"`
		BirthDate string `json:"birth_date" binding:"required"`
		Gender    string `json:"gender" binding:"required"`
		ImageUrl  string `json:"image_url" binding:"required"`
	}

	AddArticle struct {
		AuthorID     string    `json:"author_id" binding:"required"`
		JudulArtikel string    `json:"judul_artikel" binding:"required"`
		IsiArtikel   string    `json:"isi_artikel" binding:"required"`
		Author       string    `json:"author" binding:"required"`
		ImageURL     string    `json:"image_url"  binding:"required"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	}

	AddHistory struct {
		Uid            string    `json:"uid" binding:"required"`
		TotalKolestrol float64   `json:"total_kolestrol" binding:"required"`
		Tingkat        string    `json:"tingkat" binding:"required"`
		ImageUrl       string    `json:"image_url" binding:"required"`
		CreatedAt      time.Time `json:"created_at"`
		UpdateAt       time.Time `json:"updated_at"`
	}

	GetAllHistory struct {
		ID             string    `json:"id"`
		UID            string    `json:"uid"`
		TotalKolestrol string    `json:"total_kolestrol"`
		Tingkat        string    `json:"tingkat"`
		ImageURL       string    `json:"image_url"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	UpdateArticle struct {
		ID           string    `json:"id" binding:"required"`
		AuthorID     string    `json:"author_id" binding:"required"`
		JudulArtikel string    `json:"judul_artikel" binding:"required"`
		IsiArtikel   string    `json:"isi_artikel" binding:"required"`
		Author       string    `json:"author" binding:"required"`
		ImageURL     string    `json:"image_url" binding:"required"`
		UpdatedAt    time.Time `json:"updated_at"`
	}
)
