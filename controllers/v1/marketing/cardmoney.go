package marketing

import (
	"github.com/astaxie/beego"
	"log"
	"luakit/common"
	"luakit/controllers/v1/apibase"
	"luakit/models"
	"luakit/utils"
)

func (this *CardMoneyController) Prepare() {
	log.Println("Prepare")
	platform := apibase.AuthPlatformBase(this.Controller)
	this.Platform = *platform
}

type CardMoneyController struct {
	beego.Controller
	Platform models.UcenterPlatform
}

func (this *CardMoneyController) ExcelOut() {
	c := struct {
		Id int64 `json:"id"`
	}{}
	utils.GetPostJson(this.Controller, &c)

	//查询所有卡片
	//card_list,err;

	u := common.ExcelUtils{}
	err, quotation_no := u.Test(c.Id)
	if err != nil {
		log.Println(err.Error())
		utils.ApiErr(this.Controller, err.Error())
	}

	utils.ApiOk(this.Controller, "获取成功", struct {
		Path string `json:"path"`
	}{
		Path: "http://export.ddgongjiang.com/" + quotation_no + ".xlsx",
	})
}
