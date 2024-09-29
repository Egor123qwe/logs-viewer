package module

import (
	"net/http"

	"github.com/Egor123qwe/logs-viewer/internal/handler/model"
	"github.com/Egor123qwe/logs-viewer/internal/handler/model/module"
	srvmodel "github.com/Egor123qwe/logs-viewer/internal/model/log"
	modulesrv "github.com/Egor123qwe/logs-viewer/internal/service/log"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	srv modulesrv.Service
}

func New(service modulesrv.Service) Handler {
	return Handler{
		srv: service,
	}
}

func (h Handler) Init(router *gin.RouterGroup) {
	router.GET("/modules", h.getModules)
	router.GET("/init", h.initModule)
}

func (h Handler) getModules(c *gin.Context) {
	filter := c.Query("filter")

	modules, err := h.srv.GetModules(c, srvmodel.ModuleReq{NameFilter: filter})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResp{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, modules)
}

func (h Handler) initModule(c *gin.Context) {
	name := c.Query("module")

	if name == "" {
		c.JSON(http.StatusBadRequest, model.ErrorResp{Error: model.EmptyModuleNameErr.Error()})
		return
	}

	id, err := h.srv.InitModule(c, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.ErrorResp{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, module.InitResp{ID: id})
}
