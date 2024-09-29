package module

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
	router.GET("/modules", h.getModules)
	router.POST("/modules", h.initModule)
}

func (h Handler) getModules(c *gin.Context) {

}

func (h Handler) initModule(router *gin.Context) {

}
