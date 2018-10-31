// Copyright 2018 RosenLo
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package hbs

import (
	"time"

	"github.com/RosenLo/falcon-api/config"
	"github.com/open-falcon/falcon-plus/common/model"
	"github.com/spf13/viper"
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
