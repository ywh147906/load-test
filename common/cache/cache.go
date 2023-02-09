package cache

import (
	"context"

	"github.com/ywh147906/load-test/common/values"
)

type ICache interface {
	Get(key string) (val interface{}, exist bool)
	Set(key string, val interface{})
	Del(key string)
}

func GetCache(ctx context.Context, roleId values.RoleId) ICache {
	return gCache
}

var gCache ICache
