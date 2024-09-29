package log

import (
	"net/http"

	"github.com/Egor123qwe/logs-viewer/internal/handler/model"
	"github.com/Egor123qwe/logs-viewer/internal/handler/model/log"
	logmodel "github.com/Egor123qwe/logs-viewer/internal/model/log"
	logsrv "github.com/Egor123qwe/logs-viewer/internal/service/log"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	srv logsrv.Service
}

func New(service logsrv.Service) Handler {
	return Handler{
		srv: service,
	}
}

func (h Handler) Init(router *gin.RouterGroup) {
	router.GET("/logs", h.getLogs)
}

func (h Handler) getLogs(c *gin.Context) {
	var reqFilter log.Filter

	if err := c.BindJSON(&reqFilter); err != nil {
		c.JSON(http.StatusBadRequest, model.ErrorResp{Error: err.Error()})
		return
	}

	filter := logmodel.LogFilter{
		TraceID:     reqFilter.TraceID,
		ModuleID:    reqFilter.ModuleID,
		Message:     reqFilter.Message,
		StartTime:   reqFilter.StartTime,
		EndTime:     reqFilter.EndTime,
		CountOnPage: reqFilter.CountOnPage,
		Page:        reqFilter.Page,
	}

	if reqFilter.Level != nil {
		level := logmodel.ConvertLevelName(*reqFilter.Level)
		if level == logmodel.Invalid {
			c.JSON(http.StatusBadRequest, model.ErrorResp{Error: model.InvalidLevelErr.Error()})
			return
		}

		filter.Level = &level
	}

	resp, err := h.srv.GetLogs(c, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResp{Error: err.Error()})
		return
	}

	result := log.GetResp{
		PagesCount: resp.Total,
	}

	for _, l := range resp.Logs {
		result.Logs = append(result.Logs, log.Log{
			ID:      l.ID,
			TraceID: l.TraceID,
			Module:  l.Module,
			Time:    l.Time,
			Level:   l.Level,
			Message: l.Message,
		})
	}

	c.JSON(http.StatusOK, result)
}
