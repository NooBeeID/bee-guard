package registercache

import (
	"context"

	"github.com/NooBeeID/bee-guard/infra/contracts"
)

type cache struct {
	client contracts.Cache
}

// StoreSession implements login.contractCacheRepository.
func (c *cache) StoreSession(ctx context.Context, sessionId string, value string) (err error) {
	err = c.client.Set(ctx, sessionId, value)
	return
}

func New(client contracts.Cache) *cache {
	return &cache{
		client: client,
	}
}
