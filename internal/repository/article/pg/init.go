package pg

import (
	"context"
	errorCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/error"
	"github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/sqlc"
	aModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/article"
	auRepo "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/repository/auth"
	"github.com/jackc/pgx/v4"
)

type pgArticleRepository struct {
	querier sqlc.Querier
	auRepo  auRepo.Repository
}

func (r pgArticleRepository) UpdateArticle(ctx context.Context, arg aModel.UpdateArticle, au aModel.AuthArticle) (string, error) {
	id, err := r.querier.UpdateArticle(ctx, sqlc.UpdateArticleParams{
		ID:           arg.ID,
		AuthorID:     arg.AuthorID,
		JudulArticle: arg.JudulArticle,
		IsiArticle:   arg.IsiArticle,
		Author:       arg.Author,
		ImageUrl:     arg.ImageUrl,
	})
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("Article not found")
	}
	return id, err
}

func (r pgArticleRepository) CreateArticle(ctx context.Context, arg aModel.AddArticle, au aModel.AuthArticle) (string, error) {
	id, err := r.querier.CreateArticle(ctx, sqlc.CreateArticleParams(arg))
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("Article not found")
	}
	return id, err
}

func (r pgArticleRepository) GetAllArticles(ctx context.Context, au aModel.AuthArticle) ([]aModel.Article, error) {
	art, err := r.querier.GetAllArticles(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]aModel.Article, len(art))
	for i, article := range art {
		result[i] = aModel.Article(article)
	}
	return result, nil
}

func (r pgArticleRepository) GetArticle(ctx context.Context, id string, au aModel.AuthArticle) (aModel.Article, error) {
	art, err := r.querier.GetArticle(ctx, id)
	if err == pgx.ErrNoRows {
		return aModel.Article{}, errorCommon.NewNotFoundError("Article not found")
	}
	return aModel.Article(art), err
}

func (r pgArticleRepository) DeleteArticle(ctx context.Context, id string, au aModel.AuthArticle) error {
	err := r.querier.DeleteArticle(ctx, id)
	if err == pgx.ErrNoRows {
		return errorCommon.NewNotFoundError("Article not found")
	}
	return err
}

func NewPGArticleRepository(querier sqlc.Querier) pgArticleRepository {
	return pgArticleRepository{querier: querier}
}
