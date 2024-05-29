import service from '@/utils/request'

// @Tags Market
// @Summary 创建市场
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Market true "创建市场"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /market/createMarket [post]
export const createMarket = (data) => {
  return service({
    url: '/market/createMarket',
    method: 'post',
    data
  })
}

// @Tags Market
// @Summary 删除市场
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Market true "删除市场"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /market/deleteMarket [delete]
export const deleteMarket = (params) => {
  return service({
    url: '/market/deleteMarket',
    method: 'delete',
    params
  })
}

// @Tags Market
// @Summary 批量删除市场
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除市场"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /market/deleteMarket [delete]
export const deleteMarketByIds = (params) => {
  return service({
    url: '/market/deleteMarketByIds',
    method: 'delete',
    params
  })
}

// @Tags Market
// @Summary 更新市场
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Market true "更新市场"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /market/updateMarket [put]
export const updateMarket = (data) => {
  return service({
    url: '/market/updateMarket',
    method: 'put',
    data
  })
}

// @Tags Market
// @Summary 用id查询市场
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Market true "用id查询市场"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /market/findMarket [get]
export const findMarket = (params) => {
  return service({
    url: '/market/findMarket',
    method: 'get',
    params
  })
}

// @Tags Market
// @Summary 分页获取市场列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取市场列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /market/getMarketList [get]
export const getMarketList = (params) => {
  return service({
    url: '/market/getMarketList',
    method: 'get',
    params
  })
}
