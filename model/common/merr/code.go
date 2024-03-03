package merr

type Msg struct {
	Code int    `json:"code"`
	Msg  string `json:"message"`
}

// 示例：获取异常码：response.YES.Code  获取异常描述：response.YES.SetMsg
var (
	NO  = Msg{100, "请求失败"}
	YES = Msg{200, "请求成功"}

	ConfigErr  = Msg{Code: 401, Msg: "配置错误"}
	ParamErr   = Msg{Code: 402, Msg: "参数错误"}
	ServiceErr = Msg{Code: 403, Msg: "服务错误"}
	UnknownErr = Msg{Code: 404, Msg: "未知错误"}

	ServerErr = Msg{Code: 500, Msg: "未知错误"}
)


