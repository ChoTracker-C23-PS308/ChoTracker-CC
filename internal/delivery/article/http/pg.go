package http

import (
	httpCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/http"
	aModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/article"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (d HTTPArtikelDelivery) updateArticle(c *gin.Context) {
	context := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(aModel.AuthArticle)

	id := c.Param("id")

	var article httpCommon.UpdateArticle
	if err := c.ShouldBindJSON(&article); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	nid, err := d.articleRepo.CreateArticle(context, aModel.AddArticle{
		ID:           id,
		AuthorID:     article.AuthorID,
		JudulArticle: article.JudulArtikel,
		IsiArticle:   article.IsiArtikel,
		Author:       article.Author,
		ImageUrl:     article.ImageURL,
		UpdatedAt:    article.UpdatedAt,
	}, au)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, httpCommon.Response{
		Data: nid,
	})
}

func (d HTTPArtikelDelivery) getArticle(c *gin.Context) {
	context := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(aModel.AuthArticle)

	id := c.Param("id")

	u, err := d.articleRepo.GetArticle(context, id, au)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, httpCommon.Response{Data: u})
}

func (d HTTPArtikelDelivery) addArticle(c *gin.Context) {
	context := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(aModel.AuthArticle)

	var article httpCommon.AddArticle
	if err := c.ShouldBindJSON(&article); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	nid, err := d.articleRepo.CreateArticle(context, aModel.AddArticle{
		ID:           article.ID,
		AuthorID:     article.AuthorID,
		JudulArticle: article.JudulArtikel,
		IsiArticle:   article.IsiArtikel,
		Author:       article.Author,
		ImageUrl:     article.ImageURL,
		CreatedAt:    article.CreatedAt,
		UpdatedAt:    article.UpdatedAt,
	}, au)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, httpCommon.Response{
		Data: nid,
	})
}
