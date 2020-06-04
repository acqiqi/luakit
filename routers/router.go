package routers

import (
	"github.com/astaxie/beego"
	"luakit/controllers"
	"luakit/controllers/sm"
	"luakit/controllers/v1/message"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.ErrorController(&controllers.ErrorController{})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/sm",
			beego.NSAutoRouter(
				&sm.ApiController{},
			),
		),
		beego.NSNamespace("/message",
			beego.NSAutoRouter(
				&message.ApiController{},
			),
		),
	)
	beego.AddNamespace(ns)
	//beego.AutoRouter(&sm.SmController{})
}
