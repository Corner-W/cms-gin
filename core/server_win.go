//go:build windows
// +build windows

package core

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func initServer(address string, router *gin.Engine) server {
	s := &http.Server{
		Addr:              address,
		Handler:           router,
		ReadTimeout:       60 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       5 * 60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	//global.GVA_LOG.Info("http.Server:", zap.String("", fmt.Sprintf("s:%v", s)))

	return s
}
