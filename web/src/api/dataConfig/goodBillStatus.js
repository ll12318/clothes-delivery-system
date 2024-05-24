import service from '@/utils/request'

// @Tags GoodBillStatus
// @Summary 创建货单状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodBillStatus true "创建货单状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /gbs/createGoodBillStatus [post]
export const createGoodBillStatus = (data) => {
  return service({
    url: '/gbs/createGoodBillStatus',
    method: 'post',
    data
  })
}

// @Tags GoodBillStatus
// @Summary 删除货单状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodBillStatus true "删除货单状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gbs/deleteGoodBillStatus [delete]
export const deleteGoodBillStatus = (params) => {
  return service({
    url: '/gbs/deleteGoodBillStatus',
    method: 'delete',
    params
  })
}

// @Tags GoodBillStatus
// @Summary 批量删除货单状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除货单状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gbs/deleteGoodBillStatus [delete]
export const deleteGoodBillStatusByIds = (params) => {
  return service({
    url: '/gbs/deleteGoodBillStatusByIds',
    method: 'delete',
    params
  })
}

// @Tags GoodBillStatus
// @Summary 更新货单状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodBillStatus true "更新货单状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gbs/updateGoodBillStatus [put]
export const updateGoodBillStatus = (data) => {
  return service({
    url: '/gbs/updateGoodBillStatus',
    method: 'put',
    data
  })
}

// @Tags GoodBillStatus
// @Summary 用id查询货单状态
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GoodBillStatus true "用id查询货单状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gbs/findGoodBillStatus [get]
export const findGoodBillStatus = (params) => {
  return service({
    url: '/gbs/findGoodBillStatus',
    method: 'get',
    params
  })
}

// @Tags GoodBillStatus
// @Summary 分页获取货单状态列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取货单状态列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gbs/getGoodBillStatusList [get]
export const getGoodBillStatusList = (params) => {
  return service({
    url: '/gbs/getGoodBillStatusList',
    method: 'get',
    params
  })
}
