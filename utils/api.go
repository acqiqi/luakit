package utils

import (
	"github.com/astaxie/beego"
	"log"
)

// Api 数据模型
type ApiDataStruct struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

const (
	ApiSuccess = 0
	ApiError   = 1
	Api404     = 404
	Api501     = 501
	Api502     = 502

	ApiNotAuth          = 45000 //未登录
	ApiNotBindMobile    = 41001
	ApiNotBindUserInfo  = 41002 //三方用户未注册
	ApiNotReg           = 41004 //ucenter未注册
	ApiThirdNotReg      = 41005 //三方信息不存在
	ApiPlatformErrAppid = 42001 //平台appid不正确
	ApiPlatformErr      = 42002 //平台错误
	ApiJsonDecodeErr    = 49001 //Json解析错误
)

var statusText = map[int]string{
	ApiSuccess:          "操作成功",
	ApiError:            "操作失败",
	Api404:              "页面丢了",
	Api501:              "页面丢了",
	Api502:              "页面丢了",
	ApiNotAuth:          "未登录或登录超时",
	ApiNotBindMobile:    "没有绑定手机号",
	ApiNotBindUserInfo:  "没有绑定用户信息",
	ApiNotReg:           "UserCenter未注册",
	ApiThirdNotReg:      "三方信息不存在",
	ApiPlatformErrAppid: "平台appid不正确",
	ApiPlatformErr:      "平台有误",
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

// 返回数据
type CallBackStrcut struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// 带分页数据结构体
type PageDataStruct struct {
	Page  int64       `json:"page"`
	Limit int64       `json:"limit"`
	Lists interface{} `json:"lists"`
}

// Api 成功
func ApiJsonOk(msg string, data interface{}) (cb ApiDataStruct) {
	return ApiDataStruct{
		Code: 0,
		Msg:  msg,
		Data: data,
	}
}

// Api 失败
func ApiJsonErr(msg string) (cb ApiDataStruct) {
	return ApiDataStruct{
		Code: 1,
		Msg:  msg,
		Data: nil,
	}
}

func ApiJsonOpt(code int, msg string, data interface{}) (cb ApiDataStruct) {
	return ApiDataStruct{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

// 控制器返回成功
func ApiOk(c beego.Controller, msg string, data interface{}) {
	db := ApiDataStruct{
		Code: 0,
		Msg:  msg,
		Data: data,
	}
	c.Data["json"] = db
	c.ServeJSON()
	c.StopRun()
}

// 控制器返回失败
func ApiErr(c beego.Controller, msg string) {
	db := ApiDataStruct{
		Code: 1,
		Msg:  msg,
		Data: nil,
	}
	c.Data["json"] = db
	c.ServeJSON()
	c.StopRun()
}

// 控制器返回其他
func ApiOpt(c beego.Controller, code int, msg string, data interface{}) {
	db := ApiDataStruct{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	c.Data["json"] = db
	c.ServeJSON()
	c.StopRun()
}

// 获取post 过来的json 实例
func GetPostJson(c beego.Controller, cb interface{}) {
	data := c.Ctx.Input.RequestBody
	if err := JsonDecode(string(data), &cb); err != nil {
		log.Println("参数解析有误", err.Error())
		ApiErr(c, "参数解析有误")
		return
	}
	return
}
