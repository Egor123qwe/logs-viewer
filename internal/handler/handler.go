package handler

import (
	"github.com/Egor123qwe/logs-viewer/internal/service"
)

type Handler struct {
}

func New(srv service.Service) Handler {
	handler := Handler{}

	return handler
}
