package api

import (
	"fmt"

	"github.com/Egor123qwe/logs-viewer/api/log"
)

type Service interface {
	Log() *log.Client
}

type service struct {
	log *log.Client
}

func New() (Service, error) {
	config := newConfig()

	logClient, err := log.New(config.logs.host, config.logs.port)
	if err != nil {
		return nil, fmt.Errorf("failed to create [log-storage] client: %w", err)
	}

	srv := service{
		log: logClient,
	}

	return srv, nil
}

func (s service) Log() *log.Client {
	return s.log
}
