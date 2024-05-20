package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/dataConfig"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    example.ServiceGroup
	DataConfigServiceGroup dataConfig.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
