// 自动生成模板TransactionDetails
package transaction

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// 交易详情 结构体  TransactionDetails
type TransactionDetails struct {
	global.GVA_MODEL
	PreTransactionAmount  *float64       `json:"preTransactionAmount" form:"preTransactionAmount" gorm:"column:pre_transaction_amount;comment:交易前金额;"`    //交易前金额
	TransactionAmount     *float64       `json:"transactionAmount" form:"transactionAmount" gorm:"column:transaction_amount;comment:交易金额;"`               //交易金额
	PostTransactionAmount *float64       `json:"postTransactionAmount" form:"postTransactionAmount" gorm:"column:post_transaction_amount;comment:交易后金额;"` //交易后金额
	User                  system.SysUser `json:"user" gorm:"foreignKey:UserId;references:ID"`                                                             // 用户
	UserId                uint           `json:"userId" form:"userId" gorm:"column:user_id;comment:;"`                                                    // 用户
}

// TableName 交易详情 TransactionDetails自定义表名 transaction_details
func (TransactionDetails) TableName() string {
	return "transaction_details"
}
