package logincache

import (
	"context"

	"github.com/NooBeeID/bee-guard/infra/contracts"
)

type cacheRedis struct {
	client contracts.Cache
}

// StoreSession implements login.contractCacheRepository.
func (c *cacheRedis) StoreSession(ctx context.Context, sessionId string, value string) (err error) {
	return
}

func New(client contracts.Cache) *cacheRedis {
	return &cacheRedis{
		client: client,
	}
}
