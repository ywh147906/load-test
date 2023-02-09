package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync/atomic"
	"time"

	servicepb "github.com/ywh147906/load-test/common/proto/service"

	"github.com/ywh147906/load-test/common/logger"
	"github.com/ywh147906/load-test/common/msgcreate"
	"github.com/ywh147906/load-test/common/network/stdtcp"
	"github.com/ywh147906/load-test/common/pprof"
	"github.com/ywh147906/load-test/common/proto/cppbattle"
	lessservicepb "github.com/ywh147906/load-test/common/proto/less_service"
	"github.com/ywh147906/load-test/common/proto/models"
	"github.com/ywh147906/load-test/common/protocol"

	"go.uber.org/zap/zapcore"
)

var count uint64

type Client struct {
	log          *logger.Logger
	sess         *stdtcp.Session
	battleClient *Client
	type_        int64
	roleId       string
}

func NewClient(addr string, log *logger.Logger) *Client {
	c := &Client{
		log: log,
	}
	rand.Seed(time.Now().UnixNano())
	stdtcp.Connect(addr, time.Second*3, true, c, log, false)
	return c
}

// var userId = time.Now().UnixNano()
var userId = int64(283001002)

func (this_ *Client) OnConnected(session *stdtcp.Session) {
	if this_.type_ == 0 {
		this_.sess = session
		u := strconv.Itoa(int(atomic.AddInt64(&userId, 1)))
		//u := "ywhxx-05"
		rl := &lessservicepb.User_RoleLoginRequest{
			UserId:        u,
			ServerId:      11,
			AppKey:        "",
			Language:      0,
			RuleVersion:   "",
			Version:       0,
			ClientVersion: "0.0.1",
		}
		session.SetMeta(rl.UserId)
		respLogin := &lessservicepb.User_RoleLoginResponse{}
		err := session.RPCRequestOut(nil, rl, respLogin)
		if err != nil {
			panic(err)
		}

		respCenter := &servicepb.GameBattle_GetCurrBattleInfoResponse{}
		err1 := session.RPCRequestOut(nil, &servicepb.GameBattle_GetCurrBattleInfoRequest{}, respCenter)
		if err1 != nil {
			panic(err1)
		}
		fmt.Println("CurBattleInfoResponse:", respCenter)

		resp := &servicepb.GameBattle_CPPEnterBattleResponse{}
		err = session.RPCRequestOut(nil, &servicepb.GameBattle_CPPEnterBattleRequest{Pos: &cppbattle.CPPBattle_Vec2{X: 0, Y: 0}, MapId: respCenter.HungupMapId, BattleServerId: respCenter.BattleId}, resp)
		if err != nil {
			panic(err)
			os.Exit(1)
		}

		this_.battleClient = &Client{log: this_.log, type_: 1, roleId: respLogin.RoleId}

		stdtcp.Connect(resp.Ip+":"+strconv.Itoa(int(resp.Port)), time.Second*3, true, this_.battleClient, this_.log, false)

	} else {
		push := cppbattle.NSNB_AuthTCPPushToServer{
			RoleId: this_.roleId,
			Token:  "",
		}
		err := session.Send(&models.ServerHeader{RoleId: this_.roleId}, &push)
		if err != nil {
			panic(err)
		}
	}
}

func (this_ *Client) OnDisconnected(session *stdtcp.Session, err error) {
	fmt.Println(session.GetMeta(), err)
}

func (this_ *Client) OnRequest(session *stdtcp.Session, rpcIndex uint32, msgName string, frame []byte) {

}

func (this_ *Client) OnMessage(session *stdtcp.Session, msgName string, frame []byte) {
	h := &models.ServerHeader{}
	msg := msgcreate.NewMessage(msgName)
	err := protocol.DecodeInternal(frame, h, msg)
	if err != nil {
		panic(err)
	}

	if msgName == (&models.PING{}).XXX_MessageName() {
		_ = session.Send(nil, &models.PONG{})
	} else {
		atomic.AddUint64(&count, 1)
		fmt.Println(msgcreate.MessageName(msg), msg)
	}
}

func (this_ *Client) Close() {
	this_.sess.Close(nil)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	log := logger.MustNew(zapcore.DebugLevel, &logger.Options{
		Console:     "stdout",
		FilePath:    nil,
		RemoteAddr:  nil,
		InitFields:  nil,
		Development: true,
	})
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(atomic.LoadUint64(&count))
		}
	}()
	pprof.Start(log, ":6011")
	m := map[*Client]struct{}{}
	for i := 0; i < 40; i++ {
		c := NewClient("10.23.50.87:8071", log)
		m[c] = struct{}{}
		time.Sleep(time.Millisecond)
	}
	for {
		time.Sleep(time.Second)
	}
}
