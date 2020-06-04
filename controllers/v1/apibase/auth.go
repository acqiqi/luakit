package apibase

import (
	"github.com/astaxie/beego"
	"log"
	"luakit/models"
	"luakit/utils"
)

// Platform 控制器验证
func AuthPlatformBase(c beego.Controller) (platform *models.UcenterPlatform) {
	platform_key := c.Ctx.Request.Header.Get("PlatformKey")
	if platform_key == "" {
		utils.ApiOpt(c, utils.ApiPlatformErr, utils.GetApiMsg(utils.ApiPlatformErr), nil)
	}
	log.Println(platform_key)

	platform, err := models.GetPlatformByKey(platform_key)
	if err != nil {
		log.Println(err.Error())
		utils.ApiOpt(c, utils.ApiPlatformErr, utils.GetApiMsg(utils.ApiPlatformErr), nil)
	}
	return
}
