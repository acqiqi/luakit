package partner

import (
	"github.com/astaxie/beego"
	"log"
	"luakit/controllers/v1/apibase"
	"luakit/models"
	"luakit/utils"
)

func (this *CouponController) Prepare() {
	log.Println("Prepare")
	platform := apibase.Auth2PlatformBase(this.Controller)
	this.Platform = *platform
}

type CouponController struct {
	beego.Controller
	Platform models.UcenterPlatform
}

func (this *CouponController) GetInfo() {
	c := struct {
		CouponId int64 `json:"coupon_id"`
	}{}
	utils.GetPostJson(this.Controller, &c)
	conpon, err := models.GetMarketingCouponById(c.CouponId)
	if err != nil {
		utils.ApiErr(this.Controller, "优惠券不存在")
	}

	utils.ApiOk(this.Controller, "获取成功", conpon)
}
