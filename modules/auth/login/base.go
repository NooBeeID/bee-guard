package login

import (
	"database/sql"

	"github.com/NooBeeID/bee-guard/infra/router"
	"github.com/NooBeeID/bee-guard/modules"
)

type Login struct {
	beeRouter *router.Router
	db        *sql.DB
}

func New(cfg *modules.ConfigService) *Login {
	if cfg == nil {
		return &Login{}
	}

	return &Login{
		beeRouter: cfg.Router,
		db:        cfg.Db,
	}
}

func (l *Login) SetRouter(r any) *Login {
	l.beeRouter = router.New(r)
	return l
}

func (l *Login) SetSQL(db *sql.DB) *Login {
	l.db = db
	return l
}

func (l *Login) Run() {
	svc := NewService(nil)
	handler := NewHandler(svc)

	l.beeRouter.Post("/auth/login", handler.Login)
}
