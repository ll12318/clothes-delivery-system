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

type StallApi struct {
}

var stallService = service.ServiceGroupApp.DataConfigServiceGroup.StallService

// CreateStall 创建档口
// @Tags Stall
// @Summary 创建档口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataConfig.Stall true "创建档口"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /stall/createStall [post]
func (stallApi *StallApi) CreateStall(c *gin.Context) {
	var stall dataConfig.Stall
	err := c.ShouldBindJSON(&stall)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := stallService.CreateStall(&stall); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteStall 删除档口
// @Tags Stall
// @Summary 删除档口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataConfig.Stall true "删除档口"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /stall/deleteStall [delete]
func (stallApi *StallApi) DeleteStall(c *gin.Context) {
	ID := c.Query("ID")
	if err := stallService.DeleteStall(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteStallByIds 批量删除档口
// @Tags Stall
// @Summary 批量删除档口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /stall/deleteStallByIds [delete]
func (stallApi *StallApi) DeleteStallByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := stallService.DeleteStallByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateStall 更新档口
// @Tags Stall
// @Summary 更新档口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataConfig.Stall true "更新档口"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /stall/updateStall [put]
func (stallApi *StallApi) UpdateStall(c *gin.Context) {
	var stall dataConfig.Stall
	err := c.ShouldBindJSON(&stall)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := stallService.UpdateStall(stall); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindStall 用id查询档口
// @Tags Stall
// @Summary 用id查询档口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dataConfig.Stall true "用id查询档口"
// @Success 200 {object} response.Response{data=object{restall=dataConfig.Stall},msg=string} "查询成功"
// @Router /stall/findStall [get]
func (stallApi *StallApi) FindStall(c *gin.Context) {
	ID := c.Query("ID")
	if restall, err := stallService.GetStall(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"restall": restall}, c)
	}
}

// GetStallList 分页获取档口列表
// @Tags Stall
// @Summary 分页获取档口列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dataConfigReq.StallSearch true "分页获取档口列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /stall/getStallList [get]
func (stallApi *StallApi) GetStallList(c *gin.Context) {
	var pageInfo dataConfigReq.StallSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := stallService.GetStallInfoList(pageInfo); err != nil {
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

// GetStallPublic 不需要鉴权的档口接口
// @Tags Stall
// @Summary 不需要鉴权的档口接口
// @accept application/json
// @Produce application/json
// @Param data query dataConfigReq.StallSearch true "分页获取档口列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /stall/getStallPublic [get]
func (stallApi *StallApi) GetStallPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的档口接口信息",
	}, "获取成功", c)
}
