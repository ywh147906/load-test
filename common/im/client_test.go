package im

import (
	"context"
	"testing"

	"github.com/ywh147906/load-test/common/consulkv"
	"github.com/ywh147906/load-test/common/logger"
	"github.com/ywh147906/load-test/common/proto/models"
	"github.com/ywh147906/load-test/common/utils"
	"github.com/ywh147906/load-test/common/values/env"

	"go.uber.org/zap"
)

func init() {
	serverId := env.GetServerId()
	log := logger.MustNew(zap.DebugLevel, &logger.Options{
		Console: "stdout",
		// FilePath:   []string{fmt.Sprintf("./%s.log", models.ServerType_GameServer.String())},
		RemoteAddr: nil,
		InitFields: map[string]interface{}{
			"serverType": models.ServerType_GameServer,
			"serverId":   serverId,
		},
		Development: true,
		//Discard:     true,
	})
	logger.SetDefaultLogger(log)
	cnf, err := consulkv.NewConfig(env.GetString(env.CONF_PATH), env.GetString(env.CONF_HOSTS), log)
	utils.Must(err)
	Init(cnf)
}

// TestMarquee 测试跑马灯
func TestMarquee(t *testing.T) {
	DefaultClient.SendMessage(context.Background(), &Message{
		Type:       MsgTypeBroadcast,
		RoleID:     "admin",
		RoleName:   "admin",
		Content:    "测试跑马灯",
		ParseType:  ParseTypeSys,
		IsMarquee:  true,
		IsVolatile: true,
	})
}

func TestBanPost(t *testing.T) {
	DefaultClient.BanPost(context.Background(), "321", 60)
}

func TestUnbanPost(t *testing.T) {
	DefaultClient.UnBanPost(context.Background(), "321")
}
