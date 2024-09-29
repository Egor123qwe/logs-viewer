package handler

import (
	"net/http"

	"github.com/Egor123qwe/logs-viewer/internal/service"
)

type Handler struct {
	HTTP http.Handler
}

func New(srv service.Service) Handler {
	handler := Handler{}

	return handler
}
