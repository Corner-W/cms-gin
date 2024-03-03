package config

type Mysql struct {
	Path            string `mapstructure:"path" json:"path" yaml:"path"`                                      // 服务器地址:端口
	Config          string `mapstructure:"config" json:"config" yaml:"config"`                                // 高级配置
	Dbname          string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`                              // 数据库名
	Username        string `mapstructure:"username" json:"username" yaml:"username"`                          // 数据库用户名
	Password        string `mapstructure:"password" json:"password" yaml:"password"`                          // 数据库密码
	MaxIdleConns    int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`          // 空闲中的最大连接数
	MaxOpenConns    int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`          // 打开到数据库的最大连接数
	ConnMaxLifetime int    `mapstructure:"conn-max-lifetime" json:"connMaxLifetime" yaml:"conn-max-lifetime"` // 连接最大生存时间, 单位: 分钟
	LogMode         string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`                           // 是否开启Gorm全局日志
	LogZap          bool   `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`                              // 是否通过zap写入日志文件
	LogLevel        int    `mapstructure:"log-level" json:"logLevel" yaml:"log-level"`                        // sql日志记录等级 Silent:1  Error:2  Warn:3  Info:4
	SlowQuery       int64  `mapstructure:"slow-query" json:"slowQuery" yaml:"slow-query"`                     // sql慢日志阈值
}

func (m *Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
}
