package dataConfig

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig"
    dataConfigReq "github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig/request"
)

type GoodBillStatusService struct {
}

// CreateGoodBillStatus 创建货单状态记录
// Author [piexlmax](https://github.com/piexlmax)
func (gbsService *GoodBillStatusService) CreateGoodBillStatus(gbs *dataConfig.GoodBillStatus) (err error) {
	err = global.GVA_DB.Create(gbs).Error
	return err
}

// DeleteGoodBillStatus 删除货单状态记录
// Author [piexlmax](https://github.com/piexlmax)
func (gbsService *GoodBillStatusService)DeleteGoodBillStatus(ID string) (err error) {
	err = global.GVA_DB.Delete(&dataConfig.GoodBillStatus{},"id = ?",ID).Error
	return err
}

// DeleteGoodBillStatusByIds 批量删除货单状态记录
// Author [piexlmax](https://github.com/piexlmax)
func (gbsService *GoodBillStatusService)DeleteGoodBillStatusByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]dataConfig.GoodBillStatus{},"id in ?",IDs).Error
	return err
}

// UpdateGoodBillStatus 更新货单状态记录
// Author [piexlmax](https://github.com/piexlmax)
func (gbsService *GoodBillStatusService)UpdateGoodBillStatus(gbs dataConfig.GoodBillStatus) (err error) {
	err = global.GVA_DB.Model(&dataConfig.GoodBillStatus{}).Where("id = ?",gbs.ID).Updates(&gbs).Error
	return err
}

// GetGoodBillStatus 根据ID获取货单状态记录
// Author [piexlmax](https://github.com/piexlmax)
func (gbsService *GoodBillStatusService)GetGoodBillStatus(ID string) (gbs dataConfig.GoodBillStatus, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&gbs).Error
	return
}

// GetGoodBillStatusInfoList 分页获取货单状态记录
// Author [piexlmax](https://github.com/piexlmax)
func (gbsService *GoodBillStatusService)GetGoodBillStatusInfoList(info dataConfigReq.GoodBillStatusSearch) (list []dataConfig.GoodBillStatus, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&dataConfig.GoodBillStatus{})
    var gbss []dataConfig.GoodBillStatus
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&gbss).Error
	return  gbss, total, err
}