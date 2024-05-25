// 自动生成模板Market
package dataConfig

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 市场 结构体  Market
type Market struct {
	global.GVA_MODEL
	MarketName string `json:"marketName" form:"marketName" gorm:"column:market_name;comment:;"` //市场名
	Address    string `json:"address" form:"address" gorm:"column:address;comment:;"`           //地址
	Remarks    string `json:"remarks" form:"remarks" gorm:"column:remarks;comment:;"`           //备注
}

// TableName 市场 Market自定义表名 market
func (Market) TableName() string {
	return "market"
}
