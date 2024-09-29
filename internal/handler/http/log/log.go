package log

import (
	"github.com/Egor123qwe/logs-viewer/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	srv service.Service
}

func New(service service.Service) Handler {
	return Handler{
		srv: service,
	}
}

func (h Handler) Init(router *gin.RouterGroup) {
	router.GET("/logs", h.getLogs)
}

func (h Handler) getLogs(c *gin.Context) {

}
