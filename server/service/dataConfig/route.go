package dataConfig

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig"
	dataConfigReq "github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig/request"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var currentStalls []dataConfig.Stall
		if TxErr := tx.Model(&dataConfig.Stall{}).Where("route_id = ?", ID).Find(&currentStalls).Error; err != nil {
			return TxErr
		}
		// 获取当前Stall的ID集合
		currentStallIDs := make([]uint, len(currentStalls))
		for i, stall := range currentStalls {
			currentStallIDs[i] = stall.ID
		}

		// 把currentStalls里的数据 route_id 置为null
		if len(currentStallIDs) > 0 {
			if TxErr := tx.Model(&dataConfig.Stall{}).Where("id IN ?", currentStallIDs).Update("route_id", 0).Error; err != nil {
				return TxErr
			}
		}
		TxErr := tx.Delete(&dataConfig.Route{}, "id = ?", ID).Error
		if TxErr != nil {
			return TxErr
		}
		return nil
	})
}

// DeleteRouteByIds 批量删除路线记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeService *RouteService) DeleteRouteByIds(IDs []string) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var currentStalls []dataConfig.Stall
		if TxErr := tx.Model(&dataConfig.Stall{}).Where("route_id IN ?", IDs).Find(&currentStalls).Error; err != nil {
			return TxErr
		}
		// 获取当前Stall的ID集合
		currentStallIDs := make([]uint, len(currentStalls))
		for i, stall := range currentStalls {
			currentStallIDs[i] = stall.ID
		}

		// 把currentStalls里的数据 route_id 置为null
		if len(currentStallIDs) > 0 {
			if TxErr := tx.Model(&dataConfig.Stall{}).Where("id IN ?", currentStallIDs).Update("route_id", 0).Error; err != nil {
				return TxErr
			}
		}
		TxErr := tx.Delete(&dataConfig.Route{}, "id IN ?", IDs).Error
		if TxErr != nil {
			return TxErr
		}
		return nil
	})
}

// UpdateRoute 更新路线记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeService *RouteService) UpdateRoute(route dataConfig.Route) (err error) {

	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var currentStalls []dataConfig.Stall
		// route stall表 锁死 查询也锁了
		//if TxErr := tx.Exec("LOCK TABLES stall WRITE, route WRITE").Error; TxErr != nil {
		//	return TxErr
		//}
		//defer tx.Exec("UNLOCK TABLES") // 在事务结束前释放表锁

		// 锁stall表 查询不锁 还会把事务内的所有表都锁住
		if TxErr := tx.Model(&dataConfig.Stall{}).Clauses(clause.Locking{Strength: "UPDATE"}).Where("route_id = ?", route.ID).Find(&currentStalls).Error; err != nil {
			return TxErr
		}

		// 锁stall表 查询不锁 只锁查询到的数据 不会锁住其他数据
		//if TxErr := tx.Model(&dataConfig.Stall{}).Set("gorm:query_option", "FOR UPDATE").Where("route_id = ?", route.ID).Find(&currentStalls).Error; err != nil {
		//	return TxErr
		//}

		// 获取当前Stall的ID集合
		currentStallIDs := make([]uint, len(currentStalls))
		for i, stall := range currentStalls {
			currentStallIDs[i] = stall.ID
		}

		// 把currentStalls里的数据 route_id 置为null
		if len(currentStallIDs) > 0 {
			if TxErr := tx.Model(&dataConfig.Stall{}).Where("id IN ?", currentStallIDs).Update("route_id", 0).Error; TxErr != nil {
				return TxErr
			}
		}

		newStalls := route.Stalls
		newStallsIDs := make([]uint, len(newStalls))
		for i, stall := range newStalls {
			newStallsIDs[i] = stall.ID
		}
		if len(newStallsIDs) > 0 {
			if TxErr := tx.Model(&dataConfig.Stall{}).Where("id IN ?", newStallsIDs).Update("route_id", route.ID).Error; TxErr != nil {
				return TxErr
			}
		}
		// 锁住 route查询到的数据
		TxErr := tx.Model(&dataConfig.Route{}).Where("id = ?", route.ID).Updates(&route).Error
		//TxErr := tx.Model(&dataConfig.Route{}).Set("gorm:query_option", "FOR UPDATE").Where("id = ?", route.ID).Updates(&route).Error
		//time.Sleep(19 * time.Second)

		if TxErr != nil {
			return TxErr
		}
		return nil
	})
	/**
	err = global.GVA_DB.Model(&dataConfig.Route{}).Where("id = ?", route.ID).Updates(&route).Error
	return err
	*/
}

// GetRoute 根据ID获取路线记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeService *RouteService) GetRoute(ID string) (route dataConfig.Route, err error) {
	err = global.GVA_DB.Preload("Stalls").Preload("User").Where("id = ?", ID).First(&route).Error
	return
}

// GetRouteInfoList 分页获取路线记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeService *RouteService) GetRouteInfoList(info dataConfigReq.RouteSearch) (list []dataConfig.Route, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dataConfig.Route{}).Preload("Stalls").Preload("User")
	var routes []dataConfig.Route
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.RouteName != "" {
		db = db.Where("route_name LIKE ?", "%"+info.RouteName+"%")
	}
	if info.Urgent != nil {
		db = db.Where("urgent = ?", info.Urgent)
	}

	if info.UserIds != "" && info.UserIds != "[]" {
		// UserIds是一个字符串[10,11] 转换成int类型数组
		var userIds []int
		err := json.Unmarshal([]byte(info.UserIds), &userIds)
		if err != nil {
			return routes, 0, err
		}
		if len(userIds) > 0 {
			db = db.Where("user_id IN ?", userIds)
		}
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
