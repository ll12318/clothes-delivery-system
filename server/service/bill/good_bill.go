package bill

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bill"
	billReq "github.com/flipped-aurora/gin-vue-admin/server/model/bill/request"
	dataConfigModel "github.com/flipped-aurora/gin-vue-admin/server/model/dataConfig"
	systemRes "github.com/flipped-aurora/gin-vue-admin/server/model/system/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service/dataConfig"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
	"strconv"
	time2 "time"
)

type GoodBillService struct {
}

// CreateGoodBill 创建货单记录
// Author [piexlmax](https://github.com/piexlmax)
func (gbService *GoodBillService) CreateGoodBill(gb *bill.GoodBill) (err error) {
	routeService := dataConfig.ServiceGroup{}
	routeId := gb.Stall.RouteId
	if routeId != 0 {
		route, err := routeService.GetRoute(strconv.Itoa(int(routeId)))
		if err != nil {
			return err
		}
		if route.User.ID != 0 {
			gb.TakeGoodPeopleId = route.User.ID
		}
	}
	marketId := gb.Stall.MarketId
	if marketId != 0 {
		gb.MarketId = marketId
		gb.Market = gb.Stall.Market
	}

	err = global.GVA_DB.Create(gb).Error
	return err
}

// DeleteGoodBill 删除货单记录
// Author [piexlmax](https://github.com/piexlmax)
func (gbService *GoodBillService) DeleteGoodBill(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&bill.GoodBill{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&bill.GoodBill{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteGoodBillByIds 批量删除货单记录
// Author [piexlmax](https://github.com/piexlmax)
func (gbService *GoodBillService) DeleteGoodBillByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&bill.GoodBill{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&bill.GoodBill{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateGoodBill 更新货单记录
// Author [piexlmax](https://github.com/piexlmax)
func (gbService *GoodBillService) UpdateGoodBill(gb bill.GoodBill, userUuid uuid.UUID) (err error) {
	b := *gb.FinishStatus == true
	if b {
		if gb.FinishTime == "" {
			// 如果完成状态为true，设置完成时间为当前时间
			time := time2.Now()
			gb.FinishTime = time.Format("2006-01-02 15:04:05")
		}
		if gb.FinishPeopleId == 0 {
			// 如果完成状态为true，设置完成人为当前调用者
			userService := system.UserService{}
			userInfo, err := userService.GetUserInfo(userUuid)
			// 如果获取用户信息失败
			if err != nil {
				return err
			}
			gb.FinishPeopleId = userInfo.ID
		}
	}
	err = global.GVA_DB.Model(&bill.GoodBill{}).Where("id = ?", gb.ID).Updates(&gb).Error
	return err
}

// GetGoodBill 根据ID获取货单记录
// Author [piexlmax](https://github.com/piexlmax)
func (gbService *GoodBillService) GetGoodBill(ID string) (gb bill.GoodBill, err error) {
	err = global.GVA_DB.Preload("Stall.Market").Where("id = ?", ID).First(&gb).Error
	return
}

// GetGoodBillInfoList 分页获取货单记录
// Author [piexlmax](https://github.com/piexlmax)
func (gbService *GoodBillService) GetGoodBillInfoList(info billReq.GoodBillSearch, userUuid uuid.UUID, userId uint) (list []bill.GoodBill, total int64, err error) {

	db := global.GVA_DB.Model(&bill.GoodBill{}).Preload("CreatedByUserInfo", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "nick_name") // 选择需要的字段
	}).Preload("Stall.Market").Preload("TakeGoodPeople").Preload("GoodBillStatus").Preload("FinishPeople")
	// 查询用户的角色

	userService := system.UserService{}
	userInfo, err := userService.GetUserInfo(userUuid)
	if err != nil {
		return
	}
	authorities := userInfo.Authorities

	isAdmin := false
	isUser := false
	isDriver := false
	// 搜索下单人
	if info.CreatedBy != 0 {
		db = db.Where("created_by = ?", info.CreatedBy)
	}
	// 搜索拿货人
	if info.TakeGoodPeopleId != 0 {
		db = db.Where("take_good_people_id = ?", info.TakeGoodPeopleId)
	}

	if info.FinishStatus != nil {
		if *info.FinishStatus == false {
			db = db.Where("finish_status IS NULL OR finish_status = ?", *info.FinishStatus)
		}
		if *info.FinishStatus == true {
			db = db.Where("finish_status = ?", *info.FinishStatus)
		}
	}

	if info.DriverVerify != nil {
		if *info.DriverVerify == false {
			db = db.Where("driver_verify IS NULL OR driver_verify = ?", *info.DriverVerify)
		}
		if *info.DriverVerify == true {
			db = db.Where("driver_verify = ?", *info.DriverVerify)
		}
	}

	// 判断用户角色是不是超级管理员 管理员 888 用户998 司机999
	for _, authority := range authorities {
		authorityId := authority.AuthorityId
		// 判断用户角色是不是超级管理员 管理员 888 用户998 司机999
		if authorityId == 888 {
			isAdmin = true
		} else if authorityId == 998 {
			isUser = true
		} else if authorityId == 999 {
			isDriver = true
		}
	}
	if isAdmin {
		// 如果是超级管理员，查询所有数据
		list, total, err = gbService.GetGoodBillInfoListByAdmin(db, info)
		if err != nil {
			return
		}
		return list, total, err

	} else if isUser {
		// 如果是普通用户，查询自己创建的数据
		list, total, err = gbService.GetGoodBillInfoListByUser(db, info, userId)
		if err != nil {
			return
		}
		return list, total, err

	} else if isDriver {
		// 如果是司机
	}
	return list, total, err

}

func (gbService *GoodBillService) GetGoodBillInfoListByAdmin(db *gorm.DB, info billReq.GoodBillSearch) (list []bill.GoodBill, total int64, err error) {
	limit := info.PageSize
	var gbs []bill.GoodBill
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&gbs).Error
	gbService.handlerCreatedBySimpleUser(gbs)
	return gbs, total, err
}

func (gbService *GoodBillService) handlerCreatedBySimpleUser(gbs []bill.GoodBill) {
	for i := 0; i < len(gbs); i++ {
		// 查询创建人信息
		gbs[i].CreatedBySimpleUser = systemRes.CreatedBySimpleUser{
			ID:       gbs[i].CreatedByUserInfo.ID,
			NickName: gbs[i].CreatedByUserInfo.NickName,
		}
	}
}

// GetGoodBillInfoListByUser
func (gbService *GoodBillService) GetGoodBillInfoListByUser(db *gorm.DB, info billReq.GoodBillSearch, userId uint) (list []bill.GoodBill, total int64, err error) {
	limit := info.PageSize
	var gbs []bill.GoodBill
	offset := info.PageSize * (info.Page - 1)
	// 只查询自己创建的数据
	db = db.Where("created_by = ?", userId)
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&gbs).Error
	gbService.handlerCreatedBySimpleUser(gbs)
	return gbs, total, err

}

// getGoodBillMarketListByDriver
func (gbService *GoodBillService) GetGoodBillMarketListByDriver(userId uint, info billReq.GoodBillMarketListSearch) (markets []dataConfigModel.Market, total int64, err error) {
	var goodBills []bill.GoodBill

	finishStatus := info.FinishStatus
	driverVerify := info.DriverVerify
	db := global.GVA_DB.Model(&bill.GoodBill{})
	if *finishStatus == true {
		db.
			Select("market_id").
			Where("take_good_people_id = ? AND finish_status = ?", userId, true).
			Preload("Market")
	}

	if *finishStatus == false {
		db.
			Select("market_id").
			Where("take_good_people_id = ?", userId).
			Where("finish_status IS NULL OR finish_status = ?", false).
			Preload("Market")
	}

	if driverVerify != nil {
		if *driverVerify == true {
			db.
				Select("market_id").
				Where("take_good_people_id = ? AND driver_verify = ?", userId, true).
				Preload("Market")
		}
		if *driverVerify == false {
			db.
				Select("market_id").
				Where("take_good_people_id = ? AND driver_verify IS NULL OR driver_verify = ?", userId, false).
				Preload("Market")
		}
	}
	// 查询TakeGoodPeopleId为userId的数据并且查询出MarketId的数据
	err = db.
		Find(&goodBills).
		Error

	if err != nil {
		return nil, 0, err
	}

	// 获取 Market集合并去重
	marketMap := make(map[uint]dataConfigModel.Market)
	for _, goodBill := range goodBills {
		marketMap[goodBill.MarketId] = goodBill.Market
	}

	// 转换 map 为 slice
	for _, market := range marketMap {
		markets = append(markets, market)
	}

	total = int64(len(markets))
	return markets, total, nil
}

// GetGoodBillListByMarketId 根据marketId获取货单列表
func (gbService *GoodBillService) GetGoodBillListByMarketId(userId uint, info billReq.GoodBillMarketListSearch) (goodBills []bill.GoodBill, total int64, err error) {
	marketId := info.MarketId
	finishStatus := info.FinishStatus
	driverVerify := info.DriverVerify
	db := global.GVA_DB.Model(&bill.GoodBill{}).
		Preload("CreatedByUserInfo", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "nick_name") // 选择需要的字段
		}).
		Preload("Stall.Market").
		Preload("TakeGoodPeople").
		Preload("GoodBillStatus").
		Preload("FinishPeople").Preload("Market")
	if marketId != 0 {
		db = db.Where("market_id = ?", marketId)
	}
	db.Where("take_good_people_id = ?", userId)
	if *finishStatus == true {
		err = db.Where("finish_status = ?", true).
			Find(&goodBills).Error
	}
	if *finishStatus == false {
		err = db.Where("finish_status IS NULL OR finish_status = ?", false).
			Find(&goodBills).Error
	}
	if driverVerify != nil {
		if *driverVerify == true {
			err = db.Where("driver_verify = ?", true).
				Find(&goodBills).Error
		}
		if *driverVerify == false {
			err = db.Where("driver_verify IS NULL OR driver_verify = ?", false).
				Find(&goodBills).Error
		}
	}
	db.Count(&total)

	return goodBills, total, err
}
