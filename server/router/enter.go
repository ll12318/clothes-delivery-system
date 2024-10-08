package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/bill"
	"github.com/flipped-aurora/gin-vue-admin/server/router/dataConfig"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/transaction"
)

type RouterGroup struct {
	System      system.RouterGroup
	Example     example.RouterGroup
	DataConfig  dataConfig.RouterGroup
	Bill        bill.RouterGroup
	Transaction transaction.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
