package service

import (
	"github.com/Egor123qwe/logs-viewer/api"
	"github.com/Egor123qwe/logs-viewer/internal/service/auth"
	"github.com/Egor123qwe/logs-viewer/internal/service/log"
)

type Service interface {
	Logs() log.Service
	Auth() auth.Service
}

type service struct {
	log  log.Service
	auth auth.Service
}

func New(api api.Service) Service {
	return &service{
		log:  log.New(api),
		auth: auth.New(),
	}
}

func (s *service) Logs() log.Service {
	return s.log
}

func (s *service) Auth() auth.Service {
	return s.auth
}
