package http

import (
	"net/http"

	"github.com/Egor123qwe/logs-viewer/internal/handler/http/auth"
	"github.com/Egor123qwe/logs-viewer/internal/handler/http/log"
	"github.com/Egor123qwe/logs-viewer/internal/handler/http/module"
	"github.com/Egor123qwe/logs-viewer/internal/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	auth auth.Service
}

func New(srv service.Service) http.Handler {
	router := gin.Default()

	h := &handler{
		auth: auth.New(router.Group("/auth")),
	}

	api := router.Group("/api")
	api.Use(h.auth.Auth)

	log.New(srv).Init(api.Group("/log"))
	module.New(srv).Init(api.Group("/module"))

	return router
}
