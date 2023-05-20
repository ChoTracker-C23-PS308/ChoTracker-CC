package http

import (
	httpCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/http"
	aModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/article"
	uModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (d HTTPArtikelDelivery) deleteArticle(c *gin.Context) {
	context := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	id := c.Param("id")

	err := d.articleRepo.DeleteArticle(context, id, aModel.AuthArticle(au))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, httpCommon.Response{
		Message: "Article successfully deleted",
	})
}

func (d HTTPArtikelDelivery) updateArticle(c *gin.Context) {
	context := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	id := c.Param("id")

	var article httpCommon.UpdateArticle
	if err := c.ShouldBindJSON(&article); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	nid, err := d.articleRepo.UpdateArticle(context, aModel.UpdateArticle{
		ID:           id,
		AuthorID:     article.AuthorID,
		JudulArticle: article.JudulArtikel,
		IsiArticle:   article.IsiArtikel,
		Author:       article.Author,
		ImageUrl:     article.ImageURL,
		UpdatedAt:    article.UpdatedAt,
	}, aModel.AuthArticle(au))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, httpCommon.Response{
		Data:    nid,
		Message: "Data successfully updated",
	})
}

func (d HTTPArtikelDelivery) addArticle(c *gin.Context) {
	context := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

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
	}, aModel.AuthArticle(au))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, httpCommon.Response{
		Data:    nid,
		Message: "Data successfully created",
	})
}

func (d HTTPArtikelDelivery) getAllArticles(c *gin.Context) {
	context := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	u, err := d.articleRepo.GetAllArticles(context, aModel.AuthArticle(au))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, httpCommon.Response{Data: u})
}

func (d HTTPArtikelDelivery) getArticle(c *gin.Context) {
	context := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	id := c.Param("id")

	u, err := d.articleRepo.GetArticle(context, id, aModel.AuthArticle(au))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, httpCommon.Response{Data: u})
}
