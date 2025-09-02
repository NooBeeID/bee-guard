package register

import (
	"github.com/NooBeeID/bee-guard/infra/contracts"
	"github.com/NooBeeID/bee-guard/infra/modules"
	"github.com/NooBeeID/bee-guard/infra/router"
)

type Register struct {
	// contract *loginContract
	handler contracts.Handler
	path    string
}

// Run implements modules.Modules.
func (l *Register) Run(router *router.Router) {
	router.Post(l.getPath(), l.handler.Handle)
}

// GetName implements modules.Modules.
func (l *Register) GetName() string {
	return "register"
}

func (l *Register) getPath() string {
	if l.path == "" {
		return "/auth/register"
	}
	return l.path
}

func New(cfg *modules.ConfigService) *Register {
	if cfg == nil {
		return &Register{
			path: "/auth/register",
		}
	}

	return &Register{
		handler: cfg.Handler,
	}
}
