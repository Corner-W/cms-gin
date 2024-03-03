package global

import (
	"cms/common/enums"

	

	"github.com/songzhibin97/gkit/cache/local_cache"

	"golang.org/x/sync/singleflight"

	"go.uber.org/zap"

	"cms/config"


	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB *gorm.DB

	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_LOG                 *zap.Logger

	GVA_Concurrency_Control             = &singleflight.Group{}

	BlackCache local_cache.Cache

	Env enums.EnvEnum

	EnvPath string
)
