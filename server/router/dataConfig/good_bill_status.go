package dataConfig

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoodBillStatusRouter struct {
}

// InitGoodBillStatusRouter 初始化 货单状态 路由信息
func (s *GoodBillStatusRouter) InitGoodBillStatusRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	gbsRouter := Router.Group("gbs").Use(middleware.OperationRecord())
	gbsRouterWithoutRecord := Router.Group("gbs")
	gbsRouterWithoutAuth := PublicRouter.Group("gbs")

	var gbsApi = v1.ApiGroupApp.DataConfigApiGroup.GoodBillStatusApi
	{
		gbsRouter.POST("createGoodBillStatus", gbsApi.CreateGoodBillStatus)   // 新建货单状态
		gbsRouter.DELETE("deleteGoodBillStatus", gbsApi.DeleteGoodBillStatus) // 删除货单状态
		gbsRouter.DELETE("deleteGoodBillStatusByIds", gbsApi.DeleteGoodBillStatusByIds) // 批量删除货单状态
		gbsRouter.PUT("updateGoodBillStatus", gbsApi.UpdateGoodBillStatus)    // 更新货单状态
	}
	{
		gbsRouterWithoutRecord.GET("findGoodBillStatus", gbsApi.FindGoodBillStatus)        // 根据ID获取货单状态
		gbsRouterWithoutRecord.GET("getGoodBillStatusList", gbsApi.GetGoodBillStatusList)  // 获取货单状态列表
	}
	{
	    gbsRouterWithoutAuth.GET("getGoodBillStatusPublic", gbsApi.GetGoodBillStatusPublic)  // 获取货单状态列表
	}
}
