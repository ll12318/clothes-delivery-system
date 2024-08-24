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
	Device              string                       `json:"device" form:"device" gorm:"column:device;comment:下单设备 0 小程序 1 网页端;default:'1';"` // 下单设备 0 小程序 1 网页端
	BillNumber          string                       `json:"billNumber" form:"billNumber" gorm:"column:bill_number;comment:单据编号;"`

	// 微信订单号
	WechatOrderId string `json:"wechatOrderId" form:"wechatOrderId" gorm:"column:wechat_order_id;comment:微信订单号;"` // 微信订单号
	// 支付方式
	PayType string `json:"payType" form:"payType" gorm:"column:pay_type;comment:支付方式;default:'网页端'"`
	// 是否手动订单 0 自动 1 手动
	IsManual string `json:"isManual" form:"isManual" gorm:"column:is_manual;comment:是否手动订单;default:'0';"`
	// 是否支付完成 0 未完成 1 完成
	IsPay string `json:"isPay" form:"isPay" gorm:"column:is_pay;comment:是否支付完成;default:'0';"`
	// 售后状态 0 未售后 1 退款
	AfterSaleStatus string `json:"afterSaleStatus" form:"afterSaleStatus" gorm:"column:after_sale_status;comment:售后状态;default:'0';"`
	// 同意退款 0 不处理 1 同意 2 拒绝
	AgreeRefund string `json:"agreeRefund" form:"agreeRefund" gorm:"column:agree_refund;comment:同意退款;default:'0';"`
	// 退款完成 0 未完成 1 完成
	RefundStatus string `json:"refundStatus" form:"refundStatus" gorm:"column:refund_status;comment:退款完成;default:'0';"`
	// 退款订单号
	RefundOrderId string `json:"refundOrderId" form:"refundOrderId" gorm:"column:refund_order_id;comment:退款订单号;"`

	// 折扣后金额
	DiscountAmount float64 `json:"discountAmount" form:"discountAmount" gorm:"column:discount_amount;comment:折扣后金额;"`
	// 折扣率
	DiscountRate float64 `json:"discountRate" form:"discountRate" gorm:"column:discount_rate;comment:折扣率;"`
}

// TableName 货单 GoodBill自定义表名 good_bill
func (GoodBill) TableName() string {
	return "good_bill"
}
