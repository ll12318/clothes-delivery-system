// 自动生成模板GoodBillStatus
package dataConfig

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 货单状态 结构体  GoodBillStatus
type GoodBillStatus struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:;"`                                 //名称
	Color     string `json:"color" form:"color" gorm:"default:#3B82F6;column:color;comment:;"`              //背景颜色
	FontColor string `json:"fontColor" form:"fontColor" gorm:"default:#FFFFFF;column:font_color;comment:;"` // 字体颜色
}

// TableName 货单状态 GoodBillStatus自定义表名 good_bill_status
func (GoodBillStatus) TableName() string {
	return "good_bill_status"
}
