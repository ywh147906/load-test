package divination

import (
	"fmt"

	"github.com/ywh147906/load-test/assert"
	pbSvr "github.com/ywh147906/load-test/common/proto/service"
	"github.com/ywh147906/load-test/core"
)

type service struct {
}

func New(ctx *core.RoleContext) core.ILoadTestModule {
	return &service{}
}

func (s service) Process(ctx *core.RoleContext) {
	s.getDivinationInfo(ctx)
	s.divinationOnce(ctx)
	s.cheatResetTimes(ctx)
}

func (s *service) getDivinationInfo(ctx *core.RoleContext) {
	req := &pbSvr.Divination_DivinationInfoRequest{}
	_, res, err := ctx.Request(req)
	assert.Nil(ctx, err)
	_ = res.(*pbSvr.Divination_DivinationInfoResponse)
}

func (s *service) divinationOnce(ctx *core.RoleContext) {
	req := &pbSvr.Divination_DivinationOnceRequest{}
	_, res, err := ctx.Request(req)
	assert.Nil(ctx, err)
	result := res.(*pbSvr.Divination_DivinationOnceResponse)
	if result.Info.AvailableCount == result.Info.TotalCount {
		panic(fmt.Errorf("divination faild"))
	}
}

func (s *service) cheatResetTimes(ctx *core.RoleContext) {
	req := &pbSvr.Divination_CheatResetDivinationRequest{}
	_, _, err := ctx.Request(req)
	assert.Nil(ctx, err)
}
