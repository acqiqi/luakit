package routers

import (
	"github.com/astaxie/beego"
	"luakit/controllers"
	"luakit/controllers/sm"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.ErrorController(&controllers.ErrorController{})
	beego.AutoRouter(&sm.SmController{})
}
