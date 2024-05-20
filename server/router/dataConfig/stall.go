package dataConfig

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type StallRouter struct {
}

// InitStallRouter 初始化 档口 路由信息
func (s *StallRouter) InitStallRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	stallRouter := Router.Group("stall").Use(middleware.OperationRecord())
	stallRouterWithoutRecord := Router.Group("stall")
	stallRouterWithoutAuth := PublicRouter.Group("stall")

	var stallApi = v1.ApiGroupApp.DataConfigApiGroup.StallApi
	{
		stallRouter.POST("createStall", stallApi.CreateStall)             // 新建档口
		stallRouter.DELETE("deleteStall", stallApi.DeleteStall)           // 删除档口
		stallRouter.DELETE("deleteStallByIds", stallApi.DeleteStallByIds) // 批量删除档口
		stallRouter.PUT("updateStall", stallApi.UpdateStall)              // 更新档口
	}
	{
		stallRouterWithoutRecord.GET("findStall", stallApi.FindStall)       // 根据ID获取档口
		stallRouterWithoutRecord.GET("getStallList", stallApi.GetStallList) // 获取档口列表
	}
	{
		stallRouterWithoutAuth.GET("getStallPublic", stallApi.GetStallPublic) // 获取档口列表
	}
}
