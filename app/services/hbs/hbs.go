package hbs

import (
	"time"

	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/spf13/viper"
	"github.com/open-falcon/falcon-plus/modules/api/config"
)

var (
	rpcClient   *config.SingleConnRpcClient
	connTimeout int
)

func Start(address string) {
	connTimeout = viper.GetInt("hbs.conn_timeout")
	rpcClient = &config.SingleConnRpcClient{
		RpcServer: address,
		Timeout:   time.Duration(int64(connTimeout)) * time.Millisecond,
	}
}

func GetHostStrategies(hostId int) (resp model.HostStrategy, err error) {
	req := model.HostStrategyRequest{HostId: hostId}
	err = rpcClient.Call("Hbs.GetHostStrategies", req, &resp)
	return
}
