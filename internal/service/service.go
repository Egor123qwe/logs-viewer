package service

import (
	"github.com/Egor123qwe/logs-viewer/api"
	"github.com/Egor123qwe/logs-viewer/internal/service/log"
)

type Service interface {
	Logs() log.Service
}

type service struct {
	log log.Service
}

func New(api api.Service) Service {
	return &service{
		log: log.New(api),
	}
}

func (s *service) Logs() log.Service {
	return s.log
}
