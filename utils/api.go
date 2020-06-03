package utils

// Api 数据模型
type ApiDataStruct struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	ApiSuccess          = 0
	ApiError            = 1
	Api404              = 404
	ApiNotAuth          = 45000 //未登录
	ApiNotBindMobile    = 41001
	ApiNotBindUserInfo  = 41002 //三方用户未注册
	ApiNotReg           = 41004 //ucenter未注册
	ApiThirdNotReg      = 41005 //三方信息不存在
	ApiPlatformErrAppid = 42001 //平台appid不正确
	ApiJsonDecodeErr    = 49001 //Json解析错误
)

var statusText = map[int]string{
	ApiSuccess:          "操作成功",
	ApiError:            "操作失败",
	Api404:              "页面丢了",
	ApiNotAuth:          "未登录或登录超时",
	ApiNotBindMobile:    "没有绑定手机号",
	ApiNotBindUserInfo:  "没有绑定用户信息",
	ApiNotReg:           "UserCenter未注册",
	ApiThirdNotReg:      "三方信息不存在",
	ApiPlatformErrAppid: "平台appid不正确",
	ApiJsonDecodeErr:    "Json解析错误",
}

//获取code对应错误字段
func GetApiMsg(status int) (msg string) {
	return statusText[status]
}

// 提供空 obj类型
func GetEmptyStruct() interface{} {
	return struct {
	}{}
}

// Api 成功
func ApiOk(msg string, data interface{}) (cb ApiDataStruct) {
	return ApiDataStruct{
		Code: 0,
		Msg:  msg,
		Data: data,
	}
}

// Api 失败
func ApiErr(msg string) (cb ApiDataStruct) {
	return ApiDataStruct{
		Code: 1,
		Msg:  msg,
		Data: nil,
	}
}

func ApiOpt(code int, msg string, data interface{}) (cb ApiDataStruct) {
	return ApiDataStruct{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
