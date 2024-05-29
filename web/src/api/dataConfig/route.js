import service from '@/utils/request'

// @Tags Route
// @Summary 创建路线
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Route true "创建路线"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /route/createRoute [post]
export const createRoute = (data) => {
  return service({
    url: '/route/createRoute',
    method: 'post',
    data
  })
}

// @Tags Route
// @Summary 删除路线
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Route true "删除路线"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /route/deleteRoute [delete]
export const deleteRoute = (params) => {
  return service({
    url: '/route/deleteRoute',
    method: 'delete',
    params
  })
}

// @Tags Route
// @Summary 批量删除路线
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除路线"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /route/deleteRoute [delete]
export const deleteRouteByIds = (params) => {
  return service({
    url: '/route/deleteRouteByIds',
    method: 'delete',
    params
  })
}

// @Tags Route
// @Summary 更新路线
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Route true "更新路线"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /route/updateRoute [put]
export const updateRoute = (data) => {
  return service({
    url: '/route/updateRoute',
    method: 'put',
    data
  })
}

// @Tags Route
// @Summary 用id查询路线
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Route true "用id查询路线"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /route/findRoute [get]
export const findRoute = (params) => {
  return service({
    url: '/route/findRoute',
    method: 'get',
    params
  })
}

// @Tags Route
// @Summary 分页获取路线列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取路线列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /route/getRouteList [get]
export const getRouteList = (params) => {
  return service({
    url: '/route/getRouteList',
    method: 'get',
    params
  })
}
