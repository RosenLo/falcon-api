package falcon_portal

import (
	con "github.com/open-falcon/falcon-plus/modules/api/config"
)

// +---------+------------------+------+-----+---------+-------+
// | Field   | Type             | Null | Key | Default | Extra |
// +---------+------------------+------+-----+---------+-------+
// | grp_id  | int(10) unsigned | NO   | PRI | NULL    |       |
// | host_id | int(11)          | NO   | PRI | NULL    |       |
// +---------+------------------+------+-----+---------+-------+

type GrpHost struct {
	GrpID  int64 `json:"grp_id" gorm:"column:grp_id"`
	HostID int64 `json:"host_id" gorm:"column:host_id"`
}

func (this GrpHost) TableName() string {
	return "grp_host"
}

func (this GrpHost) Existing() (int64, bool) {
	db := con.Con()
	grp := GrpHost{}
	db.Falcon.Table(this.TableName()).Where("grp_id = ? and host_id = ?", this.GrpID, this.HostID).Scan(&grp)
	if grp.HostID != 0 {
		return grp.HostID, true
	} else {
		return 0, false
	}
}
