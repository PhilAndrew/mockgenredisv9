package mock_v9

import (
	"context"

	"github.com/bsm/redislock"
	"github.com/redis/go-redis/v9"
)

type Cache interface {
	TestMGet(ctx context.Context) error
}

// to take the cycle lock.
type RedisV2 struct {
	client redis.Cmdable
	locker *redislock.Client
}

func NewRedisCacheV2(client redis.Cmdable) Cache {
	cache := new(RedisV2)
	cache.client = client

	return cache
}

// into a new response and returned.
func (cache *RedisV2) TestMGet(ctx context.Context) error {
	cache.client.MGet(ctx, "cacheKey", "deleteKey").Result()
	return nil
}
