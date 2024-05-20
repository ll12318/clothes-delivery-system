package dataConfig

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig"
	dataConfigReq "github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig/request"
	"gorm.io/gorm"
)

type RouteService struct {
}

// CreateRoute 创建路线记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeService *RouteService) CreateRoute(route *dataConfig.Route) (err error) {
	err = global.GVA_DB.Create(route).Error
	return err
}

// DeleteRoute 删除路线记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeService *RouteService) DeleteRoute(ID string) (err error) {
	err = global.GVA_DB.Delete(&dataConfig.Route{}, "id = ?", ID).Error
	return err
}

// DeleteRouteByIds 批量删除路线记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeService *RouteService) DeleteRouteByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]dataConfig.Route{}, "id in ?", IDs).Error
	return err
}

// UpdateRoute 更新路线记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeService *RouteService) UpdateRoute(route dataConfig.Route) (err error) {

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) (err error) {
		var currentStalls []dataConfig.Stall
		if err := tx.Model(&dataConfig.Stall{}).Where("route_id = ?", route.ID).Find(&currentStalls).Error; err != nil {
			return err
		}

		// 获取当前Stall的ID集合
		currentStallIDs := make([]uint, len(currentStalls))
		for i, stall := range currentStalls {
			currentStallIDs[i] = stall.ID
		}

		// 把currentStalls里的数据 route_id 置为null
		if len(currentStallIDs) > 0 {
			if err := tx.Model(&dataConfig.Stall{}).Where("id IN ?", currentStallIDs).Update("route_id", 0).Error; err != nil {
				return err
			}
		}

		newStalls := route.Stalls
		newStallsIDs := make([]uint, len(newStalls))
		for i, stall := range newStalls {
			newStallsIDs[i] = stall.ID
		}
		if len(newStallsIDs) > 0 {
			if err := tx.Model(&dataConfig.Stall{}).Where("id IN ?", newStallsIDs).Update("route_id", route.ID).Error; err != nil {
				return err
			}
		}

		err = tx.Model(&dataConfig.Route{}).Where("id = ?", route.ID).Updates(&route).Error

		return err
	})
	return err
	/**
	err = global.GVA_DB.Model(&dataConfig.Route{}).Where("id = ?", route.ID).Updates(&route).Error
	return err
	*/
}

// GetRoute 根据ID获取路线记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeService *RouteService) GetRoute(ID string) (route dataConfig.Route, err error) {
	err = global.GVA_DB.Preload("Stalls").Where("id = ?", ID).First(&route).Error
	return
}

// GetRouteInfoList 分页获取路线记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeService *RouteService) GetRouteInfoList(info dataConfigReq.RouteSearch) (list []dataConfig.Route, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dataConfig.Route{}).Preload("Stalls")
	var routes []dataConfig.Route
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.RouteName != "" {
		db = db.Where("route_name LIKE ?", "%"+info.RouteName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["route_name"] = true
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

	err = db.Find(&routes).Error
	return routes, total, err
}
