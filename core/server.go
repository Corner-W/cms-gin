package core

import (
	"fmt"

	"cms/global"
	"cms/initialize"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {


	Router := initialize.Routers()

	address := fmt.Sprintf("0.0.0.0:%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	global.GVA_LOG.Info(`
	欢迎使用 ccs
	默认自动化文档地址: http://127.0.0.1:8080/swagger/index.html
`, zap.String("address", address))
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
