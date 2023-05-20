package http

import (
	httpCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/http"
	hModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/history"
	uModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (d HTTPHistoryDelivery) getHistories(c *gin.Context) {
	context := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	uid := c.Param("uid")

	u, err := d.historyRepo.GetHistory(context, uid, hModel.AuthHistory(au))
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, httpCommon.Response{Data: u})
}

func (d HTTPHistoryDelivery) addHistory(c *gin.Context) {
	context := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	var history httpCommon.AddHistory
	if err := c.ShouldBindJSON(&history); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	nid, err := d.historyRepo.CreateHistory(context, hModel.AddHistory{
		ID:             history.ID,
		Uid:            history.Uid,
		TotalKolestrol: history.TotalKolestrol,
		Tingkat:        history.Tingkat,
		ImageUrl:       history.ImageUrl,
	}, hModel.AuthHistory(au))

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusCreated, httpCommon.Response{
		Data:    nid,
		Message: "Data successfully created",
	})
}

func (d HTTPHistoryDelivery) deleteHistory(c *gin.Context) {
	context := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	id := c.Param("id")

	err := d.historyRepo.DeleteHistory(context, id, hModel.AuthHistory(au))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, httpCommon.Response{
		Message: "History successfully deleted",
	})
}
