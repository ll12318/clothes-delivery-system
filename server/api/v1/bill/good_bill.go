package bill

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bill"
	billReq "github.com/flipped-aurora/gin-vue-admin/server/model/bill/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	transactionModel "github.com/flipped-aurora/gin-vue-admin/server/model/transaction"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/service/transaction"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GoodBillApi struct {
}

var gbService = service.ServiceGroupApp.BillServiceGroup.GoodBillService

// CreateGoodBill 创建货单
// @Tags GoodBill
// @Summary 创建货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bill.GoodBill true "创建货单"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /gb/createGoodBill [post]
func (gbApi *GoodBillApi) CreateGoodBill(c *gin.Context) {
	var gb bill.GoodBill
	var isPay = false
	err := c.ShouldBindJSON(&gb)
	userUuid := utils.GetUserUuid(c)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	gb.CreatedBy = utils.GetUserID(c)
	// 微信支付 判断
	if gb.WechatOrderId != "" {
		if gb.PayType == "微信支付" {
			tdService := transaction.TransactionDetailsService{}
			// 查询微信订单号 查询交易详情 判断交易详情是否存在
			transactionDetailByWechatOrderId, err := tdService.GetTransactionDetailsByWechatOrderId(gb.WechatOrderId)
			if err == nil && transactionDetailByWechatOrderId.ID != 0 {
				response.FailWithMessage("微信订单号重复", c)
				return
			}

			//if td.WechatOrderId != "" {
			q, err := global.WeChatPay.QueryOrder(gb.WechatOrderId)
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
			StallPrice := gb.Stall.Price * 100
			if q.Amount.Total != StallPrice {
				response.FailWithMessage("微信订单金额和货单金额不一致", c)
				return
			}
			TransactionAmount := float64(0)
			PreTransactionAmount := float64(0)
			PostTransactionAmount := float64(0)
			err = tdService.CreateTransactionDetails(&transactionModel.TransactionDetails{
				TransactionAmount:     &gb.Stall.Price,
				PreTransactionAmount:  &PreTransactionAmount,
				PostTransactionAmount: &PostTransactionAmount,
				UserId:                gb.CreatedBy,
				WechatOrderId:         gb.WechatOrderId,
			})
			if err != nil {
				response.FailWithMessage("创建微信交易详情失败", c)
				return
			}

			TransactionAmount -= gb.Stall.Price
			err = tdService.CreateTransactionDetails(&transactionModel.TransactionDetails{
				TransactionAmount:     &TransactionAmount,
				PreTransactionAmount:  &PreTransactionAmount,
				PostTransactionAmount: &PostTransactionAmount,
				UserId:                gb.CreatedBy,
				WechatOrderId:         gb.WechatOrderId,
			})
			if err != nil {
				response.FailWithMessage("创建微信交易详情失败", c)
				return
			}
			isPay = true
			gb.IsPay = "1"
		}
	}

	// todo 余额支付

	if err := gbService.CreateGoodBill(&gb, userUuid, c, isPay); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {

		//orderRequest := wechat.OrderRequest{
		//	AppID:       global.WeChatPay.AppID,
		//	MchID:       global.WeChatPay.MchID,
		//	Description: "Test order",
		//	OutTradeNo:  gb.BillNumber, // 使用唯一订单号生成函数
		//	NotifyURL:   "https://your_notify_url",
		//	Amount: wechat.Amount{
		//		Total:    1,
		//		Currency: "CNY",
		//	},
		//	Payer: wechat.Payer{
		//		OpenID: "o1u8W7abV4UTsrgfuk6ewZ8gZa2c",
		//	},
		//}

		//rep, err := global.WeChatPay.CreateOrder(orderRequest)
		//if err != nil {
		//	log.Fatalf("Error creating order: %v", err)
		//}
		//发送创建订单请求
		//err = global.WeChatPay.QueryOrder()
		//if err != nil {
		//	log.Fatalf("Error creating order: %v", err)
		//}

		//打印响应
		response.OkWithMessage("创建成功", c)
		//response.OkWithData(gin.H{"resp": "rep"}, c)
	}
}

// DeleteGoodBill 删除货单
// @Tags GoodBill
// @Summary 删除货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bill.GoodBill true "删除货单"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /gb/deleteGoodBill [delete]
func (gbApi *GoodBillApi) DeleteGoodBill(c *gin.Context) {
	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := gbService.DeleteGoodBill(ID, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteGoodBillByIds 批量删除货单
// @Tags GoodBill
// @Summary 批量删除货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /gb/deleteGoodBillByIds [delete]
func (gbApi *GoodBillApi) DeleteGoodBillByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	if err := gbService.DeleteGoodBillByIds(IDs, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateGoodBill 更新货单
// @Tags GoodBill
// @Summary 更新货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bill.GoodBill true "更新货单"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /gb/updateGoodBill [put]
func (gbApi *GoodBillApi) UpdateGoodBill(c *gin.Context) {
	var gb bill.GoodBill
	err := c.ShouldBindJSON(&gb)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	gb.UpdatedBy = utils.GetUserID(c)
	userUuid := utils.GetUserUuid(c)
	if err := gbService.UpdateGoodBill(gb, userUuid, c); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindGoodBill 用id查询货单
// @Tags GoodBill
// @Summary 用id查询货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bill.GoodBill true "用id查询货单"
// @Success 200 {object} response.Response{data=object{regb=bill.GoodBill},msg=string} "查询成功"
// @Router /gb/findGoodBill [get]
func (gbApi *GoodBillApi) FindGoodBill(c *gin.Context) {
	ID := c.Query("ID")
	if regb, err := gbService.GetGoodBill(ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regb": regb}, c)
	}
}

// GetGoodBillList 分页获取货单列表
// @Tags GoodBill
// @Summary 分页获取货单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query billReq.GoodBillSearch true "分页获取货单列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /gb/getGoodBillList [get]
func (gbApi *GoodBillApi) GetGoodBillList(c *gin.Context) {
	var pageInfo billReq.GoodBillSearch
	err := c.ShouldBindQuery(&pageInfo)
	userId := utils.GetUserID(c)
	userUuid := utils.GetUserUuid(c)
	// 获取用户角色组
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := gbService.GetGoodBillInfoList(pageInfo, userUuid, userId); err != nil {
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

// GetGoodBillPublic 不需要鉴权的货单接口
// @Tags GoodBill
// @Summary 不需要鉴权的货单接口
// @accept application/json
// @Produce application/json
// @Param data query billReq.GoodBillSearch true "分页获取货单列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /gb/getGoodBillPublic [get]
func (gbApi *GoodBillApi) GetGoodBillPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的货单接口信息",
	}, "获取成功", c)
}

// getGoodBillMarketListByDriver 司机端获取货单列表

func (gbApi *GoodBillApi) GetGoodBillMarketListByDriver(c *gin.Context) {

	var pageInfo billReq.GoodBillMarketListSearch
	err := c.ShouldBindQuery(&pageInfo)
	userID := utils.GetUserID(c)

	list, total, err := gbService.GetGoodBillMarketListByDriver(userID, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:  list,
		Total: total,
	}, "获取成功", c)
}

// getGoodBillListByMarketId 根据marketId获取货单列表
func (gbApi *GoodBillApi) GetGoodBillListByMarketId(c *gin.Context) {
	var pageInfo billReq.GoodBillMarketListSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userID := utils.GetUserID(c)
	list, total, err := gbService.GetGoodBillListByMarketId(userID, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:  list,
		Total: total,
	}, "获取成功", c)
}
