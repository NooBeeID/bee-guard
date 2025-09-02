package beeguard

import (
	"database/sql"

	beeguardcache "github.com/NooBeeID/bee-guard/beeguard/cache"
	"github.com/NooBeeID/bee-guard/infra/contracts"
	"github.com/NooBeeID/bee-guard/infra/modules"
	"github.com/NooBeeID/bee-guard/infra/router"
	"github.com/NooBeeID/bee-guard/modules/auth/login"
	"github.com/NooBeeID/bee-guard/modules/auth/register"
	registercache "github.com/NooBeeID/bee-guard/modules/auth/register/resources/cache"
	registerdatabase "github.com/NooBeeID/bee-guard/modules/auth/register/resources/database"

	logincache "github.com/NooBeeID/bee-guard/modules/auth/login/resources/cache"
	logindatabase "github.com/NooBeeID/bee-guard/modules/auth/login/resources/database"
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

func (b *BeeGuard) SetCache(cache contracts.Cache) *BeeGuard { b.cache = cache; return b }

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
	var cache contracts.Cache
	if b.cache == nil {
		cache = beeguardcache.NewMemoryCache()
		b.cache = cache
	}
	if b.modules == nil {
		loginMod := login.New(&modules.ConfigService{
			Router: b.router,
			Handler: login.NewHandler(login.NewService(
				logindatabase.New(b.db),
				logincache.New(b.cache),
			)),
		})

		registerMod := register.New(&modules.ConfigService{
			Router: b.router,
			Handler: register.NewHandler(register.NewService(
				registerdatabase.New(b.db),
				registercache.New(b.cache),
			)),
		})

		list := modules.NewModules(b.router, loginMod, registerMod)
		b.SetModules(list.GetModules()...)
	}
	return b
}
