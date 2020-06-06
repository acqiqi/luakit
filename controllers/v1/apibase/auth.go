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

// 带token 控制器验证
func Auth2PlatformBase(c beego.Controller) (platform *models.UcenterPlatform) {
	platform_key := c.Ctx.Request.Header.Get("PlatformKey")
	if platform_key == "" {
		utils.ApiOpt(c, utils.ApiPlatformErr, utils.GetApiMsg(utils.ApiPlatformErr), nil)
	}

	platform, err := models.GetPlatformByKey(platform_key)
	if err != nil {
		log.Println(err.Error())
		utils.ApiOpt(c, utils.ApiPlatformErr, utils.GetApiMsg(utils.ApiPlatformErr), nil)
	}

	token := c.Ctx.Request.Header.Get("Token")
	if token == "" {
		utils.ApiOpt(c, utils.ApiNotAuth, "请设置正确的Token", nil)
	}

	platform_key = utils.GetCacheString("VK_PTOKEN" + token)

	if platform.PlatformKey != platform_key {
		utils.ApiErr(c, "登录超时或Token不正确")
	}
	return
}
