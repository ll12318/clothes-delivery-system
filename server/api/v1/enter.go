package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/bill"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/dataConfig"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/transaction"
)

type ApiGroup struct {
	SystemApiGroup      system.ApiGroup
	ExampleApiGroup     example.ApiGroup
	DataConfigApiGroup  dataConfig.ApiGroup
	BillApiGroup        bill.ApiGroup
	TransactionApiGroup transaction.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
