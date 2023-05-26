package article

import (
	"context"
	aModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/article"
	uModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/user"
)

type Repository interface {
	CreateArticle(ctx context.Context, arg aModel.AddArticle, au uModel.AuthUser) (string, error)
	GetArticle(ctx context.Context, id string, au uModel.AuthUser) (aModel.Article, error)
	GetAllArticles(ctx context.Context, au uModel.AuthUser) ([]aModel.Article, error)
	UpdateArticle(ctx context.Context, arg aModel.UpdateArticle, au uModel.AuthUser) (string, error)
	DeleteArticle(ctx context.Context, id string, au uModel.AuthUser) error
}
