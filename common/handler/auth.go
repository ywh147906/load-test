package handler

import (
	"github.com/ywh147906/load-test/common/ctx"
	"github.com/ywh147906/load-test/common/errmsg"
	"github.com/ywh147906/load-test/common/proto/models"
)

func GameServerAuth(next HandleFunc) HandleFunc {
	return func(ctx *ctx.Context) (err *errmsg.ErrMsg) {
		if ctx.ServerType != models.ServerType_GameServer {
			return errmsg.NewProtocolErrorInfo("invalid server type request")
		}
		return next(ctx)
	}
}

func GMAuth(next HandleFunc) HandleFunc {
	return func(ctx *ctx.Context) (err *errmsg.ErrMsg) {
		if ctx.ServerType != models.ServerType_GMServer {
			return errmsg.NewProtocolErrorInfo("invalid server type request")
		}
		return next(ctx)
	}
}
