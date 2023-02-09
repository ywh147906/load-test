package racingrank

import (
	"context"
	"time"

	"github.com/ywh147906/load-test/common/consulkv"
	"github.com/ywh147906/load-test/common/errmsg"
	"github.com/ywh147906/load-test/common/logger"
	"github.com/ywh147906/load-test/common/proto/dao"
	"github.com/ywh147906/load-test/common/utils"

	"go.uber.org/zap"
)

type Emit struct {
	log *logger.Logger
	kw  *kafka.Writer
}

type Config struct {
	Addr  []string `json:"addr"`
	Topic string   `json:"topic"`
}

var emitter *Emit

func Init(cfg *consulkv.Config) {
	conf := &Config{}
	utils.Must(cfg.Unmarshal("racingrank/kafka", conf))
	emitter = NewEmit(logger.DefaultLogger, conf)
}

func Close() {
	if emitter != nil {
		emitter.kw.Close()
	}
}

func NewEmit(log *logger.Logger, cfg *Config) *Emit {
	writer := &kafka.Writer{
		Addr:         kafka.TCP(cfg.Addr...),
		Topic:        cfg.Topic,
		Balancer:     &kafka.Murmur2Balancer{},
		RequiredAcks: kafka.RequireAll,
		MaxAttempts:  5,
		BatchSize:    200,
		BatchTimeout: 100 * time.Millisecond,
	}
	return &Emit{
		log: log,
		kw:  writer,
	}
}

func (e *Emit) Emit(ctx context.Context, rrm *dao.RacingRankMatch) *errmsg.ErrMsg {
	value, _ := rrm.Marshal()
	if len(value) <= 0 {
		return nil
	}
	if err := e.kw.WriteMessages(ctx, kafka.Message{
		Key:   []byte(rrm.RoleId),
		Value: value,
	}); err != nil {
		e.log.Error("emit racingrank match data err", zap.Error(err), zap.String("data", string(value)))
		return errmsg.NewInternalErr(err.Error())
	}
	return nil
}

func Emitting(ctx context.Context, rrm *dao.RacingRankMatch) *errmsg.ErrMsg {
	return emitter.Emit(ctx, rrm)
}
