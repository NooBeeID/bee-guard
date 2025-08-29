package login

import (
	"context"
	"log/slog"

	"github.com/NooBeeID/bee-guard/modules/auth/login/adapter"
)

type contractDBRepository interface {
	GetAuthByEmail(ctx context.Context, email string) (auth adapter.Auth, err error)
}

type contractCacheRepository interface {
}

type service struct {
	repo  contractDBRepository
	cache contractCacheRepository
}

// Login implements contractService.
func (s service) Login(ctx context.Context, req Request) (Response, error) {
	auth, err := s.repo.GetAuthByEmail(ctx, req.Email)
	if err != nil {
		slog.ErrorContext(ctx, "Error GetAuthByEmail", slog.Any("error", err.Error()), slog.String("email", req.Email))
		return Response{}, err
	}

	if err := auth.VerifyPassword(req.Password); err != nil {
		slog.ErrorContext(ctx, "Error VerifyPassword", slog.Any("error", err.Error()), slog.String("email", req.Email))
		return Response{}, err
	}

	return Response{}, nil
}

func NewService(repo contractDBRepository, cache contractCacheRepository) service {
	return service{repo: repo, cache: cache}
}

func (s service) UseCache() bool {
	return s.cache != nil
}
