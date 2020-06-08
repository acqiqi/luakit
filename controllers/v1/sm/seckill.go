package ucenter

import (
	"github.com/astaxie/beego"
	"log"
	"luakit/controllers/v1/apibase"
	"luakit/models"
)

type SeckillController struct {
	beego.Controller
	Platform models.UcenterPlatform
}

func (this *SeckillController) Prepare() {
	log.Println("Prepare")
	platform := apibase.Auth2PlatformBase(this.Controller)
	this.Platform = *platform
}

// 秒杀下单
func (this *SeckillController) kill() {

}
