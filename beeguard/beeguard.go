package beeguard

import (
	"database/sql"

	"github.com/NooBeeID/bee-guard/infra/contracts"
	"github.com/NooBeeID/bee-guard/infra/modules"
	"github.com/NooBeeID/bee-guard/infra/router"
	"github.com/NooBeeID/bee-guard/modules/auth/login"
	logincache "github.com/NooBeeID/bee-guard/modules/auth/login/resources/cache"
	loginpostgres "github.com/NooBeeID/bee-guard/modules/auth/login/resources/postgresql"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type BeeGuard struct {
	router  *router.Router
	db      *sql.DB
	cache   contracts.Cache
	modules *modules.Base
	server  modules.Server
}

func New() *BeeGuard {
	return &BeeGuard{}
}

func (b *BeeGuard) SetRouter(r any) *BeeGuard {
	beeRouter := router.New(r)
	b.router = beeRouter
	return b
}

func (b *BeeGuard) SetSqlFromSQLX(db *sqlx.DB) *BeeGuard {
	dbNative := db.DB
	return b.setSQL(dbNative)
}

func (b *BeeGuard) SetSqlFromGORM(db *gorm.DB) *BeeGuard {
	dbNative, err := db.DB()
	if err != nil {
		panic(err)
	}
	return b.setSQL(dbNative)
}

func (b *BeeGuard) setSQL(db *sql.DB) *BeeGuard {
	b.db = db
	return b
}

func (b *BeeGuard) SetModules(mods ...modules.Modules) *BeeGuard {
	list := modules.NewModules(b.router, mods...)
	b.modules = &list

	server := modules.NewServer(b.router, b.modules.GetModules()...)
	b.server = server
	return b
}

func (b *BeeGuard) SetCustomModules(mods ...modules.Modules) *BeeGuard {
	if b.modules == nil {
		b.SetDefaultModules()
	}
	b.server.AddCustomModules(mods...)
	return b
}

func (b *BeeGuard) Run() error {
	b.server.Start()
	return nil
}

func (b *BeeGuard) SetDefaultModules() *BeeGuard {
	if b.modules == nil {
		loginMod := login.New(&modules.ConfigService{
			Router: b.router,
			Handler: login.NewHandler(login.NewService(
				loginpostgres.New(b.db),
				logincache.New(b.cache),
			)),
		})

		list := modules.NewModules(b.router, loginMod)
		b.SetModules(list.GetModules()...)
	}
	return b
}
