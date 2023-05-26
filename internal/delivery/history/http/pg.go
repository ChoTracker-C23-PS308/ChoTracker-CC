package http

import (
	httpCommon "github.com/ChoTracker-C23-PS308/ChoTracker-CC/common/http"
	hModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/history"
	uModel "github.com/ChoTracker-C23-PS308/ChoTracker-CC/internal/model/user"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func (d HTTPHistoryDelivery) getHistories(c *gin.Context) {
	context := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	uid := c.Param("uid")

	u, err := d.historyRepo.GetHistory(context, uid, au)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, httpCommon.Response{
		Data:    u,
		Message: "Get Histoies Data successfully",
	})
}

func (d HTTPHistoryDelivery) addHistory(c *gin.Context) {
	context := c.Request.Context()
	au := c.MustGet(httpCommon.AUTH_USER).(uModel.AuthUser)

	// Generate id history
	hid, err := uuid.NewRandom()

	if err != nil {
		c.Error(err)
	}

	var history httpCommon.AddHistory
	if err := c.ShouldBindJSON(&history); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}

	nid, err := d.historyRepo.CreateHistory(context, hModel.AddHistory{
		ID:             hid.String(),
		Uid:            history.Uid,
		TotalKolestrol: history.TotalKolestrol,
		Tingkat:        history.Tingkat,
		ImageUrl:       history.ImageUrl,
	}, au)

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

	uid := c.Param("uid")
	id := c.Param("id")

	err := d.historyRepo.DeleteHistory(context, uid, id, au)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, httpCommon.Response{Data: id,
		Message: "History successfully deleted",
	})
}
