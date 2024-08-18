package transaction

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/transaction"
	transactionReq "github.com/flipped-aurora/gin-vue-admin/server/model/transaction/request"
	"strconv"
)

type TransactionDetailsService struct {
}

// CreateTransactionDetails 创建交易详情记录
// Author [piexlmax](https://github.com/piexlmax)
func (tdService *TransactionDetailsService) CreateTransactionDetails(td *transaction.TransactionDetails) (err error) {
	transactionAmount := *td.TransactionAmount
	userId := td.UserId
	// 获取用户余额
	ltd, err := tdService.GetTransactionDetailsByUserId(strconv.Itoa(int(userId)))
	if err != nil {
		*td.PostTransactionAmount = transactionAmount
		*td.PreTransactionAmount = 0.00
	} else {
		// 如果用户金额不为空 交易后金额 = 上次金额 + 交易金额
		newAmount := *ltd.PostTransactionAmount + transactionAmount
		td.PostTransactionAmount = &newAmount
		td.PreTransactionAmount = ltd.PostTransactionAmount
	}
	err = global.GVA_DB.Create(td).Error
	if err != nil {
		return err
	}
	// 给用户余额
	err = global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", userId).Update("balance", td.PostTransactionAmount).Error
	return err
}

// DeleteTransactionDetails 删除交易详情记录
// Author [piexlmax](https://github.com/piexlmax)
func (tdService *TransactionDetailsService) DeleteTransactionDetails(ID string) (err error) {
	err = global.GVA_DB.Delete(&transaction.TransactionDetails{}, "id = ?", ID).Error
	return err
}

// DeleteTransactionDetailsByIds 批量删除交易详情记录
// Author [piexlmax](https://github.com/piexlmax)
func (tdService *TransactionDetailsService) DeleteTransactionDetailsByIds(IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]transaction.TransactionDetails{}, "id in ?", IDs).Error
	return err
}

// UpdateTransactionDetails 更新交易详情记录
// Author [piexlmax](https://github.com/piexlmax)
func (tdService *TransactionDetailsService) UpdateTransactionDetails(td transaction.TransactionDetails) (err error) {
	err = global.GVA_DB.Model(&transaction.TransactionDetails{}).Where("id = ?", td.ID).Updates(&td).Error
	return err
}

// GetTransactionDetails 根据ID获取交易详情记录
// Author [piexlmax](https://github.com/piexlmax)
func (tdService *TransactionDetailsService) GetTransactionDetails(ID string) (td transaction.TransactionDetails, err error) {
	err = global.GVA_DB.Preload("User").Where("id = ?", ID).First(&td).Error
	return
}

func (tdService *TransactionDetailsService) GetTransactionDetailsByWechatOrderId(wechatOrderId string) (td transaction.TransactionDetails, err error) {
	err = global.GVA_DB.Preload("User").Where("wechat_order_id = ?", wechatOrderId).First(&td).Error
	return
}

// GetTransactionDetails 根据userId获取交易详情记录最后一条记录
func (tdService *TransactionDetailsService) GetTransactionDetailsByUserId(userId string) (td transaction.TransactionDetails, err error) {
	err = global.GVA_DB.Preload("User").Where("user_id = ?", userId).Last(&td).Error
	return
}

// GetTransactionDetailsInfoList 分页获取交易详情记录
// Author [piexlmax](https://github.com/piexlmax)
func (tdService *TransactionDetailsService) GetTransactionDetailsInfoList(info transactionReq.TransactionDetailsSearch) (list []transaction.TransactionDetails, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&transaction.TransactionDetails{}).Preload("User")
	var tds []transaction.TransactionDetails
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

	err = db.Find(&tds).Error
	return tds, total, err
}
