package dataConfig

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig"
	dataConfigReq "github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig/request"
)

type MarketService struct {
}

// CreateMarket 创建市场记录
// Author [piexlmax](https://github.com/piexlmax)
func (marketService *MarketService) CreateMarket(market *dataConfig.Market) (err error) {
	err = global.GVA_DB.Create(market).Error
	return err
}

// DeleteMarket 删除市场记录
// Author [piexlmax](https://github.com/piexlmax)
func (marketService *MarketService) DeleteMarket(ID string) (err error) {
	err = global.GVA_DB.Delete(&dataConfig.Market{}, "id = ?", ID).Error
	return err
}

// DeleteMarketByIds 批量删除市场记录
// Author [piexlmax](https://github.com/piexlmax)
func (marketService *MarketService) DeleteMarketByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]dataConfig.Market{}, "id in ?", IDs).Error
	return err
}

// UpdateMarket 更新市场记录
// Author [piexlmax](https://github.com/piexlmax)
func (marketService *MarketService) UpdateMarket(market dataConfig.Market) (err error) {
	err = global.GVA_DB.Model(&dataConfig.Market{}).Where("id = ?", market.ID).Updates(&market).Error
	return err
}

// GetMarket 根据ID获取市场记录
// Author [piexlmax](https://github.com/piexlmax)
func (marketService *MarketService) GetMarket(ID string) (market dataConfig.Market, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&market).Error
	return
}

// GetMarketInfoList 分页获取市场记录
// Author [piexlmax](https://github.com/piexlmax)
func (marketService *MarketService) GetMarketInfoList(info dataConfigReq.MarketSearch) (list []dataConfig.Market, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dataConfig.Market{})
	var markets []dataConfig.Market
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.MarketName != "" {
		db = db.Where("market_name LIKE ?", "%"+info.MarketName+"%")
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
	}
	if info.Remarks != "" {
		db = db.Where("remarks LIKE ?", "%"+info.Remarks+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["market_name"] = true
	orderMap["address"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&markets).Error
	return markets, total, err
}
