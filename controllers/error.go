package controllers

import (
	"github.com/astaxie/beego"
	"luakit/utils"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	utils.ApiOpt(c.Controller, utils.Api404, utils.GetApiMsg(utils.Api404), nil)
	return
}

func (c *ErrorController) Error501() {
	utils.ApiOpt(c.Controller, utils.Api501, utils.GetApiMsg(utils.Api501), nil)
	return
}

func (c *ErrorController) ErrorDb() {
	utils.ApiOpt(c.Controller, utils.Api501, utils.GetApiMsg(utils.Api501), nil)
	return
}
