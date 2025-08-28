package beeguard

import (
	"database/sql"

	"github.com/NooBeeID/bee-guard/infra/router"
	"github.com/NooBeeID/bee-guard/modules"
	"github.com/NooBeeID/bee-guard/modules/auth/login"
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type BeeGuard struct {
	router *router.Router
	db     *sql.DB
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
	loginSvc := login.New(&modules.ConfigService{
		Router: b.router,
		Db:     b.db,
	})
	loginSvc.Run()
	return nil
}
