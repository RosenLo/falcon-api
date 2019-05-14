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

package host

import (
	"fmt"
	"strconv"

	h "github.com/RosenLo/falcon-api/app/helper"
	f "github.com/RosenLo/falcon-api/app/model/falcon_portal"
	"github.com/RosenLo/falcon-api/app/services/hbs"
	u "github.com/RosenLo/falcon-api/app/utils"
	log "github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type APIMaintainInput struct {
	MaintainBegin uint64 `json:"maintain_begin"`
	MaintainEnd   uint64 `json:"maintain_end"`
}

func GetHostBindToWhichHostGroup(c *gin.Context) {
	HostIdTmp := c.Params.ByName("host_id")
	if HostIdTmp == "" {
		h.JSONR(c, badstatus, "host id is missing")
		return
	}
	hostID, err := strconv.Atoi(HostIdTmp)
	if err != nil {
		log.Debugf("HostId: %v", HostIdTmp)
		h.JSONR(c, badstatus, err)
		return
	}
	grpHostMap := []f.GrpHost{}
	db.Falcon.Select("grp_id").Where("host_id = ?", hostID).Find(&grpHostMap)
	grpIds := []int64{}
	for _, g := range grpHostMap {
		grpIds = append(grpIds, g.GrpID)
	}
	hostgroups := []f.HostGroup{}
	if len(grpIds) != 0 {
		grpIdsStr, _ := u.ArrInt64ToString(grpIds)
		db.Falcon.Where(fmt.Sprintf("id in (%s)", grpIdsStr)).Find(&hostgroups)
	}
	h.JSONR(c, hostgroups)
	return
}

func GetHostGroupWithTemplate(c *gin.Context) {
	grpIDtmp := c.Params.ByName("host_group")
	if grpIDtmp == "" {
		h.JSONR(c, badstatus, "grp id is missing")
		return
	}
	grpID, err := strconv.Atoi(grpIDtmp)
	if err != nil {
		log.Debugf("grpIDtmp: %v", grpIDtmp)
		h.JSONR(c, badstatus, err)
		return
	}
	hostgroup := f.HostGroup{ID: int64(grpID)}
	if dt := db.Falcon.Find(&hostgroup); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	hosts := []f.Host{}
	grpHosts := []f.GrpHost{}
	if dt := db.Falcon.Where("grp_id = ?", grpID).Find(&grpHosts); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	for _, grph := range grpHosts {
		var host f.Host
		db.Falcon.Find(&host, grph.HostID)
		if host.ID != 0 {
			hosts = append(hosts, host)
		}
	}
	h.JSONR(c, map[string]interface{}{
		"hostgroup": hostgroup,
		"hosts":     hosts,
	})
	return
}

func GetGrpsRelatedHost(c *gin.Context) {
	hostIDtmp := c.Params.ByName("host_id")
	if hostIDtmp == "" {
		h.JSONR(c, badstatus, "host id is missing")
		return
	}
	hostID, err := strconv.Atoi(hostIDtmp)
	if err != nil {
		log.Debugf("host id: %v", hostIDtmp)
		h.JSONR(c, badstatus, err)
		return
	}

	host := f.Host{ID: int64(hostID)}
	if dt := db.Falcon.Find(&host); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	grps := host.RelatedGrp()
	h.JSONR(c, grps)
	return
}

func GetTplsRelatedHost(c *gin.Context) {
	hostIDtmp := c.Params.ByName("host_id")
	if hostIDtmp == "" {
		h.JSONR(c, badstatus, "host id is missing")
		return
	}
	hostID, err := strconv.Atoi(hostIDtmp)
	if err != nil {
		log.Debugf("host id: %v", hostIDtmp)
		h.JSONR(c, badstatus, err)
		return
	}
	host := f.Host{ID: int64(hostID)}
	if dt := db.Falcon.Find(&host); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	tpls := host.RelatedTpl()
	h.JSONR(c, tpls)
	return
}

