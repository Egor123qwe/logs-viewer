package http

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Egor123qwe/logs-viewer/internal/server/launcher"
	"github.com/Egor123qwe/logs-viewer/internal/util/CORS"
)

type server struct {
	srv    *http.Server
	config Config
}

func New(router http.Handler, config Config) launcher.Server {
	return &server{
		srv: &http.Server{
			Addr:        fmt.Sprintf(":%d", config.Port),
			Handler:     CORS.New(router),
			ReadTimeout: config.ReadTime,
		},
		config: config,
	}
}

func (s *server) Serve(ctx context.Context) error {
	errCh := make(chan error)

	go func() {
		errCh <- s.srv.ListenAndServe()
	}()

	log.Printf("http-server: serving on http://localhost%s\n", s.srv.Addr)

	select {
	case err := <-errCh:
		return fmt.Errorf("http-server: %w", err)

	case <-ctx.Done():
		ctx, cancel := context.WithTimeout(context.Background(), s.config.ShutdownTime)
		defer cancel()

		if err := s.srv.Shutdown(ctx); err != nil {
			log.Println("http-server: Shutdown error: " + err.Error())
		}

		log.Println("http-server: server stopped successfully.")
	}

	return nil
}
