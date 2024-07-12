package conroller

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/snusEbjoer/cipherService/internal/handlers"
)

type Conroller struct {
	service *handlers.Service
}

const TTL = 10 * time.Second

func NewController(cache *redis.Client, ctx context.Context) *Conroller {
	s := handlers.NewEncodeService(cache, ctx)
	return &Conroller{
		service: s,
	}
}
func (c *Conroller) EncodeSHA256() handlers.Handler {
	return c.service.EncodeSHA256()
}

func (c *Conroller) EncodeMD5() handlers.Handler {
	return c.service.EncodeMD5()
}
