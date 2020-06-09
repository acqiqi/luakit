package ucenter

import (
	"github.com/astaxie/beego"
	"log"
	"luakit/controllers/v1/apibase"
	"luakit/models"
	"luakit/utils"
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

// 执行秒杀服务插入队列
func (this *SeckillController) PushKillQeue() {
	bindData := struct {
		SeckillId int64 `json:"seckill_id"`
	}{}
	utils.GetPostJson(this.Controller, &bindData)

}

// 秒杀下单
func (this *SeckillController) Kill() {

}
