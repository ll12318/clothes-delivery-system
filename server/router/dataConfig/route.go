package dataConfig

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RouteRouter struct {
}

// InitRouteRouter 初始化 路线 路由信息
func (s *RouteRouter) InitRouteRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	routeRouter := Router.Group("route").Use(middleware.OperationRecord())
	routeRouterWithoutRecord := Router.Group("route")
	routeRouterWithoutAuth := PublicRouter.Group("route")

	var routeApi = v1.ApiGroupApp.DataConfigApiGroup.RouteApi
	{
		routeRouter.POST("createRoute", routeApi.CreateRoute)             // 新建路线
		routeRouter.DELETE("deleteRoute", routeApi.DeleteRoute)           // 删除路线
		routeRouter.DELETE("deleteRouteByIds", routeApi.DeleteRouteByIds) // 批量删除路线
		routeRouter.PUT("updateRoute", routeApi.UpdateRoute)              // 更新路线
	}
	{
		routeRouterWithoutRecord.GET("findRoute", routeApi.FindRoute)       // 根据ID获取路线
		routeRouterWithoutRecord.GET("getRouteList", routeApi.GetRouteList) // 获取路线列表
	}
	{
		routeRouterWithoutAuth.GET("getRoutePublic", routeApi.GetRoutePublic) // 获取路线列表
	}
}
