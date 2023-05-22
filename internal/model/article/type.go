package article

import "time"

type (
	AuthArticle struct {
		ID   string
		Role uint32
	}

	Article struct {
		ID           string
		AuthorID     string
		JudulArticle string
		IsiArticle   string
		Author       string
		ImageUrl     string
		CreatedAt    time.Time
		UpdatedAt    time.Time
	}

	AddArticle struct {
		ID           string
		AuthorID     string
		JudulArticle string
		IsiArticle   string
		Author       string
		ImageUrl     string
	}

	UpdateArticle struct {
		ID           string
		AuthorID     string
		JudulArticle string
		IsiArticle   string
		Author       string
		ImageUrl     string
		UpdatedAt    time.Time
	}
)
