package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/dataConfig"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	DataConfig dataConfig.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
