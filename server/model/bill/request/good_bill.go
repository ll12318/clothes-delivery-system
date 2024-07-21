package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type GoodBillSearch struct {
	StartCreatedAt   *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt     *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	CreatedBy        uint       `json:"createdBy" form:"createdBy"`
	TakeGoodPeopleId uint       `json:"takeGoodPeopleId" form:"takeGoodPeopleId"`
	FinishStatus     *bool      `json:"finishStatus" form:"finishStatus"`
	DriverVerify     *bool      `json:"driverVerify" form:"driverVerify"` //
	Device           string     `json:"device" form:"device"`             // 0 小程序 1 网页端
	request.PageInfo
}

// GoodBillMarketListSearch
type GoodBillMarketListSearch struct {
	FinishStatus *bool `json:"finishStatus" form:"finishStatus"`
	DriverVerify *bool `json:"driverVerify" form:"driverVerify"`
	MarketId     uint  `json:"marketId" form:"marketId"`
}
