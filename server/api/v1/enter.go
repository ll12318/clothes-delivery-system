package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/bill"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/dataConfig"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup     system.ApiGroup
	ExampleApiGroup    example.ApiGroup
	DataConfigApiGroup dataConfig.ApiGroup
	BillApiGroup       bill.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
