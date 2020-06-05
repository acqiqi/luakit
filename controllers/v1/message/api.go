package message

import (
	"github.com/astaxie/beego"
	"log"
	"luakit/controllers/v1/apibase"
	"luakit/models"
	"luakit/utils"
	"strings"
)

type ApiController struct {
	beego.Controller
	Platform models.UcenterPlatform
}

func (this *ApiController) Prepare() {
	log.Println("Prepare")
	platform := apibase.AuthPlatformBase(this.Controller)
	this.Platform = *platform
}

type CB struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Key  string `json:"key"`
	P    interface{}
}

func (this *ApiController) PublishMessageQueue() {
	bindData := struct {
		Cuid       int64             `json:"cuid"`
		MessageKey string            `orm:"message_key" json:"message_key" binding:"required"`
		PushData   map[string]string `json:"push_data"  binding:"required"`
	}{}

	utils.GetPostJson(this.Controller, &bindData)

	mt, err := models.GetMessageTplByMessageKey(bindData.MessageKey)
	if err != nil {
		utils.ApiErr(this.Controller, "消息模板不存在")
	}
	if mt.PlatformKey != this.Platform.PlatformKey {
		utils.ApiErr(this.Controller, "模板有误")
	}

	if mt.IsMsg == 0 {
		utils.ApiOk(this.Controller, "当前消息模板不需要发送消息", nil)
	}
	//
	//实例化model 插入tplid
	queueModel := new(models.MessageQueue)
	queueModel.SmallTplId = mt.SmallTplId
	queueModel.MsgTplId = mt.Id
	queueModel.MessageKey = mt.MessageKey
	queueModel.PlatformKey = this.Platform.PlatformKey
	//插入判定
	queueModel.IsSms = mt.IsSms
	queueModel.IsEmail = mt.IsEmail
	queueModel.IsFormId = mt.IsFormId
	//插入地址
	queueModel.PathType = mt.PathType
	queueModel.PathId = mt.PathId

	queueModel.Title = mt.Title
	queueModel.Desc = mt.Desc
	queueModel.Content = mt.Content
	queueModel.SmallTplContent = mt.SmallTplContent
	queueModel.SmsContent = mt.SmsContent
	queueModel.EmailTitle = mt.EmailTitle
	queueModel.EmailContent = mt.EmailContent
	queueModel.SmallTplPath = mt.SmallTplPath
	queueModel.Flag = 1
	for k, v := range bindData.PushData {
		//fmt.Println(k, ":", v)

		key_name := "{{" + k + "}}"

		queueModel.Title = strings.Replace(queueModel.Title, key_name, v, -1)
		queueModel.Desc = strings.Replace(queueModel.Desc, key_name, v, -1)
		queueModel.Content = strings.Replace(queueModel.Content, key_name, v, -1)
		queueModel.SmallTplContent = strings.Replace(queueModel.SmallTplContent, key_name, v, -1)
		queueModel.SmsContent = strings.Replace(queueModel.SmsContent, key_name, v, -1)
		queueModel.EmailTitle = strings.Replace(queueModel.EmailTitle, key_name, v, -1)
		queueModel.EmailContent = strings.Replace(queueModel.EmailContent, key_name, v, -1)
		queueModel.SmallTplPath = strings.Replace(queueModel.SmallTplPath, key_name, v, -1)
		queueModel.PathId = strings.Replace(queueModel.PathId, key_name, v, -1)
		queueModel.PathType = strings.Replace(queueModel.PathType, key_name, v, -1)
	}
	if mt.IsUcId != 1 {
		// 目前只支持Ucenter 形式请求
		utils.ApiErr(this.Controller, "用户权限方式不正确")
	}
	queueModel.IsUcId = mt.IsUcId //插入是否使用ucenter的id

	if queueModel.IsUcId == 1 {
		//处理ucenter业务
		queueModel.Cuid = bindData.Cuid
		//查询用户是否存在
		user, err := models.GetUcenterUsersById(bindData.Cuid)
		if err != nil {
			utils.ApiErr(this.Controller, "用户不存在")
		}
		queueModel.Cuid = user.Id
	} else {
		//处理三方系统自己的用户业务
		// 暂时不处理
	}
	_, err = models.AddMessageQueue(queueModel)
	if err != nil {
		log.Println(err.Error())
	}
	utils.ApiOk(this.Controller, "提交成功", queueModel)
}

func (this *ApiController) PublishMessage() {
	c := new(CB)
	utils.GetPostJson(this.Controller, &c)

	log.Println("进来了")
	c.Key = this.Platform.PlatformKey
	c.P = this.Platform
	this.Data["json"] = c

	this.ServeJSON()
}
