package dataConfig

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MarketRouter struct {
}

// InitMarketRouter 初始化 市场 路由信息
func (s *MarketRouter) InitMarketRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	marketRouter := Router.Group("market").Use(middleware.OperationRecord())
	marketRouterWithoutRecord := Router.Group("market")
	marketRouterWithoutAuth := PublicRouter.Group("market")

	var marketApi = v1.ApiGroupApp.DataConfigApiGroup.MarketApi
	{
		marketRouter.POST("createMarket", marketApi.CreateMarket)             // 新建市场
		marketRouter.DELETE("deleteMarket", marketApi.DeleteMarket)           // 删除市场
		marketRouter.DELETE("deleteMarketByIds", marketApi.DeleteMarketByIds) // 批量删除市场
		marketRouter.PUT("updateMarket", marketApi.UpdateMarket)              // 更新市场
	}
	{
		marketRouterWithoutRecord.GET("findMarket", marketApi.FindMarket)       // 根据ID获取市场
		marketRouterWithoutRecord.GET("getMarketList", marketApi.GetMarketList) // 获取市场列表
	}
	{
		marketRouterWithoutAuth.GET("getMarketPublic", marketApi.GetMarketPublic) // 获取市场列表
	}
}
