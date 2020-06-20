package platform

import (
	"github.com/astaxie/beego"
	uuid "github.com/satori/go.uuid"
	"log"
	"luakit/controllers/v1/apibase"
	"luakit/models"
	"luakit/utils"
	"time"
)

type PublicController struct {
	beego.Controller
	Platform models.UcenterPlatform
}

func (this *PublicController) Prepare() {
	log.Println("Prepare")
	platform := apibase.AuthPlatformBase(this.Controller)
	this.Platform = *platform
}

// 获取平台的Token 用于鉴权
func (this *PublicController) GetToken() {
	bindData := struct {
		PlatformSecret string `json:"platform_secret"`
	}{}
	utils.GetPostJson(this.Controller, &bindData)
	if bindData.PlatformSecret != this.Platform.PlatformSecret {
		utils.ApiErr(this.Controller, "秘钥不正确")
	}

	//创建Token
	token := uuid.NewV4().String()
	log.Println(token)
	utils.SetCacheString("VK_PTOKEN"+token, this.Platform.PlatformKey, time.Second*7200)

	utils.ApiOk(this.Controller, "获取成功", struct {
		Token   string `json:"token"`
		Timeout int    `json:"timeout"`
	}{
		Token:   token,
		Timeout: 7200,
	})
}
