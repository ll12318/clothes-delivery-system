package bill

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type GoodBillRouter struct {
}

// InitGoodBillRouter 初始化 货单 路由信息
func (s *GoodBillRouter) InitGoodBillRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	gbRouter := Router.Group("gb").Use(middleware.OperationRecord())
	gbRouterWithoutRecord := Router.Group("gb")
	gbRouterWithoutAuth := PublicRouter.Group("gb")

	var gbApi = v1.ApiGroupApp.BillApiGroup.GoodBillApi
	{
		gbRouter.POST("createGoodBill", gbApi.CreateGoodBill)             // 新建货单
		gbRouter.DELETE("deleteGoodBill", gbApi.DeleteGoodBill)           // 删除货单
		gbRouter.DELETE("deleteGoodBillByIds", gbApi.DeleteGoodBillByIds) // 批量删除货单
		gbRouter.PUT("updateGoodBill", gbApi.UpdateGoodBill)              // 更新货单
	}
	{
		gbRouterWithoutRecord.GET("findGoodBill", gbApi.FindGoodBill)       // 根据ID获取货单
		gbRouterWithoutRecord.GET("getGoodBillList", gbApi.GetGoodBillList) // 获取货单列表
		// 司机端获取货单列表
		gbRouterWithoutRecord.GET("getGoodBillMarketListByDriver", gbApi.GetGoodBillMarketListByDriver) // 获取货单列表

		// 根据marketId获取货单列表
		gbRouterWithoutRecord.GET("getGoodBillListByMarketId", gbApi.GetGoodBillListByMarketId) // 获取货单列表
	}
	{
		gbRouterWithoutAuth.GET("getGoodBillPublic", gbApi.GetGoodBillPublic) // 获取货单列表
	}
}
