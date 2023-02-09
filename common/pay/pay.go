package pay

import (
	"github.com/go-redis/redis/v8"
	"github.com/ywh147906/load-test/common/redisclient"
	"github.com/ywh147906/load-test/common/utils"
)

const PayQueueKey = "pay_queue"

func GetKey(roleId values.RoleId) string {
	return utils.GenDefaultRedisKey(values.Pay, values.Hash, roleId)
}

func GetPayRedis() redis.Cmdable {
	return redisclient.GetDefaultRedis()
}

func GetPayQueueRedis() redis.Cmdable {
	return redisclient.GetDefaultRedis()
}
