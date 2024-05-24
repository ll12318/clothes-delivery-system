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
	CreatedByUserInfo   system.SysUser               `json:"-" form:"createdByUserInfo" gorm:"foreignKey:CreatedBy;references:ID"`
	CreatedBy           uint                         `json:"-" gorm:"column:created_by;comment:创建者"`
	UpdatedBy           uint                         `json:"-" gorm:"column:updated_by;comment:更新者"`
	DeletedBy           uint                         `json:"-" gorm:"column:deleted_by;comment:删除者"`
	CreatedBySimpleUser response.CreatedBySimpleUser `json:"createdBySimpleUser"  gorm:"-"`
	Stall               dataConfig.Stall             `json:"stall" gorm:"foreignKey:StallId;references:ID"`
	StallId             uint                         `json:"stallId" form:"stallId" gorm:"column:stall_id;comment:;"`
	Urgent              *bool                        `json:"urgent" form:"urgent" gorm:"column:urgent;comment:是否加急;"` // 是否紧急
	TakeGoodPeople      system.SysUser               `json:"takeGoodPeople" form:"takeGoodPeople" gorm:"foreignKey:TakeGoodPeopleId;references:ID"`
	TakeGoodPeopleId    uint                         `json:"takeGoodPeopleId" form:"takeGoodPeopleId" gorm:"column:take_good_people_id;comment: 拿货人id;"`
	TakeGoodNum         int                          `json:"takeGoodNum" form:"takeGoodNum" gorm:"column:take_good_num;comment:拿货数量;"` // 拿货数量
}

// TableName 货单 GoodBill自定义表名 good_bill
func (GoodBill) TableName() string {
	return "good_bill"
}
