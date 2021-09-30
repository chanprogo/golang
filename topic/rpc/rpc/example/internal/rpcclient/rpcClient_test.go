package rpcclient

import (
	"testing"

	"example/app"
	"example/conf"
	"example/internal/protodatasvr"
	"example/internal/rpcserver"
)

func init() {
	conf.NewConfiger("../../../cmd/smsrpcsvr/app.conf")

	app.NewRPCSvr(conf.Configer.RpcConf.RpcPort)
	protodatasvr.RegisterEmailServiceServer(app.RPCSvr, &rpcserver.SendEmailServer{})
}

func TestRpcClient(t *testing.T) {
	RpcClient()
}
