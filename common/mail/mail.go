package mail

import (
	"github.com/ywh147906/load-test/common/redisclient"
	"github.com/ywh147906/load-test/common/utils"
	"github.com/ywh147906/load-test/common/values"
)

const EntireMailKey = "entire_mail"

func GetMailRedis() redis.Cmdable {
	return redisclient.GetDefaultRedis()
}

func GetMailKey(roleId values.RoleId) string {
	return utils.GenDefaultRedisKey(values.Mail, values.Hash, roleId)
}
