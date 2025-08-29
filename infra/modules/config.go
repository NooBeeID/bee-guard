package modules

import (
	"database/sql"

	"github.com/NooBeeID/bee-guard/infra/contracts"
	"github.com/NooBeeID/bee-guard/infra/router"
)

type ConfigService struct {
	Router *router.Router
	DB     *sql.DB
	Cache contracts.Cache
	Handler contracts.Handler
}


