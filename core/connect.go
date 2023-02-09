package core

import (
	"github.com/ywh147906/load-test/common/errmsg"
	modelsPb "github.com/ywh147906/load-test/common/proto/models"

	"github.com/gogo/protobuf/proto"
)

type PushFunc func(ctx *RoleContext, msg proto.Message) *errmsg.ErrMsg

type IConnect interface {
	Request(req proto.Message) (resp *modelsPb.Resp, res proto.Message, err *errmsg.ErrMsg)
	AsyncSend(req proto.Message) (err *errmsg.ErrMsg)
	RegisterPush(name string, f PushFunc)
	Close()
}
