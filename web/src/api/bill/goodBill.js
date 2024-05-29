import service from "@/utils/request";

// @Tags GoodBill
// @Summary 创建货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodBill true "创建货单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /gb/createGoodBill [post]
export const createGoodBill = (data) => {
  return service({
    url: "/gb/createGoodBill",
    method: "post",
    data,
  });
};

// @Tags GoodBill
// @Summary 删除货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodBill true "删除货单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gb/deleteGoodBill [delete]
export const deleteGoodBill = (params) => {
  return service({
    url: "/gb/deleteGoodBill",
    method: "delete",
    params,
  });
};

// @Tags GoodBill
// @Summary 批量删除货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除货单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /gb/deleteGoodBill [delete]
export const deleteGoodBillByIds = (params) => {
  return service({
    url: "/gb/deleteGoodBillByIds",
    method: "delete",
    params,
  });
};

// @Tags GoodBill
// @Summary 更新货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.GoodBill true "更新货单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /gb/updateGoodBill [put]
export const updateGoodBill = (data) => {
  return service({
    url: "/gb/updateGoodBill",
    method: "put",
    data,
  });
};

// @Tags GoodBill
// @Summary 用id查询货单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.GoodBill true "用id查询货单"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /gb/findGoodBill [get]
export const findGoodBill = (params) => {
  return service({
    url: "/gb/findGoodBill",
    method: "get",
    params,
  });
};

// @Tags GoodBill
// @Summary 分页获取货单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取货单列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /gb/getGoodBillList [get]
export const getGoodBillList = (params) => {
  return service({
    url: "/gb/getGoodBillList",
    method: "get",
    params,
  });
};

// getGoodBillMarketListByDriver
export const getGoodBillMarketListByDriver = (params) => {
  return service({
    url: "/gb/getGoodBillMarketListByDriver",
    method: "get",
    params,
  });
};

// getGoodBillListByMarketId
export const getGoodBillListByMarketId = (params) => {
  return service({
    url: "/gb/getGoodBillListByMarketId",
    method: "get",
    params,
  });
};
