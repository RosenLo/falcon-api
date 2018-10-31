// Copyright 2017 Xiaomi, Inc.
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

package controller

import (
	"net/http"

	"github.com/RosenLo/falcon-api/app/controller/alarm"
	"github.com/RosenLo/falcon-api/app/controller/dashboard_graph"
	"github.com/RosenLo/falcon-api/app/controller/dashboard_screen"
	"github.com/RosenLo/falcon-api/app/controller/expression"
	"github.com/RosenLo/falcon-api/app/controller/graph"
	"github.com/RosenLo/falcon-api/app/controller/host"
	"github.com/RosenLo/falcon-api/app/controller/mockcfg"
	"github.com/RosenLo/falcon-api/app/controller/strategy"
	"github.com/RosenLo/falcon-api/app/controller/template"
	"github.com/RosenLo/falcon-api/app/controller/uic"
	"github.com/RosenLo/falcon-api/app/utils"
	"github.com/gin-gonic/gin"
)

func StartGin(port string, r *gin.Engine) {
	r.Use(utils.CORS())
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, I'm Falcon+ (｡A｡)")
	})
	graph.Routes(r)
	uic.Routes(r)
	template.Routes(r)
	strategy.Routes(r)
	host.Routes(r)
	expression.Routes(r)
	mockcfg.Routes(r)
	dashboard_graph.Routes(r)
	dashboard_screen.Routes(r)
	alarm.Routes(r)
	r.Run(port)
}
