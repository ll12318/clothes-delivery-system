// 自动生成模板GoodBill
package bill

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// 货单 结构体  GoodBill
type GoodBill struct {
 global.GVA_MODEL 
      Remarks  string `json:"remarks" form:"remarks" gorm:"column:remarks;comment:;"`  //备注 
      CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
      UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
      DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 货单 GoodBill自定义表名 good_bill
func (GoodBill) TableName() string {
  return "good_bill"
}

