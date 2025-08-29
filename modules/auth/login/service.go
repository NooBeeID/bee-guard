package login

import (
	"context"
	"log/slog"

	"github.com/NooBeeID/bee-guard/entity"
)

type contractDBRepository interface {
	GetAuthByEmail(ctx context.Context, email string) (auth entity.Auth, err error)
}

type contractCacheRepository interface {
	StoreSession(ctx context.Context, sessionId string, value string) (err error)
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

	token, err :=  auth.GenerateSession();
	if err != nil {
		slog.ErrorContext(ctx, "Error GenerateToken", slog.Any("error", err.Error()), slog.String("email", req.Email))
		return Response{}, err
	}

	if s.useCache() {
		if err := s.cache.StoreSession(ctx, "session_key_user", token.Token); err != nil {
			slog.ErrorContext(ctx, "Error StoreSession", slog.Any("error", err.Error()), slog.String("email", req.Email))
			return Response{}, err
		}
	}

	return Response{
		Token: token.Token,
		Type: "Bearer",
	}, nil
}

func NewService(repo contractDBRepository, cache contractCacheRepository) service {
	return service{repo: repo, cache: cache}
}

func (s service) useCache() bool {
	return s.cache != nil
}
