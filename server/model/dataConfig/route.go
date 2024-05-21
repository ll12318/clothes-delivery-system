// 自动生成模板Route
package dataConfig

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// 路线 结构体  Route
type Route struct {
	global.GVA_MODEL
	RouteName string         `json:"routeName" form:"routeName" gorm:"column:route_name;comment:;" binding:"required"` //路线名称
	Remarks   string         `json:"remarks" form:"remarks" gorm:"column:remarks;comment:;"`                           //备注
	Stalls    []Stall        `json:"stalls" gorm:"foreignKey:RouteId;references:ID"`
	User      system.SysUser `json:"user" gorm:"foreignKey:UserId;references:ID"`
	UserId    uint           `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`
}

// TableName 路线 Route自定义表名 route
func (Route) TableName() string {
	return "route"
}
