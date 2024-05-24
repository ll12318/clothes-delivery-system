package dataConfig

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig"
	dataConfigReq "github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig/request"
)

type StallService struct {
}

// CreateStall 创建档口记录
// Author [piexlmax](https://github.com/piexlmax)
func (stallService *StallService) CreateStall(stall *dataConfig.Stall) (err error) {
	err = global.GVA_DB.Create(stall).Error
	return err
}

// DeleteStall 删除档口记录
// Author [piexlmax](https://github.com/piexlmax)
func (stallService *StallService) DeleteStall(ID string) (err error) {
	err = global.GVA_DB.Delete(&dataConfig.Stall{}, "id = ?", ID).Error
	return err
}

// DeleteStallByIds 批量删除档口记录
// Author [piexlmax](https://github.com/piexlmax)
func (stallService *StallService) DeleteStallByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]dataConfig.Stall{}, "id in ?", IDs).Error
	return err
}

// UpdateStall 更新档口记录
// Author [piexlmax](https://github.com/piexlmax)
func (stallService *StallService) UpdateStall(stall dataConfig.Stall) (err error) {
	err = global.GVA_DB.Model(&dataConfig.Stall{}).Where("id = ?", stall.ID).Updates(&stall).Error
	return err
}

// GetStall 根据ID获取档口记录
// Author [piexlmax](https://github.com/piexlmax)
func (stallService *StallService) GetStall(ID string) (stall dataConfig.Stall, err error) {
	err = global.GVA_DB.Preload("Market").Where("id = ?", ID).First(&stall).Error
	return
}

// GetStallInfoList 分页获取档口记录
// Author [piexlmax](https://github.com/piexlmax)
func (stallService *StallService) GetStallInfoList(info dataConfigReq.StallSearch) (list []dataConfig.Stall, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dataConfig.Stall{}).Preload("Market")
	var stalls []dataConfig.Stall
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	if info.FilterOccupancy == true {
		db = db.Where("route_id = ?", 0)
	} else {
		if info.MarketId != 0 {
			db = db.Where("market_id = ?", info.MarketId)
		}
	}

	if info.Urgent != nil {
		db = db.Where("urgent = ?", info.Urgent)
	}

	if info.Stall != "" {
		db = db.Where("stall LIKE ?", "%"+info.Stall+"%")
	}
	if info.StallNumber != "" {
		db = db.Where("stall_number LIKE ?", "%"+info.StallNumber+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["stall"] = true
	orderMap["stall_number"] = true
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

	err = db.Find(&stalls).Error
	return stalls, total, err
}
