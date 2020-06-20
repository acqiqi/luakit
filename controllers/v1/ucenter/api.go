package ucenter

import (
	"github.com/astaxie/beego"
	"log"
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
	platform := apibase.Auth2PlatformBase(this.Controller)
	this.Platform = *platform
}

// 根据用户手机号获取用户信息
func (this *ApiController) QueryUserInfoOrMobile() {
	bindData := struct {
		Mobile string `json:"mobile"`
	}{}
	utils.GetPostJson(this.Controller, &bindData)

	m, err := models.GetUcenterUsersByMobile(bindData.Mobile)
	if err != nil {
		log.Println(err.Error())
		utils.ApiErr(this.Controller, "用户不存在")
	}

	utils.ApiOk(this.Controller, "获取成功", m)
}

// 根据用户id获取用户信息
func (this *ApiController) QueryUserInfoOrCuid() {
	bindData := struct {
		Cuid int64 `json:"cuid"`
	}{}
	utils.GetPostJson(this.Controller, &bindData)

	m, err := models.GetUcenterUsersById(bindData.Cuid)
	if err != nil {
		log.Println(err.Error())
		utils.ApiErr(this.Controller, "用户不存在")
	}

	utils.ApiOk(this.Controller, "获取成功", m)
}

func (this *ApiController) QueryUsers() {
	bindData := utils.PageDataStruct{}
	utils.GetPostJson(this.Controller, &bindData)

	if bindData.Page == 0 {
		bindData.Page = 1
	}
	if bindData.Limit == 0 {
		bindData.Limit = 20
	}

	maps := make(map[string]string)
	//maps["status__gt"] = "0"

	list, err := models.GetAllUcenterUsers(maps, []string{}, []string{"id"}, []string{"desc"}, bindData.Limit*(bindData.Page-1), bindData.Limit)
	if err != nil {
		log.Println(err.Error())
		utils.ApiErr(this.Controller, "获取失败")
	}
	utils.ApiOk(this.Controller, "获取成功", utils.PageDataStruct{
		Page:  bindData.Page,
		Limit: bindData.Limit,
		Lists: list,
	})
}
