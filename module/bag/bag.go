package bag

import (
	"github.com/ywh147906/load-test/assert"
	"github.com/ywh147906/load-test/common/errmsg"
	pbSvr "github.com/ywh147906/load-test/common/proto/service"
	"github.com/ywh147906/load-test/common/values"
	"github.com/ywh147906/load-test/core"

	"github.com/gogo/protobuf/proto"
	"go.uber.org/zap"
)

// 每个用户对应单独的一个service实例
type service struct {
	ItemId values.ItemId
}

func New(ctx *core.RoleContext) core.ILoadTestModule {
	s := &service{}
	s.registerPush(ctx)
	return s
}

func (s *service) Process(ctx *core.RoleContext) {
	s.reset(ctx)
	s.cheatAddItem(ctx)
	s.getBag(ctx)
	s.useItem(ctx)
}

func (s *service) reset(ctx *core.RoleContext) {
	s.ItemId = 10007
}

func (s *service) registerPush(ctx *core.RoleContext) {
	msg := &pbSvr.Bag_ItemUpdatePush{}
	ctx.RegisterPush(msg.XXX_MessageName(), func(ctx *core.RoleContext, msg proto.Message) *errmsg.ErrMsg {
		m := msg.(*pbSvr.Bag_ItemUpdatePush)
		ctx.Debug("Bag_ItemUpdatePush", zap.Any("items", m.Items))
		return nil
	})
}

func (s *service) cheatAddItem(ctx *core.RoleContext) {
	req := &pbSvr.Bag_AddItemsRequest{Items: map[int64]int64{
		10000: 10,
		10001: 10,
		10002: 10,
		10007: 10,
	}}
	_, res, err := ctx.Request(req)
	assert.Nil(ctx, err)
	assert.NotNil(ctx, res)
	_ = res.(*pbSvr.Bag_AddItemsResponse)
}

func (s *service) getBag(ctx *core.RoleContext) {
	req := &pbSvr.Bag_GetBagInfoRequest{}
	_, res, err := ctx.Request(req)
	assert.Nil(ctx, err)
	assert.NotNil(ctx, res)
	_ = res.(*pbSvr.Bag_GetBagInfoResponse)
}

func (s *service) useItem(ctx *core.RoleContext) {
	req := &pbSvr.Bag_UseItemRequest{
		ItemId: 10007,
		Count:  1,
		Choose: nil,
	}
	_, res, err := ctx.Request(req)
	assert.Nil(ctx, err)
	assert.NotNil(ctx, res)
	_ = res.(*pbSvr.Bag_UseItemResponse)
}
