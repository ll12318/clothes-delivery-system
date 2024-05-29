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

type MarketApi struct {
}

var marketService = service.ServiceGroupApp.DataConfigServiceGroup.MarketService

// CreateMarket 创建市场
// @Tags Market
// @Summary 创建市场
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataConfig.Market true "创建市场"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /market/createMarket [post]
func (marketApi *MarketApi) CreateMarket(c *gin.Context) {
	var market dataConfig.Market
	err := c.ShouldBindJSON(&market)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := marketService.CreateMarket(&market); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMarket 删除市场
// @Tags Market
// @Summary 删除市场
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataConfig.Market true "删除市场"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /market/deleteMarket [delete]
func (marketApi *MarketApi) DeleteMarket(c *gin.Context) {
	ID := c.Query("ID")
	if err := marketService.DeleteMarket(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMarketByIds 批量删除市场
// @Tags Market
// @Summary 批量删除市场
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /market/deleteMarketByIds [delete]
func (marketApi *MarketApi) DeleteMarketByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := marketService.DeleteMarketByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMarket 更新市场
// @Tags Market
// @Summary 更新市场
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataConfig.Market true "更新市场"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /market/updateMarket [put]
func (marketApi *MarketApi) UpdateMarket(c *gin.Context) {
	var market dataConfig.Market
	err := c.ShouldBindJSON(&market)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := marketService.UpdateMarket(market); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMarket 用id查询市场
// @Tags Market
// @Summary 用id查询市场
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dataConfig.Market true "用id查询市场"
// @Success 200 {object} response.Response{data=object{remarket=dataConfig.Market},msg=string} "查询成功"
// @Router /market/findMarket [get]
func (marketApi *MarketApi) FindMarket(c *gin.Context) {
	ID := c.Query("ID")
	if remarket, err := marketService.GetMarket(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remarket": remarket}, c)
	}
}

// GetMarketList 分页获取市场列表
// @Tags Market
// @Summary 分页获取市场列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dataConfigReq.MarketSearch true "分页获取市场列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /market/getMarketList [get]
func (marketApi *MarketApi) GetMarketList(c *gin.Context) {
	var pageInfo dataConfigReq.MarketSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := marketService.GetMarketInfoList(pageInfo); err != nil {
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

// GetMarketPublic 不需要鉴权的市场接口
// @Tags Market
// @Summary 不需要鉴权的市场接口
// @accept application/json
// @Produce application/json
// @Param data query dataConfigReq.MarketSearch true "分页获取市场列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /market/getMarketPublic [get]
func (marketApi *MarketApi) GetMarketPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的市场接口信息",
	}, "获取成功", c)
}
