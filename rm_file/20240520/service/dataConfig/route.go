package dataConfig

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig"
    dataConfigReq "github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig/request"
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
func (routeService *RouteService)DeleteRoute(ID string) (err error) {
	err = global.GVA_DB.Delete(&dataConfig.Route{},"id = ?",ID).Error
	return err
}

// DeleteRouteByIds 批量删除路线记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeService *RouteService)DeleteRouteByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]dataConfig.Route{},"id in ?",IDs).Error
	return err
}

// UpdateRoute 更新路线记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeService *RouteService)UpdateRoute(route dataConfig.Route) (err error) {
	err = global.GVA_DB.Model(&dataConfig.Route{}).Where("id = ?",route.ID).Updates(&route).Error
	return err
}

// GetRoute 根据ID获取路线记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeService *RouteService)GetRoute(ID string) (route dataConfig.Route, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&route).Error
	return
}

// GetRouteInfoList 分页获取路线记录
// Author [piexlmax](https://github.com/piexlmax)
func (routeService *RouteService)GetRouteInfoList(info dataConfigReq.RouteSearch) (list []dataConfig.Route, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&dataConfig.Route{})
    var routes []dataConfig.Route
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.RouteName != "" {
        db = db.Where("route_name LIKE ?","%"+ info.RouteName+"%")
    }
    if info.Stalls != nil {
        db = db.Where("stalls = ?",info.Stalls)
    }
	err = db.Count(&total).Error
	if err!=nil {
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
	return  routes, total, err
}
func (routeService *RouteService)GetRouteDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)
	
	   stalls := make([]map[string]any, 0)
       global.GVA_DB.Table("stall").Select("stalls as label,stalls as value").Scan(&stalls)
	   res["stalls"] = stalls
	return
}