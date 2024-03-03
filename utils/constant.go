package utils

const (
	ZERO = 0
	NO   = 1 // 否
	YES  = 2 // 是

	StatusUnable = 1 // StatusUnable 状态禁用

	StatusEnable = 2 // StatusEnable 状态启用
	ConfigEnv    = "GVA_CONFIG"
	ConfigFile   = "config.yaml"
	//ConfigFile = "config-prod.yaml"

	// WorkSchemaScopePub 工作项方案可见范围：公开
	WorkSchemaScopePub = 1

	// WorkSchemaScopeOrg 工作项方案可见范围：部门
	WorkSchemaScopeOrg = 2

	// WorkSchemaScopePri 工作项方案可见范围：私有
	WorkSchemaScopePri = 3

	// 三方服务类型
	THIRD_SERVER_TYPE_ZADIG   = 1
	THIRD_SERVER_TYPE_JENKINS = 2

	WorkitemTypePlat    = 1 // 平台级工作项类型
	WorkitemTypeProject = 2 // 项目级工作项类型

	USER_SYSTEM = "system" // 系统标记，createdBy

	RULE_CODE_FLAG = "rule-" // 规则编码标识

	LeYanPlatOAPre    = "LY"  //乐研平台审批编码前缀
	LeYanPlatOAPreErr = "LeY" //乐研平台审批编码前缀，数据库错误情况下，保证流程通过
)
