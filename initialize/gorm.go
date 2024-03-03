package initialize

import (
	"log"
	"time"

	"cms/global"
	"cms/initialize/internal"
	mySqlLog "cms/utils/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

//
//@function: Gorm
//@description: 初始化数据库并产生数据库全局变量
//@return: *gorm.DB

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}

//
//@function: GormMysql
//@description: 初始化Mysql数据库
//@return: *gorm.DB

func GormMysql() *gorm.DB {
	m := global.GVA_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), gormConfig()); err != nil {
		//global.GVA_LOG.Error("MySQL启动异常", zap.Any("err", err))
		//os.Exit(0)
		//return nil
		return nil
	} else {
		sqlDB, _ := db.DB()
		// 以下代码指挥在执行create操作时生效，在.Transaction开启事务时不生效
		//var openedNum int64
		//if global.Env == enums.LOCAL || global.Env == enums.DEV {
		//	db.Callback().Create().Before("gorm:begin_transaction").Register("leyan-backend:begin_transaction", func(db *gorm.DB) {
		//		atomic.AddInt64(&openedNum, 1)
		//		context := db.Statement.Context
		//		global.GVA_LOG.Info("leyan-backend:begin_transaction", zap.String("info", fmt.Sprintf("Context: %s OpendNum: %v", context, openedNum)))
		//	})
		//
		//	db.Callback().Create().After("gorm:commit_or_rollback_transaction").Register("leyan-backend:commit_or_rollback_transaction", func(db *gorm.DB) {
		//		atomic.AddInt64(&openedNum, -1)
		//		context := db.Statement.Context
		//		sql := db.Statement.SQL.String()
		//		global.GVA_LOG.Info("leyan-backend:stop_transaction", zap.String("info", fmt.Sprintf("Context: %s SQL: %s OpendNum: %v", context, sql, openedNum)))
		//	})
		//}

		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime) * time.Minute)
		return db
	}
}

//
//@function: gormConfig
//@description: 根据配置决定是否开启日志
//@param: mod bool
//@return: *gorm.Config

func gormConfig() *gorm.Config {
	config := &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
	switch global.GVA_CONFIG.Mysql.LogMode {
	case "silent", "Silent":
		config.Logger = internal.Default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = internal.Default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = internal.Default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = internal.Default.LogMode(logger.Info)
	default:
		config.Logger = internal.Default.LogMode(logger.Info)
	}
	slowQuery := global.GVA_CONFIG.Mysql.SlowQuery
	logLevel := global.GVA_CONFIG.Mysql.LogLevel
	if slowQuery > 0 || logLevel > 0 { //有配置且不为默认值0时才开启sql日志文件记录
		if slowQuery == 0 { //默认记录慢日志1000毫秒
			slowQuery = 1000
		}
		if logLevel == 0 || (logLevel != 1 && logLevel != 2 && logLevel != 3 && logLevel != 4) { //默认记录sql警告日志
			logLevel = 3 //Silent:1  Error:2  Warn:3  Info:4
		}
		newLogger := logger.New(
			log.New(mySqlLog.Logger, "", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Duration(slowQuery) * time.Millisecond,
				LogLevel:      logger.LogLevel(logLevel),
				Colorful:      false,
			},
		)
		config.Logger = newLogger
	}

	return config
}
