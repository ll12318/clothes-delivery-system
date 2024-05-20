package dataConfig

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig"
	dataConfigReq "github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RouteApi struct {
}

var routeService = service.ServiceGroupApp.DataConfigServiceGroup.RouteService

// CreateRoute 创建路线
// @Tags Route
// @Summary 创建路线
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataConfig.Route true "创建路线"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /route/createRoute [post]
func (routeApi *RouteApi) CreateRoute(c *gin.Context) {
	var route dataConfig.Route
	err := c.ShouldBindJSON(&route)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := routeService.CreateRoute(&route); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRoute 删除路线
// @Tags Route
// @Summary 删除路线
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataConfig.Route true "删除路线"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /route/deleteRoute [delete]
func (routeApi *RouteApi) DeleteRoute(c *gin.Context) {
	ID := c.Query("ID")
	if err := routeService.DeleteRoute(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRouteByIds 批量删除路线
// @Tags Route
// @Summary 批量删除路线
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /route/deleteRouteByIds [delete]
func (routeApi *RouteApi) DeleteRouteByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := routeService.DeleteRouteByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRoute 更新路线
// @Tags Route
// @Summary 更新路线
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataConfig.Route true "更新路线"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /route/updateRoute [put]
func (routeApi *RouteApi) UpdateRoute(c *gin.Context) {
	var route dataConfig.Route
	err := c.ShouldBindJSON(&route)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := routeService.UpdateRoute(route); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRoute 用id查询路线
// @Tags Route
// @Summary 用id查询路线
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dataConfig.Route true "用id查询路线"
// @Success 200 {object} response.Response{data=object{reroute=dataConfig.Route},msg=string} "查询成功"
// @Router /route/findRoute [get]
func (routeApi *RouteApi) FindRoute(c *gin.Context) {
	ID := c.Query("ID")
	if reroute, err := routeService.GetRoute(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reroute": reroute}, c)
	}
}

// GetRouteList 分页获取路线列表
// @Tags Route
// @Summary 分页获取路线列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dataConfigReq.RouteSearch true "分页获取路线列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /route/getRouteList [get]
func (routeApi *RouteApi) GetRouteList(c *gin.Context) {
	var pageInfo dataConfigReq.RouteSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := routeService.GetRouteInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetRoutePublic 不需要鉴权的路线接口
// @Tags Route
// @Summary 不需要鉴权的路线接口
// @accept application/json
// @Produce application/json
// @Param data query dataConfigReq.RouteSearch true "分页获取路线列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /route/getRoutePublic [get]
func (routeApi *RouteApi) GetRoutePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的路线接口信息",
	}, "获取成功", c)
}
