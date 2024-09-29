package auth

import (
	"context"

	"github.com/Egor123qwe/logs-viewer/internal/model"
	"github.com/Egor123qwe/logs-viewer/internal/model/auth"
)

type Service interface {
	Auth(ctx context.Context, req auth.Credentials) error
}

type service struct {
	config auth.Credentials
}

func New() Service {
	return &service{
		config: newConfig(),
	}
}

func (s service) Auth(ctx context.Context, req auth.Credentials) error {
	if s.config.Username != req.Username || s.config.Password != req.Password {
		return model.InvalidCredentialsErr
	}

	return nil
}
