// Package server configures and starts servers for handling incoming requests.
package server

import (
	"context"
	"sync"

	"github.com/Egor123qwe/logs-viewer/internal/server/launcher"
	"github.com/Egor123qwe/logs-viewer/internal/service"
	"github.com/op/go-logging"
	"golang.org/x/sync/errgroup"
)

var log = logging.MustGetLogger("server")

type server struct {
	servers []launcher.Server
}

func New(srv service.Service) (launcher.Server, error) {
	//h := handler.New(srv)

	result := &server{
		servers: []launcher.Server{},
	}

	return result, nil
}

func (s *server) Serve(ctx context.Context) error {
	gr, grCtx := errgroup.WithContext(ctx)

	// start server
	gr.Go(func() error {
		return s.serve(grCtx)
	})

	var err error

	if err = gr.Wait(); err != nil {
		log.Criticalf("app error: %v", err)
	}

	log.Infof("app: shutting down the server...")

	return err
}

func (s *server) serve(ctx context.Context) error {
	var wg sync.WaitGroup
	wg.Add(len(s.servers))

	gr, grCtx := errgroup.WithContext(ctx)

	for _, s := range s.servers {
		s := s

		gr.Go(func() error {
			defer wg.Done()

			return s.Serve(grCtx)
		})
	}

	wg.Wait()

	return gr.Wait()
}
