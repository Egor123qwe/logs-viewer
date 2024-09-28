// Package app handles starting and monitoring the server for graceful shutdown.
package app

import (
	"context"
	"fmt"

	"github.com/Egor123qwe/logs-viewer/api"
	"github.com/Egor123qwe/logs-viewer/internal/server"
	"github.com/Egor123qwe/logs-viewer/internal/service"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("app")

type App struct {
	srv service.Service
}

func New() (*App, error) {
	api, err := api.New()
	if err != nil {
		return nil, err
	}

	app := &App{
		srv: service.New(api),
	}

	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	srv, err := server.New(a.srv)
	if err != nil {
		return err
	}

	if err := srv.Serve(ctx); err != nil {
		return fmt.Errorf("server stopped with error: %w\n", err)
	}

	log.Infof("server stopped")

	return nil
}
