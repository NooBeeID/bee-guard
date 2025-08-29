package login

import (
	"github.com/NooBeeID/bee-guard/infra/contracts"
	"github.com/NooBeeID/bee-guard/infra/modules"
	"github.com/NooBeeID/bee-guard/infra/router"
)

type Login struct {
	// contract *loginContract
	handler contracts.Handler
}

// Run implements modules.Modules.
func (l *Login) Run(router *router.Router) {
	router.Post("/auth/login", l.handler.Handle)
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
		handler: cfg.Handler,
	}
}
