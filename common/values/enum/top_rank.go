package enum

import (
	"github.com/go-redis/redis/v8"
	"github.com/ywh147906/load-test/common/redisclient"
)

const (
	TopRankId    = "top_rank"
	TopRankLimit = "top_rank_limit"
)

func GetRedisClient() redis.Cmdable {
	return redisclient.GetDefaultRedis()
}
