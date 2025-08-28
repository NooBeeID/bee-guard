package modules

import (
	"database/sql"

	"github.com/NooBeeID/bee-guard/infra/router"
)

type ConfigService struct {
	Router *router.Router
	Db     *sql.DB
}
