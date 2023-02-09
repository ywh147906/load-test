package stdtcp

import (
	"net"
	"time"

	"github.com/ywh147906/load-test/common/gopool"
	"github.com/ywh147906/load-test/common/logger"

	"go.uber.org/zap"
)

func Connect(addr string, timeout time.Duration, isReconnect bool, h dispatch, log *logger.Logger, isDebug bool) {
	gopool.Submit(func() {
		conn, err := net.DialTimeout("tcp", addr, timeout)
		if err != nil {
			sess := newErrClientSession(addr, log, isDebug, h, isReconnect, timeout)
			sess.Close(err)
			log.Debug("connect failed", zap.String("remoteAddr", addr))
		} else {
			sess := newClientSession(conn, log, isDebug, h, isReconnect, timeout)
			sess.start()
			h.OnConnected(sess)
		}
	})
}
