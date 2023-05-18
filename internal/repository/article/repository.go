package article

import (
	"context"
	aModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/article"
)

type Repository interface {
	CreateArticle(ctx context.Context, arg aModel.AddArticle, au aModel.AuthArticle) (string, error)
	GetArticle(ctx context.Context, id string, au aModel.AuthArticle) (aModel.Article, error)
	UpdateArticle(ctx context.Context, arg aModel.UpdateArticle, au aModel.AuthArticle) (string, error)
}
