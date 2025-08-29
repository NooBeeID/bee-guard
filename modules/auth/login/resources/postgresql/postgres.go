package loginpostgres

import (
	"context"
	"database/sql"

	"github.com/NooBeeID/bee-guard/entity"
)

type postgres struct {
	db *sql.DB
}

// GetAuthByEmail implements login.contractDBRepository.
func (p *postgres) GetAuthByEmail(ctx context.Context, email string) (auth entity.Auth, err error) {
	return
}

func New(db *sql.DB) *postgres {
	return &postgres{
		db: db,
	}
}
