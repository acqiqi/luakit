package ucenter

import (
	"github.com/astaxie/beego"
	"log"
	"luakit/controllers/v1/apibase"
	"luakit/models"
)

type ApiController struct {
	beego.Controller
	Platform models.UcenterPlatform
}

func (this *ApiController) Prepare() {
	log.Println("Prepare")
	platform := apibase.Auth2PlatformBase(this.Controller)
	this.Platform = *platform
}
