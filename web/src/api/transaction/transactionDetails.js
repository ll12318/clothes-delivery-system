import service from '@/utils/request'

// @Tags TransactionDetails
// @Summary 创建交易详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TransactionDetails true "创建交易详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /td/createTransactionDetails [post]
export const createTransactionDetails = (data) => {
  return service({
    url: '/td/createTransactionDetails',
    method: 'post',
    data
  })
}

// @Tags TransactionDetails
// @Summary 删除交易详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TransactionDetails true "删除交易详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /td/deleteTransactionDetails [delete]
export const deleteTransactionDetails = (params) => {
  return service({
    url: '/td/deleteTransactionDetails',
    method: 'delete',
    params
  })
}

// @Tags TransactionDetails
// @Summary 批量删除交易详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除交易详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /td/deleteTransactionDetails [delete]
export const deleteTransactionDetailsByIds = (params) => {
  return service({
    url: '/td/deleteTransactionDetailsByIds',
    method: 'delete',
    params
  })
}

// @Tags TransactionDetails
// @Summary 更新交易详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TransactionDetails true "更新交易详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /td/updateTransactionDetails [put]
export const updateTransactionDetails = (data) => {
  return service({
    url: '/td/updateTransactionDetails',
    method: 'put',
    data
  })
}

// @Tags TransactionDetails
// @Summary 用id查询交易详情
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TransactionDetails true "用id查询交易详情"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /td/findTransactionDetails [get]
export const findTransactionDetails = (params) => {
  return service({
    url: '/td/findTransactionDetails',
    method: 'get',
    params
  })
}

// @Tags TransactionDetails
// @Summary 分页获取交易详情列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取交易详情列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /td/getTransactionDetailsList [get]
export const getTransactionDetailsList = (params) => {
  return service({
    url: '/td/getTransactionDetailsList',
    method: 'get',
    params
  })
}
