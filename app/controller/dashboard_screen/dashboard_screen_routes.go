// Copyright 2018 RosenLo

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

/**
 * This code was originally worte by Xiaomi, Inc. modified by RosenLo.
**/

package dashboard_screen

import (
	"net/http"

	"github.com/RosenLo/falcon-api/app/utils"
	"github.com/RosenLo/falcon-api/config"
	"github.com/gin-gonic/gin"
)

var db config.DBPool

const badstatus = http.StatusBadRequest
const expecstatus = http.StatusExpectationFailed

func Routes(r *gin.Engine) {
	db = config.Con()
	authapi := r.Group("/api/v1/dashboard")
	authapi.Use(utils.AuthSessionMidd)
	authapi.POST("/screen", ScreenCreate)
	authapi.GET("/screen/:screen_id", ScreenGet)
	authapi.GET("/screen_name/:screen_name", ScreenGetByName)
	authapi.GET("/screens/pid/:pid", ScreenGetsByPid)
	authapi.GET("/screens", ScreenGetsAll)
	authapi.DELETE("/screen/:screen_id", ScreenDelete)
	authapi.PUT("/screen/:screen_id", ScreenUpdate)
}