func GetHost(c *gin.Context) {
	hostName := c.DefaultQuery("hostname", "")
	hostIp := c.DefaultQuery("ip", "")

	var hosts []f.Host
	if hostName == "" && hostIp == "" {
		if dt := db.Falcon.Find(&hosts); dt.Error != nil {
			h.JSONR(c, expecstatus, dt.Error)
			return
		}
	}
	if hostIp != "" {
		if dt := db.Falcon.Where("ip = ?", hostIp).Find(&hosts); dt.Error != nil {
			h.JSONR(c, expecstatus, dt.Error)
			return
		}
	}
	if hostName != "" {
		if dt := db.Falcon.Where("hostname like ?", fmt.Sprintf("%s%%", hostName)).Find(&hosts); dt.Error != nil {
			h.JSONR(c, expecstatus, dt.Error)
			return
		}
	}

	h.JSONR(c, hosts)
	return
}

func MaintainHost(c *gin.Context) {
	var cmaint APIMaintainInput
	err := c.Bind(&cmaint)
	if err != nil {
		h.JSONR(c, badstatus, err)
		return
	}

	hostIDtmp := c.Params.ByName("host_id")
	if hostIDtmp == "" {
		h.JSONR(c, badstatus, "host id is missing")
		return
	}

	hostID, err := strconv.Atoi(hostIDtmp)
	if err != nil {
		log.Debugf("host id: %v", hostIDtmp)
		h.JSONR(c, badstatus, err)
		return
	}

	host := f.Host{ID: int64(hostID)}
	if _, ok := host.Existing(); !ok {
		h.JSONR(c, badstatus, err)
		return
	}
	hhost := map[string]interface{}{
		"maintain_begin": cmaint.MaintainBegin,
		"maintain_end":   cmaint.MaintainEnd,
	}
	if dt := db.Falcon.Model(&host).Where("id = ?", host.ID).Updates(hhost).Find(&host); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	h.JSONR(c, host)
	return
}

func DeleteHost(c *gin.Context) {
	hostIDTmp := c.Params.ByName("host_id")
	if hostIDTmp == "" {
		h.JSONR(c, badstatus, "host id is missing")
		return
	}
	hostID, err := strconv.Atoi(hostIDTmp)
	if err != nil {
		log.Debugf("hostIDTmp: %v", hostIDTmp)
		h.JSONR(c, badstatus, err)
		return
	}
	host := f.Host{ID: int64(hostID)}
	if dt := db.Falcon.Find(&host); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	user, _ := h.GetUser(c)
	if !user.IsAdmin() {
		return
	}

	if dt := db.Falcon.Delete(&host); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	h.JSONR(c, fmt.Sprintf("host :%v has been deleted", hostID))
	return

}

func GetHostAlone(c *gin.Context) {
	var hosts []f.Host
	var grpHost []f.GrpHost
	var hostIds []int64
	db.Falcon.Select("DISTINCT(host_id)").Find(&grpHost)
	for _, host := range grpHost {
		hostIds = append(hostIds, host.HostID)
	}
	if dt := db.Falcon.Not("id", hostIds).Find(&hosts); dt.Error != nil {
		h.JSONR(c, expecstatus, dt.Error)
		return
	}
	h.JSONR(c, hosts)
	return
}

func GetHostStrategies(c *gin.Context) {
	hostIDtmp := c.Params.ByName("host_id")
	if hostIDtmp == "" {
		h.JSONR(c, badstatus, "host id is missing")
		return
	}
	hostID, err := strconv.Atoi(hostIDtmp)
	if err != nil {
		log.Debugf("host id: %v", hostIDtmp)
		h.JSONR(c, badstatus, err)
		return
	}

	ss, err := hbs.GetHostStrategies(hostID)
	if err != nil {
		log.Error(err)
	}
	h.JSONR(c, ss)
	return
}
