package pg

import (
	"context"
	errorCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/error"
	"github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/sqlc"
	aModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/article"
	uModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/user"
	auRepo "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/repository/auth"
	"github.com/jackc/pgx/v4"
)

type pgArticleRepository struct {
	querier sqlc.Querier
	auRepo  auRepo.Repository
}

var (
	ErrUser_UserNotAuthorized = errorCommon.NewNotFoundError("USER_NOT_AUTHORIZED")
)

func (r pgArticleRepository) UpdateArticle(ctx context.Context, arg aModel.UpdateArticle, au uModel.AuthUser) (string, error) {
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

func (r pgArticleRepository) CreateArticle(ctx context.Context, arg aModel.AddArticle, au uModel.AuthUser) (string, error) {
	id, err := r.querier.CreateArticle(ctx, sqlc.CreateArticleParams(arg))
	if err == pgx.ErrNoRows {
		return "", errorCommon.NewNotFoundError("Article not found")
	}
	return id, err
}

func (r pgArticleRepository) GetArticle(ctx context.Context, id string, au uModel.AuthUser) (aModel.Article, error) {
	//if !au.IsSame(id) {
	//	return aModel.Article{}, ErrUser_UserNotAuthorized
	//}

	art, err := r.querier.GetArticle(ctx, id)
	if err == pgx.ErrNoRows {
		return aModel.Article{}, errorCommon.NewNotFoundError("Article not found")
	}
	return aModel.Article(art), err
}

func (r pgArticleRepository) GetAllArticles(ctx context.Context, au uModel.AuthUser) ([]aModel.Article, error) {
	art, err := r.querier.GetAllArticle(ctx)
	if err == pgx.ErrNoRows {
		return []aModel.Article{}, errorCommon.NewNotFoundError("Article not found")
	}

	var allArt []aModel.Article

	for i := 0; i < len(art); i++ {
		allArt = append(allArt, aModel.Article{
			ID:           art[i].ID,
			AuthorID:     art[i].AuthorID,
			JudulArticle: art[i].JudulArticle,
			IsiArticle:   art[i].IsiArticle,
			Author:       art[i].Author,
			ImageUrl:     art[i].ImageUrl,
			CreatedAt:    art[i].CreatedAt,
			UpdatedAt:    art[i].UpdatedAt,
		})
	}

	return allArt, err
}

func (r pgArticleRepository) DeleteArticle(ctx context.Context, id string, au uModel.AuthUser) error {
	err := r.querier.DeleteArticle(ctx, id)
	if err == pgx.ErrNoRows {
		return errorCommon.NewNotFoundError("Article not found")
	}
	return err
}

func NewPGArticleRepository(querier sqlc.Querier) pgArticleRepository {
	return pgArticleRepository{querier: querier}
}
