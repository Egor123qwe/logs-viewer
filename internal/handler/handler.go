package handler

import (
	"net/http"

	httpHandler "github.com/Egor123qwe/logs-viewer/internal/handler/http"
	"github.com/Egor123qwe/logs-viewer/internal/service"
)

type Handler struct {
	HTTP http.Handler
}

func New(srv service.Service) Handler {
	handler := Handler{
		HTTP: httpHandler.New(srv),
	}

	return handler
}
