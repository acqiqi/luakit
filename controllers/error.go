package controllers

import (
	"github.com/astaxie/beego"
	"luakit/utils"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["json"] = utils.ApiOpt(utils.Api404, utils.GetApiMsg(utils.Api404), nil)
	c.ServeJSON()

}

func (c *ErrorController) Error501() {
	c.Data["content"] = "server error"
	c.TplName = "501.tpl"
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "dberror.tpl"
}
