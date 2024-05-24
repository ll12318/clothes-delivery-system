// 自动生成模板Stall
package dataConfig

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 档口 结构体  Stall
type Stall struct {
	global.GVA_MODEL
	Stall       string `json:"stall" form:"stall" gorm:"column:stall;comment:;" binding:"required"`                    //档口名称
	StallNumber string `json:"stallNumber" form:"stallNumber" gorm:"column:stall_number;comment:;" binding:"required"` //档口号
	Remarks     string `json:"remarks" form:"remarks" gorm:"column:remarks;comment:;"`                                 //备注
	Market      Market `json:"market" form:"market" gorm:"foreignKey:MarketId;references:ID"`
	MarketId    uint   `json:"marketId" form:"marketId" gorm:"column:market_id;comment:;"`
	RouteId     uint   `json:"routeId" form:"routeId" gorm:"column:route_id;comment:;"`
	Urgent      *bool  `json:"urgent" form:"urgent" gorm:"column:urgent;comment:是否加急;"` // 是否紧急
	Price       int    `json:"price" form:"price" gorm:"column:price;comment: 档口价格;"`   // 档口价格
}

// TableName 档口 Stall自定义表名 stall
func (Stall) TableName() string {
	return "stall"
}
