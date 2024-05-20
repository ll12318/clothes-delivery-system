import service from '@/utils/request'

// @Tags Stall
// @Summary 创建档口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stall true "创建档口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /stall/createStall [post]
export const createStall = (data) => {
  return service({
    url: '/stall/createStall',
    method: 'post',
    data
  })
}

// @Tags Stall
// @Summary 删除档口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stall true "删除档口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stall/deleteStall [delete]
export const deleteStall = (params) => {
  return service({
    url: '/stall/deleteStall',
    method: 'delete',
    params
  })
}

// @Tags Stall
// @Summary 批量删除档口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除档口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /stall/deleteStall [delete]
export const deleteStallByIds = (params) => {
  return service({
    url: '/stall/deleteStallByIds',
    method: 'delete',
    params
  })
}

// @Tags Stall
// @Summary 更新档口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Stall true "更新档口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /stall/updateStall [put]
export const updateStall = (data) => {
  return service({
    url: '/stall/updateStall',
    method: 'put',
    data
  })
}

// @Tags Stall
// @Summary 用id查询档口
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Stall true "用id查询档口"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /stall/findStall [get]
export const findStall = (params) => {
  return service({
    url: '/stall/findStall',
    method: 'get',
    params
  })
}

// @Tags Stall
// @Summary 分页获取档口列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取档口列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /stall/getStallList [get]
export const getStallList = (params) => {
  return service({
    url: '/stall/getStallList',
    method: 'get',
    params
  })
}
