package transaction

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TransactionDetailsRouter struct {
}

// InitTransactionDetailsRouter 初始化 交易详情 路由信息
func (s *TransactionDetailsRouter) InitTransactionDetailsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	tdRouter := Router.Group("td").Use(middleware.OperationRecord())
	tdRouterWithoutRecord := Router.Group("td")
	tdRouterWithoutAuth := PublicRouter.Group("td")

	var tdApi = v1.ApiGroupApp.TransactionApiGroup.TransactionDetailsApi
	{
		tdRouter.POST("createTransactionDetails", tdApi.CreateTransactionDetails)   // 新建交易详情
		tdRouter.DELETE("deleteTransactionDetails", tdApi.DeleteTransactionDetails) // 删除交易详情
		tdRouter.DELETE("deleteTransactionDetailsByIds", tdApi.DeleteTransactionDetailsByIds) // 批量删除交易详情
		tdRouter.PUT("updateTransactionDetails", tdApi.UpdateTransactionDetails)    // 更新交易详情
	}
	{
		tdRouterWithoutRecord.GET("findTransactionDetails", tdApi.FindTransactionDetails)        // 根据ID获取交易详情
		tdRouterWithoutRecord.GET("getTransactionDetailsList", tdApi.GetTransactionDetailsList)  // 获取交易详情列表
	}
	{
	    tdRouterWithoutAuth.GET("getTransactionDetailsPublic", tdApi.GetTransactionDetailsPublic)  // 获取交易详情列表
	}
}
