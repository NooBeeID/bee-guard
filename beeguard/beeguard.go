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
	router *router.Router
	db     *sql.DB
	cache contracts.Cache
	server modules.Server
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

func (b *BeeGuard) Run() error {
	loginMod := login.New(&modules.ConfigService{
		Handler: login.NewHandler(login.NewService(
			loginpostgres.New(b.db),
			logincache.New(b.cache),
		)),
	})


	listModules := modules.NewModules(b.router, loginMod)
	server := modules.NewServer(b.router, listModules.GetModules()...)
	server.Start()
	return nil
}
