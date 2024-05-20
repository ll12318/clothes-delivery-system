// 自动生成模板Route
package dataConfig

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	"gorm.io/datatypes"
)

// 路线 结构体  Route
type Route struct {
 global.GVA_MODEL 
      RouteName  string `json:"routeName" form:"routeName" gorm:"column:route_name;comment:;" binding:"required"`  //路线名称 
      Remarks  string `json:"remarks" form:"remarks" gorm:"column:remarks;comment:;"`  //备注 
      Stalls  datatypes.JSON `json:"stalls" form:"stalls" gorm:"column:stalls;comment:;type:text;"`  //档口 
}


// TableName 路线 Route自定义表名 route
func (Route) TableName() string {
  return "route"
}

