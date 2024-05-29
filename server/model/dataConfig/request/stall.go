package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type StallSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Stall           string `json:"stall" form:"stall" `
	StallNumber     string `json:"stallNumber" form:"stallNumber" `
	MarketId        uint   `json:"marketId" form:"marketId" `
	FilterOccupancy bool   `json:"filterOccupancy" form:"filterOccupancy"`
	Urgent          *bool  `json:"urgent" form:"urgent"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
