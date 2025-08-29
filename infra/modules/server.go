package modules

import (
	"log/slog"

	"github.com/NooBeeID/bee-guard/infra/router"
)

type Server struct {
	mods   map[string]Modules
	router *router.Router
}

func NewServer(router *router.Router, mods ...Modules) Server {
	if len(mods) == 0 {
		slog.Warn("no modules provided, server will not have any functionality")
		mods = make([]Modules, 0)
	}

	var myModules = map[string]Modules{}

	// Initialize the map to hold module names and their instances
	for _, mod := range mods {
		if mod == nil {
			slog.Error("nil module provided, skipping")
			continue
		}
		if mod.GetName() == "" {
			slog.Error("module name is empty, skipping")
			continue
		}

		if _, exists := myModules[mod.GetName()]; exists {
			slog.Error("duplicate module name detected: " + mod.GetName() + ", skipping this module")
			continue
		}

		slog.Debug("adding module: " + mod.GetName())
		myModules[mod.GetName()] = mod
	}
	return Server{
		mods:   myModules,
		router: router,
	}
}

func (s *Server) AddCustomModules(mods ...Modules) {

	for _, mod := range mods {
		slog.Info("custom module " + mod.GetName())
		if mod == nil {
			slog.Error("nil module provided, skipping")
			continue
		}

		s.mods[mod.GetName()] = mod
		slog.Debug("added custom module: " + mod.GetName())
	}
}

func (s *Server) Start() error {
	for _, mod := range s.mods {
		mod.Run()
		slog.Info("module " + mod.GetName() + " started")
	}

	return nil
}
