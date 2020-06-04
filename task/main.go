package task

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/toolbox"
	"log"
	"luakit/models"
	"luakit/utils"
)

func init() {
	log.Println("Init Task")
	messageQueueTk := toolbox.NewTask("messageQueueTk", "0/1 * * * * *", messagePushQueue)
	toolbox.AddTask("messagequeuetk", messageQueueTk)
	//toolbox.StartTask()
}

func onTask() error {
	fmt.Print("heihei")
	return nil
}

func messagePushQueue() error {

	queue, err := models.MessageQueuePop()
	if err != nil {
		log.Println(err.Error())
		return err
	}

	user, err := models.GetUcenterUsersById(queue.Cuid)
	if err != nil {
		log.Println(err)
		return errors.New("用户不存在")
	}

	platform, err := models.GetPlatformByKey(queue.PlatformKey)
	if err != nil {
		log.Println(err)
		return errors.New("平台有误")
	}

	//处理消息
	if queue.IsSms == 1 {
		//发送短信
		go utils.SendSMSLuosinao(user.Mobile, queue.SmsContent)
	}

	if queue.IsFormId == 1 {
		// 处理消息模板
	}

	log.Println("处理回调")
	if platform.MessageCallbackUrl != "" {
		log.Println("开始")
		log.Printf(platform.MessageCallbackUrl)
		cb := utils.ApiDataStruct{}
		utils.HttpPostJson(platform.MessageCallbackUrl, queue, &cb)
	}

	messge := models.Message{}
	messge.Cuid = user.Id
	messge.MessageKey = queue.MessageKey
	messge.Title = queue.Title
	messge.Desc = queue.Desc
	messge.Content = queue.Content
	messge.MessageType = queue.MessageType
	messge.PathType = queue.PathType
	messge.PathId = queue.PathId
	messge.IsFormId = queue.IsFormId
	messge.SmallTplId = queue.SmallTplId
	messge.SmallTplContent = queue.SmallTplContent
	messge.SmallTplOpenid = "1" //openid
	messge.IsSms = queue.IsSms
	messge.Mobile = user.Mobile
	messge.SmsContent = queue.SmsContent
	messge.IsEmail = queue.IsEmail
	messge.Email = queue.Email
	messge.EmailTitle = queue.EmailTitle
	messge.EmailContent = queue.EmailContent
	messge.MsgTplId = queue.MsgTplId
	messge.PlatformKey = platform.PlatformKey
	messge.PushData = queue.PushData
	models.AddMessage(&messge)

	return nil

	//fmt.Print("hello world")
	//callback, _ := common.HttpJsonRequest(common.GetConfig("loop_task_url"), struct {
	//}{})
	//fmt.Print(callback)
}