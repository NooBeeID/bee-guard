package register

import (
	"context"
	"log/slog"

	"github.com/NooBeeID/bee-guard/entity"
)

type contractDBRepository interface {
	GetAuthByEmail(ctx context.Context, email string) (auth entity.Auth, err error)
	StoreAuth(ctx context.Context, auth entity.Auth) (err error)
}

type contractCacheRepository interface {
	StoreSession(ctx context.Context, sessionId string, value string) (err error)
}

type service struct {
	repo  contractDBRepository
	cache contractCacheRepository
}

// Register implements ContractService.
func (s service) Register(ctx context.Context, req Request) (res Response, err error) {
	auth, err := s.repo.GetAuthByEmail(ctx, req.Email)
	if err != nil {
		slog.ErrorContext(ctx, "Error GetAuthByEmail", slog.Any("error", err.Error()), slog.String("email", req.Email))
		return
	}

	if auth.IsExists() {
		err = entity.ErrAuthAlreadyExists
		slog.ErrorContext(ctx, "Error GetAuthByEmail", slog.Any("error", err.Error()), slog.String("email", req.Email))
		return
	}

	auth = req.toAuth()
	if err = auth.GeneratePassword(); err != nil {
		slog.ErrorContext(ctx, "Error GeneratePassword", slog.Any("error", err.Error()), slog.String("email", req.Email))
		return
	}

	if err = s.repo.StoreAuth(ctx, auth); err != nil {
		slog.ErrorContext(ctx, "Error StoreAuth", slog.Any("error", err.Error()), slog.String("email", req.Email))
		return
	}

	if s.useCache() {
		if err = s.cache.StoreSession(ctx, "session_key_user", auth.ID); err != nil {
			slog.ErrorContext(ctx, "Error StoreSession", slog.Any("error", err.Error()), slog.String("email", req.Email))
			return
		}
	}

	return
}

func NewService(repo contractDBRepository, cache contractCacheRepository) service {
	return service{repo: repo, cache: cache}
}

func (s service) useCache() bool {
	return s.cache != nil
}
