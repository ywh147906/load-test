package client_version

import (
	"github.com/ywh147906/load-test/common/redisclient"
)

const (
	ClientVersionKey = "CLIENT_VERSION"
	AuditVersionKey  = "AUDIT_VERSION"
)

func GetClientVersionRedis() redis.Cmdable {
	return redisclient.GetDefaultRedis()
}
