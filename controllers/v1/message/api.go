package message

import (
	"github.com/astaxie/beego"
	"log"
	"luakit/models"
	"luakit/utils"
)

type ApiController struct {
	beego.Controller
	Platform models.UcenterPlatform
}

func (this *ApiController) Prepare() {
	log.Println("Prepare")
	platform_key := this.Ctx.Request.Header.Get("PlatformKey")
	//log.Println(platform_key)
	if platform_key == "" {
		this.Data["json"] = utils.ApiOpt(utils.ApiPlatformErr, utils.GetApiMsg(utils.ApiPlatformErr), nil)
		this.ServeJSON()
		this.StopRun()
	}
	platform, err := models.GetPlatformByKey(platform_key)
	if err != nil {
		log.Println(err.Error())
		this.Data["json"] = utils.ApiOpt(utils.ApiPlatformErr, utils.GetApiMsg(utils.ApiPlatformErr), nil)
		this.ServeJSON()
		this.StopRun()
	}
	this.Platform = *platform

	//log.Println("进来了1")
	//c := new(CB)
	//utils.GetPostJson(this.Controller.Ctx.Input.RequestBody,&c)
	//var cc = struct {
	//	Heihei string `json:"he1ihei"`
	//}{Heihei: "11111"}
	//this.Data["json"] = cc
	//this.Platform = "heihei"
	//this.ServeJSON()
}

type CB struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Key  string `json:"key"`
	P    interface{}
}

func (this *ApiController) PublishMessage() {
	c := new(CB)
	err := utils.GetPostJson(this.Controller, &c)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("进来了")
	c.Key = this.Platform.PlatformKey
	c.P = this.Platform
	this.Data["json"] = c

	this.ServeJSON()
}
