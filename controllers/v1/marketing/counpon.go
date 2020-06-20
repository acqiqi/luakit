package marketing

import (
	"github.com/astaxie/beego"
	"log"
	"luakit/controllers/v1/apibase"
	"luakit/models"
	"luakit/utils"
	"strconv"
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

func (this *CouponController) GetUserList() {
	c := struct {
		Cuid  int64       `json:"cuid"`
		Page  int64       `json:"page"`
		Limit int64       `json:"limit"`
		Lists interface{} `json:"lists"`
	}{}
	utils.GetPostJson(this.Controller, &c)

	_, err := models.GetUcenterUsersById(c.Cuid)
	if err != nil {
		utils.ApiErr(this.Controller, "用户不存在")
	}

	maps := make(map[string]string)
	maps["cuid"] = strconv.FormatInt(c.Cuid, 10)
	maps["flag"] = "1"
	list, err := models.GetAllMarketingConpon(maps, []string{}, []string{"id"}, []string{"desc"}, c.Limit*(c.Page-1), c.Limit)

	utils.ApiOk(this.Controller, "获取成功", utils.PageDataStruct{
		Page:  c.Page,
		Limit: c.Limit,
		Lists: list,
	})

}
