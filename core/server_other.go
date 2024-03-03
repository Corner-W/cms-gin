//go:build !windows
// +build !windows

package core

import (
	"syscall"
	"time"

	"cms/global"
	"cms/utils/collection"
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20

	var expire time.Duration
	if collection.SliceContain([]string{"prod", "test", "dev"}, global.GVA_CONFIG.System.Env) {
		expire = 5 * time.Second
	} else {
		expire = 1 * time.Second
	}
	f := func() {
		global.GVA_LOG.Info("Shutting down micro server...")
		time.Sleep(expire)
	}
	_ = s.RegisterSignalHook(endless.PRE_SIGNAL, syscall.SIGINT, f)
	_ = s.RegisterSignalHook(endless.PRE_SIGNAL, syscall.SIGTERM, f)
	return s
}
