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

package template

import (
	"net/http"

	"github.com/RosenLo/falcon-api/app/utils"
	"github.com/RosenLo/falcon-api/config"
	"github.com/gin-gonic/gin"
)

var db config.DBPool

const badstatus = http.StatusBadRequest

func Routes(r *gin.Engine) {
	db = config.Con()
	tmpr := r.Group("/api/v1/template")
	tmpr.Use(utils.AuthSessionMidd)
	tmpr.GET("", GetTemplates)
	tmpr.POST("", CreateTemplate)
	tmpr.GET("/:tpl_id", GetATemplate)
	tmpr.PUT("", UpdateTemplate)
	tmpr.DELETE("/:tpl_id", DeleteTemplate)
	tmpr.POST("/action", CreateActionToTmplate)
	tmpr.PUT("/action", UpdateActionToTmplate)

	actr := r.Group("/api/v1/action")
	actr.GET("/:act_id", GetActionByID)

	//simple list for ajax use
	tmpr2 := r.Group("/api/v1/template_simple")
	tmpr2.Use(utils.AuthSessionMidd)
	tmpr2.GET("", GetTemplatesSimple)

	tmpr3 := r.Group("/api/v1/template_name")
	tmpr3.Use(utils.AuthSessionMidd)
	tmpr3.GET("/name/:tpl_name", GetTemplateByName)
}
