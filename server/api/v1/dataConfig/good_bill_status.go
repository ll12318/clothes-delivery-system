package dataConfig

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig"
    dataConfigReq "github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type GoodBillStatusApi struct {
}

var gbsService = service.ServiceGroupApp.DataConfigServiceGroup.GoodBillStatusService


// CreateGoodBillStatus 创建货单状态
// @Tags GoodBillStatus
// @Summary 创建货单状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataConfig.GoodBillStatus true "创建货单状态"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /gbs/createGoodBillStatus [post]
func (gbsApi *GoodBillStatusApi) CreateGoodBillStatus(c *gin.Context) {
	var gbs dataConfig.GoodBillStatus
	err := c.ShouldBindJSON(&gbs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := gbsService.CreateGoodBillStatus(&gbs); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteGoodBillStatus 删除货单状态
// @Tags GoodBillStatus
// @Summary 删除货单状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataConfig.GoodBillStatus true "删除货单状态"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /gbs/deleteGoodBillStatus [delete]
func (gbsApi *GoodBillStatusApi) DeleteGoodBillStatus(c *gin.Context) {
	ID := c.Query("ID")
	if err := gbsService.DeleteGoodBillStatus(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoodBillStatusByIds 批量删除货单状态
// @Tags GoodBillStatus
// @Summary 批量删除货单状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /gbs/deleteGoodBillStatusByIds [delete]
func (gbsApi *GoodBillStatusApi) DeleteGoodBillStatusByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := gbsService.DeleteGoodBillStatusByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoodBillStatus 更新货单状态
// @Tags GoodBillStatus
// @Summary 更新货单状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dataConfig.GoodBillStatus true "更新货单状态"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /gbs/updateGoodBillStatus [put]
func (gbsApi *GoodBillStatusApi) UpdateGoodBillStatus(c *gin.Context) {
	var gbs dataConfig.GoodBillStatus
	err := c.ShouldBindJSON(&gbs)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := gbsService.UpdateGoodBillStatus(gbs); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGoodBillStatus 用id查询货单状态
// @Tags GoodBillStatus
// @Summary 用id查询货单状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dataConfig.GoodBillStatus true "用id查询货单状态"
// @Success 200 {object} response.Response{data=object{regbs=dataConfig.GoodBillStatus},msg=string} "查询成功"
// @Router /gbs/findGoodBillStatus [get]
func (gbsApi *GoodBillStatusApi) FindGoodBillStatus(c *gin.Context) {
	ID := c.Query("ID")
	if regbs, err := gbsService.GetGoodBillStatus(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regbs": regbs}, c)
	}
}

// GetGoodBillStatusList 分页获取货单状态列表
// @Tags GoodBillStatus
// @Summary 分页获取货单状态列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dataConfigReq.GoodBillStatusSearch true "分页获取货单状态列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /gbs/getGoodBillStatusList [get]
func (gbsApi *GoodBillStatusApi) GetGoodBillStatusList(c *gin.Context) {
	var pageInfo dataConfigReq.GoodBillStatusSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := gbsService.GetGoodBillStatusInfoList(pageInfo); err != nil {
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

// GetGoodBillStatusPublic 不需要鉴权的货单状态接口
// @Tags GoodBillStatus
// @Summary 不需要鉴权的货单状态接口
// @accept application/json
// @Produce application/json
// @Param data query dataConfigReq.GoodBillStatusSearch true "分页获取货单状态列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /gbs/getGoodBillStatusPublic [get]
func (gbsApi *GoodBillStatusApi) GetGoodBillStatusPublic(c *gin.Context) {
    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的货单状态接口信息",
    }, "获取成功", c)
}
