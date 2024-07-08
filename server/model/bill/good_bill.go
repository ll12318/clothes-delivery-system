// 自动生成模板GoodBill
package bill

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
)

// 货单 结构体  GoodBill
type GoodBill struct {
	global.GVA_MODEL
	Remarks             string                       `json:"remarks" form:"remarks" gorm:"column:remarks;comment:;"` //备注
	CreatedByUserInfo   system.SysUser               `json:"createdByUserInfo" form:"createdByUserInfo" gorm:"foreignKey:CreatedBy;references:ID"`
	CreatedBy           uint                         `json:"-" gorm:"column:created_by;comment:创建者"`
	UpdatedBy           uint                         `json:"-" gorm:"column:updated_by;comment:更新者"`
	DeletedBy           uint                         `json:"-" gorm:"column:deleted_by;comment:删除者"`
	CreatedBySimpleUser response.CreatedBySimpleUser `json:"createdBySimpleUser"  gorm:"-"`
	Stall               dataConfig.Stall             `json:"stall" gorm:"foreignKey:StallId;references:ID"`
	StallId             uint                         `json:"stallId" form:"stallId" gorm:"column:stall_id;comment:;"`
	Urgent              *bool                        `json:"urgent" form:"urgent" gorm:"column:urgent;comment:是否加急;"` // 是否紧急
	TakeGoodPeople      system.SysUser               `json:"takeGoodPeople" form:"takeGoodPeople" gorm:"foreignKey:TakeGoodPeopleId;references:ID"`
	TakeGoodPeopleId    uint                         `json:"takeGoodPeopleId" form:"takeGoodPeopleId" gorm:"column:take_good_people_id;comment: 拿货人id;"`
	TakeGoodNum         int                          `json:"takeGoodNum" form:"takeGoodNum" gorm:"column:take_good_num;comment:拿货数量;"`                   // 拿货数量
	FinishStatus        *bool                        `json:"finishStatus" form:"finishStatus" gorm:"column:finish_status;comment:完成状态;"`                 // 完成状态
	GoodBillStatus      dataConfig.GoodBillStatus    `json:"goodBillStatus" gorm:"foreignKey:GoodBillStatusId;references:ID"`                            // 货单状态
	GoodBillStatusId    uint                         `json:"goodBillStatusId" form:"goodBillStatusId" gorm:"column:good_bill_status_id;comment:货单状态id;"` // 货单状态id
	FinishTime          string                       `json:"finishTime" form:"finishTime" gorm:"column:finish_time;comment:完成时间;"`                       // 完成时间
	FinishPeople        system.SysUser               `json:"finishPeople" form:"finishPeople" gorm:"foreignKey:FinishPeopleId;references:ID"`
	FinishPeopleId      uint                         `json:"finishPeopleId" form:"finishPeopleId" gorm:"column:finish_people_id;comment:完成人id;"` // 完成人id
	Market              dataConfig.Market            `json:"market" gorm:"foreignKey:MarketId;references:ID"`
	MarketId            uint                         `json:"marketId" form:"marketId" gorm:"column:market_id;comment:市场id;"`
	Images              string                       `json:"images" form:"images" gorm:"column:images;comment:图片;type:longtext;"`
	OrderMessage        string                       `json:"orderMessage" form:"orderMessage" gorm:"column:order_message;comment:下单人留言;type:longtext;"`   // 下单人留言
	AdminMessage        string                       `json:"adminMessage" form:"adminMessage" gorm:"column:admin_message;comment:管理员留言;type:longtext;"`   // 管理员留言
	DriverMessage       string                       `json:"driverMessage" form:"driverMessage" gorm:"column:driver_message;comment:司机留言;type:longtext;"` // 司机留言
	DriverVerify        *bool                        `json:"driverVerify" form:"driverVerify" gorm:"column:driver_verify;comment:司机核实;"`
}

// TableName 货单 GoodBill自定义表名 good_bill
func (GoodBill) TableName() string {
	return "good_bill"
}
