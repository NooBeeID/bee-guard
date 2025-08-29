package modules

import (
	"fmt"
	"log/slog"

	"github.com/NooBeeID/bee-guard/infra/router"
)

type Modules interface {
	Run()
	GetName() string
}

type Base struct {
	router  *router.Router
	modules []Modules
}

func NewModules(router *router.Router, modules ...Modules) Base {
	return Base{modules: modules, router: router}
}

func (b Base) GetModules() []Modules {
	return b.modules
}

func (b Base) Run() error {
	for _, module := range b.modules {
		slog.Debug(fmt.Sprintf("running module %v", module.GetName()))
		module.Run()
	}

	return nil
}
