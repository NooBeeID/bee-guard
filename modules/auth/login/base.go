package login

import (
	"database/sql"

	"github.com/NooBeeID/bee-guard/infra/contracts"
	"github.com/NooBeeID/bee-guard/infra/modules"
	"github.com/NooBeeID/bee-guard/infra/router"
)

type Login struct {
	beeRouter *router.Router
	db        *sql.DB
	cache     contracts.Cache

	// contract *loginContract
	handler contracts.Handler
}

// Run implements modules.Modules.
func (l *Login) Run() {
	l.beeRouter.Post("/auth/login", l.handler.Handle)
}

// GetName implements modules.Modules.
func (l *Login) GetName() string {
	return "login"
}

func New(cfg *modules.ConfigService) *Login {
	if cfg == nil {
		return &Login{}
	}


	return &Login{
		beeRouter: cfg.Router,
		db:        cfg.DB,
		handler: cfg.Handler,
	}
}