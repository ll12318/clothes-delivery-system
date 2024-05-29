package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/bill"
	"github.com/flipped-aurora/gin-vue-admin/server/service/dataConfig"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    example.ServiceGroup
	DataConfigServiceGroup dataConfig.ServiceGroup
	BillServiceGroup       bill.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
