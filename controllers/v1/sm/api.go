package sm

import (
	"github.com/astaxie/beego"
	"log"
	"luakit/common"
	"luakit/controllers/v1/apibase"
	"luakit/models"
	"luakit/utils"
)

type ApiController struct {
	beego.Controller
	Platform models.UcenterPlatform
}

func (this *ApiController) Prepare() {
	log.Println("Prepare")
	platform := apibase.AuthPlatformBase(this.Controller)
	this.Platform = *platform
}

func (this *ApiController) AccountOrder() {
	c := struct {
		UsersServiceId int64 `json:"users_service_id"`
	}{}
	utils.GetPostJson(this.Controller, &c)

	log.Println(c)
	accounts_utils := new(common.Accounts)
	accounts_utils.SmUsersServiceId = c.UsersServiceId
	if err := accounts_utils.InitSmUsersService(); err != nil {
		utils.ApiErr(this.Controller, err.Error())
	}
	if err := accounts_utils.AccountsSmServiceOrders(); err != nil {
		utils.ApiErr(this.Controller, err.Error())
	}

	utils.ApiOk(this.Controller, "操作成功", utils.GetEmptyStruct())
}
