package ucenter

import (
	"github.com/astaxie/beego"
	"log"
	"luakit/controllers/v1/apibase"
	"luakit/models"
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
