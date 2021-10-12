package oauth2

import (
	"context"

	"github.com/penkovski/oauth2min/internal/server"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) Token(ctx context.Context, req *server.TokenRequest) (*server.TokenResponse, error) {
	return nil, nil
}
