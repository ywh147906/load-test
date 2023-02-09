package main

import (
	"fmt"

	"github.com/ywh147906/load-test/common/consulkv"
	"github.com/ywh147906/load-test/common/idgenerate"
	"github.com/ywh147906/load-test/common/logger"
	"github.com/ywh147906/load-test/common/proto/models"
	"github.com/ywh147906/load-test/common/redisclient"
	"github.com/ywh147906/load-test/common/ulimit"
	"github.com/ywh147906/load-test/common/utils"
	"github.com/ywh147906/load-test/common/values/env"
	"github.com/ywh147906/load-test/core"
	"github.com/ywh147906/load-test/module"

	"go.uber.org/zap"
)

func main() {
	err := ulimit.SetRLimit()
	utils.Must(err)
	log := logger.MustNewAsync(zap.DebugLevel, &logger.Options{
		Console:    "stdout",
		FilePath:   []string{fmt.Sprintf("./%s.log", models.ServerType_LoadTest.String())},
		RemoteAddr: nil,
		InitFields: map[string]interface{}{
			"serverType": models.ServerType_LoadTest,
			"serverId":   env.GetServerId(),
		},
		Development: true,
		Discard:     true,
	})
	logger.SetDefaultLogger(log)
	cnf, err := consulkv.NewConfig(env.GetString(env.CONF_PATH), env.GetString(env.CONF_HOSTS), log)
	utils.Must(err)
	redisclient.Init(cnf)
	idgenerate.Init(redisclient.GetCommonRedis())
	module.Init()
	core.RunBoomer(env.GetString(env.LOCUST_MASTER_HOST), int(env.GetInteger(env.LOCUST_MASTER_PORT)), module.GetModuleTasks())
}
