package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/ywh147906/load-test/common/logger"
	"github.com/ywh147906/load-test/common/network/stdtcp"

	"go.uber.org/zap/zapcore"
)

type Server struct {
	acceptor *stdtcp.Acceptor
	log      *logger.Logger
	addr     string
	stdtcp.EventHandler
}

// NewServer isForClient是不是监听客户端的连接，如果是客户端，那么读取数据的方式会不一样
func NewServer(addr string, log *logger.Logger, isForClient bool) (*Server, error) {
	s := &Server{log: log, addr: addr}
	a, err := stdtcp.NewAcceptor(addr, log, true, s, isForClient)
	if err != nil {
		return nil, err
	}
	s.acceptor = a
	s.log = log
	a.Start()
	return s, nil
}

// OnConnected 当监听连接建立时,当连接服务器成功时 会调用
func (this_ *Server) OnConnected(session *stdtcp.Session) {

}

func (this_ *Server) OnDisconnected(session *stdtcp.Session, err error) {

}

func (this_ *Server) OnRequest(session *stdtcp.Session, rpcIndex uint32, msgName string, frame []byte) {

}

func (this_ *Server) OnMessage(session *stdtcp.Session, msgName string, frame []byte) {

}

func (this_ *Server) Close() {
	this_.acceptor.Close()
}

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("0.0.0.0:16667", nil))
	}()
	log := logger.MustNew(zapcore.DebugLevel, &logger.Options{
		Console:     "",
		FilePath:    nil,
		RemoteAddr:  nil,
		InitFields:  nil,
		Development: true,
	})
	s, err := NewServer("0.0.0.0:6666", log, false)
	if err != nil {
		panic(err)
	}
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGHUP)
	<-ch
	s.Close()
}
