package transaction

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/transaction"
	transactionReq "github.com/flipped-aurora/gin-vue-admin/server/model/transaction/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TransactionDetailsApi struct {
}

var tdService = service.ServiceGroupApp.TransactionServiceGroup.TransactionDetailsService

// CreateTransactionDetails 创建交易详情
// @Tags TransactionDetails
// @Summary 创建交易详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body transaction.TransactionDetails true "创建交易详情"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /td/createTransactionDetails [post]
func (tdApi *TransactionDetailsApi) CreateTransactionDetails(c *gin.Context) {
	var td transaction.TransactionDetails
	err := c.ShouldBindJSON(&td)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	authorityId := utils.GetUserAuthorityId(c)
	//
	// 如果 WechatOrderId 为空 并且 authorityId 不等于 888 则返回没权限
	if td.WechatOrderId == "" && authorityId != 888 {
		response.FailWithMessage("没有权限新建交易详情", c)
		return
	}

	// 查询微信订单号 查询交易详情 判断交易详情是否存在
	transactionDetailByWechatOrderId, err := tdService.GetTransactionDetailsByWechatOrderId(td.WechatOrderId)
	if err == nil && transactionDetailByWechatOrderId.ID != 0 {
		response.FailWithMessage("微信订单号已经存在", c)
		return
	}

	if td.WechatOrderId != "" {
		q, err := global.WeChatPay.QueryOrder(td.WechatOrderId)
		if err != nil {
			response.FailWithMessage("查询微信订单失败", c)
			return
		}
		if q.TradeState != "SUCCESS" {
			response.FailWithMessage("微信订单状态不是成功", c)
			return
		}

		if q.SuccessTime == "" {
			response.FailWithMessage("微信订单没有成功时间", c)
			return
		}

		//  判断金额是否一 致
		TransactionAmount := *td.TransactionAmount * 100
		if q.Amount.Total != TransactionAmount {
			response.FailWithMessage("微信订单金额和交易金额不一致", c)
			return
		}
	}

	if err := tdService.CreateTransactionDetails(&td); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTransactionDetails 删除交易详情
// @Tags TransactionDetails
// @Summary 删除交易详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body transaction.TransactionDetails true "删除交易详情"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /td/deleteTransactionDetails [delete]
func (tdApi *TransactionDetailsApi) DeleteTransactionDetails(c *gin.Context) {
	ID := c.Query("ID")
	if err := tdService.DeleteTransactionDetails(ID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTransactionDetailsByIds 批量删除交易详情
// @Tags TransactionDetails
// @Summary 批量删除交易详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /td/deleteTransactionDetailsByIds [delete]
func (tdApi *TransactionDetailsApi) DeleteTransactionDetailsByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := tdService.DeleteTransactionDetailsByIds(IDs); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTransactionDetails 更新交易详情
// @Tags TransactionDetails
// @Summary 更新交易详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body transaction.TransactionDetails true "更新交易详情"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /td/updateTransactionDetails [put]
func (tdApi *TransactionDetailsApi) UpdateTransactionDetails(c *gin.Context) {
	var td transaction.TransactionDetails
	err := c.ShouldBindJSON(&td)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := tdService.UpdateTransactionDetails(td); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTransactionDetails 用id查询交易详情
// @Tags TransactionDetails
// @Summary 用id查询交易详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query transaction.TransactionDetails true "用id查询交易详情"
// @Success 200 {object} response.Response{data=object{retd=transaction.TransactionDetails},msg=string} "查询成功"
// @Router /td/findTransactionDetails [get]
func (tdApi *TransactionDetailsApi) FindTransactionDetails(c *gin.Context) {
	ID := c.Query("ID")
	if retd, err := tdService.GetTransactionDetails(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retd": retd}, c)
	}
}

// GetTransactionDetailsList 分页获取交易详情列表
// @Tags TransactionDetails
// @Summary 分页获取交易详情列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query transactionReq.TransactionDetailsSearch true "分页获取交易详情列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /td/getTransactionDetailsList [get]
func (tdApi *TransactionDetailsApi) GetTransactionDetailsList(c *gin.Context) {
	var pageInfo transactionReq.TransactionDetailsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tdService.GetTransactionDetailsInfoList(pageInfo); err != nil {
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

// GetTransactionDetailsPublic 不需要鉴权的交易详情接口
// @Tags TransactionDetails
// @Summary 不需要鉴权的交易详情接口
// @accept application/json
// @Produce application/json
// @Param data query transactionReq.TransactionDetailsSearch true "分页获取交易详情列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /td/getTransactionDetailsPublic [get]
func (tdApi *TransactionDetailsApi) GetTransactionDetailsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的交易详情接口信息",
	}, "获取成功", c)
}
